// db/round.go

package rounddb

import (
	"context"
	"fmt"
	"time"

	"github.com/uptrace/bun"
)

// RoundDBImpl is the concrete implementation of the RoundDB interface using bun.
type RoundDBImpl struct {
	DB *bun.DB
}

// GetRounds retrieves all rounds.
func (r *RoundDBImpl) GetRounds(ctx context.Context) ([]*Round, error) {
	var rounds []*Round
	err := r.DB.NewSelect().
		Model(&rounds).
		Relation("Participants").
		Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch rounds: %w", err)
	}
	return rounds, nil
}

// GetRound retrieves a specific round by ID.
func (r *RoundDBImpl) GetRound(ctx context.Context, roundID int64) (*Round, error) {
	var round Round
	err := r.DB.NewSelect().
		Model(&round).
		Where("id = ?", roundID).
		Relation("Participants").
		Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch round: %w", err)
	}
	return &round, nil
}

// CreateRound creates a new round in the database.
func (r *RoundDBImpl) CreateRound(ctx context.Context, input ScheduleRoundInput) (*Round, error) {
	round := &Round{
		Title:     input.Title,
		Location:  input.Location,
		EventType: input.EventType,
		Date:      input.Date,
		Time:      input.Time,
		CreatorID: input.DiscordID,
		State:     RoundStateUpcoming, // Set initial state to "UPCOMING"
	}
	_, err := r.DB.NewInsert().
		Model(round).
		Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create round: %w", err)
	}
	return round, nil
}

// UpdateRound updates an existing round in the database.
func (r *RoundDBImpl) UpdateRound(ctx context.Context, roundID int64, input EditRoundInput) error {
	round := &Round{
		ID:        roundID,
		Title:     input.Title,
		Location:  input.Location,
		EventType: input.EventType,
		Date:      input.Date,
		Time:      input.Time,
	}

	_, err := r.DB.NewUpdate().
		Model(round).
		WherePK().
		Column("title", "location", "event_type", "date", "time"). // Use Column to specify fields
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to update round: %w", err)
	}

	return nil
}

// DeleteRound deletes a round by ID.
func (r *RoundDBImpl) DeleteRound(ctx context.Context, roundID int64) error { // No userID parameter
	_, err := r.DB.NewDelete().
		Model((*Round)(nil)).
		Where("id = ?", roundID).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete round: %w", err)
	}
	return nil
}

// SubmitScore updates the scores map for a round in the database.
func (r *RoundDBImpl) SubmitScore(ctx context.Context, roundID int64, discordID string, score int) error {
	var round Round
	err := r.DB.NewSelect().
		Model(&round).
		Where("id = ?", roundID).
		Scan(ctx)
	if err != nil {
		return fmt.Errorf("failed to fetch round: %w", err)
	}

	if round.Scores == nil {
		round.Scores = make(map[string]int)
	}
	round.Scores[discordID] = score

	_, err = r.DB.NewUpdate().
		Model(&round).
		Where("id = ?", roundID).
		Column("scores").
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to update round scores: %w", err)
	}

	return nil
}

// UpdateParticipantResponse updates a participant's response or tag number in a round.
func (r *RoundDBImpl) UpdateParticipant(ctx context.Context, roundID int64, participant Participant) error {
	var round Round
	err := r.DB.NewSelect().
		Model(&round).
		Where("id = ?", roundID).
		Scan(ctx)
	if err != nil {
		return fmt.Errorf("failed to fetch round: %w", err)
	}

	// Find the participant and update their response or tag number
	found := false
	for i, p := range round.Participants {
		if p.DiscordID == participant.DiscordID {
			if participant.Response != "" { // Update response if provided
				round.Participants[i].Response = participant.Response
			}
			if participant.TagNumber != nil { // Update tag number if provided
				round.Participants[i].TagNumber = participant.TagNumber
			}
			found = true
			break
		}
	}
	if !found {
		// If participant not found, add them to the round
		round.Participants = append(round.Participants, participant)
	}

	_, err = r.DB.NewUpdate().
		Model(&round).
		Where("id = ?", roundID).
		Column("participants").
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to update participant response: %w", err)
	}

	return nil
}

// UpdateRoundState updates the state of a round.
func (r *RoundDBImpl) UpdateRoundState(ctx context.Context, roundID int64, state RoundState) error {
	_, err := r.DB.NewUpdate().
		Model((*Round)(nil)).
		Set("state = ?", state).
		Where("id = ?", roundID).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to update round state: %w", err)
	}
	return nil
}

// GetUpcomingRounds retrieves rounds that are upcoming within the given time range.
func (r *RoundDBImpl) GetUpcomingRounds(ctx context.Context, now, oneHourFromNow time.Time) ([]*Round, error) {
	var rounds []*Round
	err := r.DB.NewSelect().
		Model(&rounds).
		Where("state = ? AND date = ? AND time BETWEEN ? AND ?", RoundStateUpcoming, now.Format("2006-01-02"), now.Format("15:04"), oneHourFromNow.Format("15:04")).
		Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch upcoming rounds: %w", err)
	}
	return rounds, nil
}

// IsRoundFinalized checks if a round is finalized.
func (r *RoundDBImpl) IsRoundFinalized(ctx context.Context, roundID int64) (bool, error) {
	var round Round
	err := r.DB.NewSelect().
		Model(&round).
		Column("finalized").
		Where("id = ?", roundID).
		Scan(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to check round finalized status: %w", err)
	}
	return round.Finalized, nil
}

// IsUserParticipant checks if a user is a participant in a round.
func (r *RoundDBImpl) IsUserParticipant(ctx context.Context, roundID int64, userID string) (bool, error) {
	// Assuming your Participant struct has a DiscordID field
	var participant Participant
	err := r.DB.NewSelect().
		Model(&participant).
		Where("jsonb_exists(participants, ?) AND participants->>? = ?", userID, "discord_id", userID). // Adjust the query based on your JSON structure
		Scan(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to check participant status: %w", err)
	}
	return true, nil // If no error, the user is a participant
}

// GetRoundState retrieves the state of a round.
func (r *RoundDBImpl) GetRoundState(ctx context.Context, roundID int64) (RoundState, error) {
	var round Round
	err := r.DB.NewSelect().
		Model(&round).
		Column("state").
		Where("id = ?", roundID).
		Scan(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get round state: %w", err)
	}
	return round.State, nil
}

// RoundExists checks if a round with the given ID exists.
func (r *RoundDBImpl) RoundExists(ctx context.Context, roundID int64) (bool, error) {
	exists, err := r.DB.NewSelect().
		Model((*Round)(nil)).
		Where("id = ?", roundID).
		Exists(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to check if round exists: %w", err)
	}
	return exists, nil
}