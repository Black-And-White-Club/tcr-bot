package roundhandlers

import (
	"fmt"
	"log/slog"

	roundevents "github.com/Black-And-White-Club/frolf-bot-shared/events/round"
	"github.com/Black-And-White-Club/frolf-bot/internal/eventutil"
	"github.com/ThreeDotsLabs/watermill/message"
)

func (h *RoundHandlers) HandleRoundCreateRequest(msg *message.Message) error {
	correlationID, _, err := eventutil.UnmarshalPayload[roundevents.RoundCreateRequestPayload](msg, h.logger)
	if err != nil {
		return fmt.Errorf("failed to unmarshal RoundCreateRequestPayload: %w", err)
	}

	h.logger.Info("Received RoundCreateRequest event",
		slog.String("correlation_id", correlationID),
	)

	if err := h.RoundService.ValidateRoundRequest(msg.Context(), msg); err != nil {
		h.logger.Error("Failed to handle RoundCreateRequest event",
			slog.String("correlation_id", correlationID),
			slog.Any("error", err),
		)
		return fmt.Errorf("failed to handle RoundCreateRequest event: %w", err)
	}

	h.logger.Info("RoundCreateRequest event processed", slog.String("correlation_id", correlationID))
	return nil
}

func (h *RoundHandlers) HandleRoundValidated(msg *message.Message) error {
	correlationID, _, err := eventutil.UnmarshalPayload[roundevents.RoundValidatedPayload](msg, h.logger)
	if err != nil {
		return fmt.Errorf("failed to unmarshal RoundValidatedPayload: %w", err)
	}

	h.logger.Info("Received RoundValidated event",
		slog.String("correlation_id", correlationID),
	)

	if err := h.RoundService.StoreRound(msg.Context(), msg); err != nil {
		h.logger.Error("Failed to handle RoundValidated event",
			slog.String("correlation_id", correlationID),
			slog.Any("error", err),
		)
		return fmt.Errorf("failed to handle RoundValidated event: %w", err)
	}

	h.logger.Info("RoundValidated event processed", slog.String("correlation_id", correlationID))
	return nil
}

func (h *RoundHandlers) HandleRoundStored(msg *message.Message) error {
	correlationID, _, err := eventutil.UnmarshalPayload[roundevents.RoundStoredPayload](msg, h.logger)
	if err != nil {
		return fmt.Errorf("failed to unmarshal RoundStoredPayload: %w", err)
	}

	h.logger.Info("Received RoundStored event",
		slog.String("correlation_id", correlationID),
	)

	if err := h.RoundService.ScheduleRoundEvents(msg.Context(), msg); err != nil {
		h.logger.Error("Failed to handle RoundStored event",
			slog.String("correlation_id", correlationID),
			slog.Any("error", err),
		)
		return fmt.Errorf("failed to handle RoundStored event: %w", err)
	}

	h.logger.Info("RoundStored event processed", slog.String("correlation_id", correlationID))
	return nil
}

func (h *RoundHandlers) HandleRoundScheduled(msg *message.Message) error {
	correlationID, _, err := eventutil.UnmarshalPayload[roundevents.RoundScheduledPayload](msg, h.logger)
	if err != nil {
		return fmt.Errorf("failed to unmarshal RoundScheduledPayload: %w", err)
	}

	h.logger.Info("Received RoundScheduled event",
		slog.String("correlation_id", correlationID),
	)

	if err := h.RoundService.PublishRoundCreated(msg.Context(), msg); err != nil {
		h.logger.Error("Failed to handle RoundScheduled event",
			slog.String("correlation_id", correlationID),
			slog.Any("error", err),
		)
		return fmt.Errorf("failed to handle RoundScheduled event: %w", err)
	}

	h.logger.Info("RoundScheduled event processed", slog.String("correlation_id", correlationID))
	return nil
}

func (h *RoundHandlers) HandleUpdateDiscordEventID(msg *message.Message) error {
	correlationID, _, err := eventutil.UnmarshalPayload[roundevents.RoundEventCreatedPayload](msg, h.logger)
	if err != nil {
		return fmt.Errorf("failed to unmarshal RoundEventCreatedPayload: %w", err)
	}

	h.logger.Info("Received RoundEventCreated event",
		slog.String("correlation_id", correlationID),
	)

	if err := h.RoundService.UpdateDiscordEventID(msg.Context(), msg); err != nil {
		h.logger.Error("Failed to handle RoundEventCreated event",
			slog.String("correlation_id", correlationID),
			slog.Any("error", err),
		)
		return fmt.Errorf("failed to handle RoundEventCreated event: %w", err)
	}

	h.logger.Info("RoundEventCreated event processed", slog.String("correlation_id", correlationID))
	return nil
}
