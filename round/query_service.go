// round/query_service.go
package round

import (
	"context"
	"fmt"
	"time"

	rounddb "github.com/Black-And-White-Club/tcr-bot/round/db"
)

// RoundQueryService handles query-related logic for rounds.
type RoundQueryService struct {
	roundDB   rounddb.RoundDB
	converter RoundConverter
}

// NewRoundQueryService creates a new RoundQueryService.
func NewRoundQueryService(roundDB rounddb.RoundDB) *RoundQueryService {
	return &RoundQueryService{
		roundDB:   roundDB,
		converter: &DefaultRoundConverter{},
	}
}

// GetRounds retrieves all rounds.
func (s *RoundQueryService) GetRounds(ctx context.Context) ([]*Round, error) {
	modelRounds, err := s.roundDB.GetRounds(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get rounds: %w", err)
	}

	var apiRounds []*Round
	for _, modelRound := range modelRounds {
		apiRounds = append(apiRounds, s.converter.ConvertModelRoundToStructRound(modelRound))
	}

	return apiRounds, nil
}

// GetRound retrieves a specific round by ID.
func (s *RoundQueryService) GetRound(ctx context.Context, roundID int64) (*Round, error) {
	return getRound(ctx, s.roundDB, s.converter, roundID)
}

// HasActiveRounds checks if there are any active rounds.
func (s *RoundQueryService) HasActiveRounds(ctx context.Context) (bool, error) {
	// 1. Check for upcoming rounds within the next hour
	now := time.Now()
	oneHourFromNow := now.Add(time.Hour)
	upcomingRounds, err := s.roundDB.GetUpcomingRounds(ctx, now, oneHourFromNow)
	if err != nil {
		return false, fmt.Errorf("failed to get upcoming rounds: %w", err)
	}
	if len(upcomingRounds) > 0 {
		return true, nil // There are upcoming rounds
	}

	// 2. If no upcoming rounds, check for rounds in progress
	rounds, err := s.GetRounds(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to get rounds: %w", err)
	}
	for _, round := range rounds {
		if round.State == RoundStateInProgress {
			return true, nil // There's a round in progress
		}
	}

	// 3. No active rounds found
	return false, nil
}
