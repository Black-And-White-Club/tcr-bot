package db

import (
	leaderboarddb "github.com/Black-And-White-Club/tcr-bot/app/modules/leaderboard/infrastructure/repositories"
	rounddb "github.com/Black-And-White-Club/tcr-bot/app/modules/round/infrastructure/repositories"
	userdb "github.com/Black-And-White-Club/tcr-bot/app/modules/user/infrastructure/repositories"
)

const (
	// DBTypePostgres represents an underlying POSTGRES database type.
	DatabaseTypePostgres string = "POSTGRES"
)

// DB provides methods for interacting with an underlying database or other storage mechanism.
type Database interface {
	leaderboarddb.LeaderboardDB
	rounddb.RoundDB
	// scoredb.ScoreDB
	userdb.UserDB
}
