package userdb

import (
	"context"

	usertypes "github.com/Black-And-White-Club/frolf-bot/app/modules/user/domain/types"
)

// UserDB is an interface for interacting with the user database.
type UserDB interface {
	// WithinTransaction(ctx context.Context, fn func(ctx context.Context) error) error
	CreateUser(ctx context.Context, user *User) error
	GetUserByDiscordID(ctx context.Context, discordID usertypes.DiscordID) (*User, error)
	GetUserRole(ctx context.Context, discordID usertypes.DiscordID) (usertypes.UserRoleEnum, error)
	UpdateUserRole(ctx context.Context, discordID usertypes.DiscordID, role usertypes.UserRoleEnum) error
}
