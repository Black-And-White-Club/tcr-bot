package roundservice

import (
	"context"
	"fmt"
	"strings"
	"time"

	roundevents "github.com/Black-And-White-Club/frolf-bot-shared/events/round"
	"github.com/Black-And-White-Club/frolf-bot/app/shared/logging"
	"github.com/Black-And-White-Club/frolf-bot/internal/eventutil"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
)

// -- Service Functions for UpdateRound Flow --

// ValidateRoundUpdateRequest validates the round update request.
func (s *RoundService) ValidateRoundUpdateRequest(ctx context.Context, msg *message.Message) error {
	_, eventPayload, err := eventutil.UnmarshalPayload[roundevents.RoundUpdateRequestPayload](msg, s.logger)
	if err != nil {
		return s.publishRoundUpdateError(msg, eventPayload, fmt.Errorf("invalid payload: %w", err))
	}

	var errs []string
	if eventPayload.RoundID == "" {
		errs = append(errs, "round ID cannot be empty")
	}
	if eventPayload.Title == nil && eventPayload.Location == nil && eventPayload.EventType == nil && eventPayload.Date == nil && eventPayload.Time == nil {
		errs = append(errs, "at least one field to update must be provided")
	}

	if len(errs) > 0 {
		errMsg := strings.Join(errs, "; ")
		err := fmt.Errorf("%s", errMsg)
		logging.LogErrorWithMetadata(ctx, s.logger, msg, "Round update input validation failed", map[string]interface{}{
			"errors": errs,
			"error":  err.Error(),
		})
		return s.publishRoundUpdateError(msg, eventPayload, fmt.Errorf("%s", errMsg))
	}

	// If validation passes, publish a "round.update.validated" event
	if err := s.publishEvent(msg, roundevents.RoundUpdateValidated, roundevents.RoundUpdateValidatedPayload{
		RoundUpdateRequestPayload: eventPayload,
	}); err != nil {
		logging.LogErrorWithMetadata(ctx, s.logger, msg, "Failed to publish round.update.validated event", map[string]interface{}{"error": err.Error()})
		return fmt.Errorf("failed to publish round.update.validated event: %w", err)
	}

	logging.LogInfoWithMetadata(ctx, s.logger, msg, "Round update input validated", map[string]interface{}{"round_id": eventPayload.RoundID})
	return nil
}

// UpdateRoundEntity updates the round entity with the new values.
func (s *RoundService) UpdateRoundEntity(ctx context.Context, msg *message.Message) error {
	_, eventPayload, err := eventutil.UnmarshalPayload[roundevents.RoundFetchedPayload](msg, s.logger)
	if err != nil {
		s.logger.Error("Unmarshal failed in UpdateRoundEntity", "error", err)
		return s.publishRoundUpdateError(msg, eventPayload.RoundUpdateRequestPayload, fmt.Errorf("invalid payload: %w", err))
	}

	// Fetch the existing round
	existingRound, err := s.RoundDB.GetRound(ctx, eventPayload.Round.ID)
	if err != nil {
		s.logger.Error("Failed to fetch round", "round_id", eventPayload.Round.ID, "error", err)
		return s.publishRoundUpdateError(msg, eventPayload.RoundUpdateRequestPayload, fmt.Errorf("failed to fetch existing round: %w", err))
	}

	// Apply updates
	if eventPayload.RoundUpdateRequestPayload.Title != nil {
		existingRound.Title = *eventPayload.RoundUpdateRequestPayload.Title
	}
	if eventPayload.RoundUpdateRequestPayload.Location != nil {
		existingRound.Location = *eventPayload.RoundUpdateRequestPayload.Location
	}
	if eventPayload.RoundUpdateRequestPayload.EventType != nil { // <-- Add this!
		existingRound.EventType = eventPayload.RoundUpdateRequestPayload.EventType
	}
	if eventPayload.RoundUpdateRequestPayload.Date != nil && eventPayload.RoundUpdateRequestPayload.Time != nil {
		existingRound.StartTime = time.Date(
			eventPayload.RoundUpdateRequestPayload.Date.Year(),
			eventPayload.RoundUpdateRequestPayload.Date.Month(),
			eventPayload.RoundUpdateRequestPayload.Date.Day(),
			eventPayload.RoundUpdateRequestPayload.Time.Hour(),
			eventPayload.RoundUpdateRequestPayload.Time.Minute(),
			0, 0, time.UTC,
		)
	}

	// Update the round in the database
	err = s.RoundDB.UpdateRound(ctx, existingRound.ID, existingRound)
	if err != nil {
		s.logger.Error("Failed to update round entity", "round_id", existingRound.ID, "error", err)
		return s.publishRoundUpdateError(msg, eventPayload.RoundUpdateRequestPayload, fmt.Errorf("failed to update round entity: %w", err))
	}

	// Successfully updated round → Publish "round.updated" event
	err = s.publishEvent(msg, roundevents.RoundUpdated, roundevents.RoundUpdatedPayload{
		RoundID: existingRound.ID,
	})
	if err != nil {
		s.logger.Error("Failed to publish round.updated event", "round_id", existingRound.ID, "error", err)
		return fmt.Errorf("failed to publish round.updated event: %w", err)
	}

	s.logger.Info("Round successfully updated", "round_id", existingRound.ID)
	return nil
}

