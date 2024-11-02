package telegram

// GiveawayCompleted represents a service message about the completion of a giveaway without public winners.
//
// See "GiveawayCompleted" https://core.telegram.org/bots/api#backgroundtypefill
type GiveawayCompleted struct {
	// (Required) Number of winners in the giveaway
	WinnerCount int `json:"winner_count"`

	// (Optional) Number of undistributed prizes
	UnclaimedPrizeCount *int `json:"unclaimed_prize_count,omitempty"`

	// (Optional) Message with the giveaway that was completed, if it wasn't deleted
	GiveawayMessage *Message `json:"giveaway_message,omitempty"`

	// (Optional) True, if the giveaway is a Telegram Star giveaway. Otherwise, currently, the giveaway is a Telegram Premium giveaway.
	IsStarGiveaway *bool `json:"is_star_giveaway,omitempty"`
}
