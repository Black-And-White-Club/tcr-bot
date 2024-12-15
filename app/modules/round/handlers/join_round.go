package roundhandlers

import (
	"context"
	"encoding/json"
	"fmt"

	roundcommands "github.com/Black-And-White-Club/tcr-bot/app/modules/round/commands"
	rounddb "github.com/Black-And-White-Club/tcr-bot/app/modules/round/db"
	rounddto "github.com/Black-And-White-Club/tcr-bot/app/modules/round/dto"
	watermillutil "github.com/Black-And-White-Club/tcr-bot/internal/watermill"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

// JoinRoundHandler handles the JoinRoundRequest command.
type JoinRoundHandler struct {
	roundDB  rounddb.RoundDB
	eventBus watermillutil.PubSuber
}

// NewJoinRoundHandler creates a new JoinRoundHandler.
func NewJoinRoundHandler(roundDB rounddb.RoundDB, eventBus watermillutil.PubSuber) *JoinRoundHandler {
	return &JoinRoundHandler{
		roundDB:  roundDB,
		eventBus: eventBus,
	}
}

// Handle processes the JoinRoundRequest command.
func (h *JoinRoundHandler) Handle(ctx context.Context, msg *message.Message) error {
	var dto rounddto.JoinRoundInput
	if err := json.Unmarshal(msg.Payload, &dto); err != nil {
		return fmt.Errorf("failed to unmarshal JoinRoundRequest: %w", err)
	}

	// 1. Validate the command
	if dto.RoundID <= 0 {
		return fmt.Errorf("invalid RoundID")
	}
	if dto.DiscordID == "" {
		return fmt.Errorf("invalid DiscordID")
	}
	if dto.Response == "" {
		return fmt.Errorf("invalid Response")
	}

	exists, err := h.roundDB.RoundExists(ctx, dto.RoundID)
	if err != nil {
		return fmt.Errorf("failed to check if round exists: %w", err)
	}
	if !exists {
		return fmt.Errorf("round does not exist")
	}

	// Check if the user is already a participant using GetParticipant
	_, err = h.roundDB.GetParticipant(ctx, dto.RoundID, dto.DiscordID)
	if err == nil {
		return fmt.Errorf("user is already a participant in this round")
	}

	// 2. Publish a GetTagNumberRequest event to the leaderboard module
	getTagNumberRequest := roundcommands.GetTagNumberRequest{
		DiscordID: dto.DiscordID,
	}
	payload, err := json.Marshal(getTagNumberRequest)
	if err != nil {
		return fmt.Errorf("failed to marshal GetTagNumberRequest: %w", err)
	}
	if err := h.eventBus.Publish("leaderboard.get.tag.number.request", message.NewMessage(watermill.NewUUID(), payload)); err != nil {
		return fmt.Errorf("failed to publish GetTagNumberRequest: %w", err)
	}

	// 3. Add the user as a participant to the round in the database (without tag number initially)
	participant := rounddb.Participant{
		DiscordID: dto.DiscordID,
		Response:  rounddb.Response(dto.Response), // Convert to rounddb.Response
	}
	err = h.roundDB.UpdateParticipant(ctx, dto.RoundID, participant)
	if err != nil {
		return fmt.Errorf("failed to add participant to round: %w", err)
	}

	// 4. Fetch the updated participant from the database (now with tag number)
	updatedParticipant, err := h.roundDB.GetParticipant(ctx, dto.RoundID, dto.DiscordID)
	if err != nil {
		return fmt.Errorf("failed to get updated participant: %w", err)
	}

	// 5. Publish a ParticipantJoinedRound event (using the updated participant)
	event := ParticipantJoinedRoundEvent{
		RoundID:     dto.RoundID,
		Participant: updatedParticipant,
	}
	payload, err = json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal ParticipantJoinedRoundEvent: %w", err)
	}
	if err := h.eventBus.Publish(TopicJoinRound, message.NewMessage(watermill.NewUUID(), payload)); err != nil {
		return fmt.Errorf("failed to publish ParticipantJoinedRoundEvent: %w", err)
	}

	return nil
}