// StoreRoundUpdate updates the round in the database.
func (s *RoundService) StoreRoundUpdate(ctx context.Context, msg *message.Message) error {
	_, eventPayload, err := eventutil.UnmarshalPayload[roundevents.RoundEntityUpdatedPayload](msg, s.logger)
	if err != nil {
		return s.publishRoundUpdateError(msg, roundevents.RoundUpdateRequestPayload{}, fmt.Errorf("invalid payload: %w", err))
	}

	if err := s.RoundDB.UpdateRound(ctx, eventPayload.Round.ID, &eventPayload.Round); err != nil {
		return s.publishRoundUpdateError(msg, roundevents.RoundUpdateRequestPayload{}, err)
	}

	// Publish "round.updated" event
	if err := s.publishEvent(msg, roundevents.RoundUpdated, roundevents.RoundUpdatedPayload{
		RoundID: eventPayload.Round.ID,
	}); err != nil {
		logging.LogErrorWithMetadata(ctx, s.logger, msg, "Failed to publish round.updated event", map[string]interface{}{"error": err.Error()})
		return fmt.Errorf("failed to publish round.updated event: %w", err)
	}

	logging.LogInfoWithMetadata(ctx, s.logger, msg, "Round updated in database", map[string]interface{}{"round_id": eventPayload.Round.ID})
	return nil
}

// publishRoundUpdateError publishes a round.update.error event with details.
func (s *RoundService) publishRoundUpdateError(msg *message.Message, input roundevents.RoundUpdateRequestPayload, err error) error {
	payload := roundevents.RoundUpdateErrorPayload{
		CorrelationID:      middleware.MessageCorrelationID(msg),
		RoundUpdateRequest: &input,
		Error:              err.Error(),
	}

	if pubErr := s.publishEvent(msg, roundevents.RoundUpdateError, payload); pubErr != nil {
		s.ErrorReporter.ReportError(middleware.MessageCorrelationID(msg), "Failed to publish round.update.error event", pubErr, "original_error", err.Error())
		return fmt.Errorf("failed to publish round.update.error event: %w, original error: %w", pubErr, err)
	}

	return err
}

// UpdateScheduledRoundEvents updates the scheduled events for a round.
func (s *RoundService) UpdateScheduledRoundEvents(ctx context.Context, msg *message.Message) error {
	_, eventPayload, err := eventutil.UnmarshalPayload[roundevents.RoundUpdatedPayload](msg, s.logger)
	if err != nil {
		return fmt.Errorf("invalid payload: %w", err)
	}

	// Cancel existing scheduled events
	if err := s.EventBus.CancelScheduledMessage(ctx, eventPayload.RoundID); err != nil {
		return fmt.Errorf("failed to cancel existing scheduled events: %w", err)
	}

	// Publish an event to schedule new events
	if err := s.publishEvent(msg, roundevents.RoundScheduleUpdate, roundevents.RoundScheduleUpdatePayload{
		RoundID: eventPayload.RoundID,
	}); err != nil {
		return fmt.Errorf("failed to publish round.schedule.update event: %w", err)
	}

	return nil
}
