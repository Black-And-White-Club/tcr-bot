// models.go
package rounddb

import (
	"time"
)

// Round represents a single round in the tournament.
type Round struct {
	ID           string        `bun:"id,pk,autoincrement"`
	Title        string        `bun:"title,notnull"`
	Location     string        `bun:"location"`
	EventType    *string       `bun:"event_type"`
	Date         time.Time     `bun:"date,notnull"`
	Time         time.Time     `bun:"time,notnull"`
	Finalized    bool          `bun:"finalized,notnull"`
	CreatorID    string        `bun:"discord_id,notnull"`
	State        RoundState    `bun:"state,notnull"`
	Participants []Participant `bun:"participants,type:jsonb"`
}

// Response represents the possible responses for a participant.
type Response string

// Define the possible response values as constants.
const (
	ResponseAccept    Response = "ACCEPT"
	ResponseTentative Response = "TENTATIVE"
	ResponseDecline   Response = "DECLINE"
)

// RoundState represents the state of a round.
type RoundState string

// Enum constants for RoundState
const (
	RoundStateUpcoming   RoundState = "UPCOMING"
	RoundStateInProgress RoundState = "IN_PROGRESS"
	RoundStateFinalized  RoundState = "FINALIZED"
	RoundStateDeleted    RoundState = "DELETED"
)

// Participant represents a user participating in a round.
type Participant struct {
	DiscordID string   `json:"discord_id"`
	TagNumber *int     `json:"tag_number"`
	Response  Response `json:"response"`
	Score     *int     `json:"score"`
}
