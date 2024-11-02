package telegram

// GiveawayWinners represents a message about the completion of a giveaway with public winners.
//
// See "GiveawayWinners" https://core.telegram.org/bots/api#backgroundtypefill
type GiveawayWinners struct {
	// (Required) The chat that created the giveaway
	Chat Chat `json:"chat"`

	// (Required) Identifier of the message with the giveaway in the chat
	GiveawayMessageID int `json:"giveaway_message_id"`

	// (Required) Point in time (Unix timestamp) when winners of the giveaway were selected
	WinnersSelectionDate int64 `json:"winners_selection_date"`

	// (Required) Total number of winners in the giveaway
	WinnerCount int `json:"winner_count"`

	// (Required) List of up to 100 winners of the giveaway
	Winners []User `json:"winners"`

	// (Optional) The number of other chats the user had to join in order to be eligible for the giveaway
	AdditionalChatCount *int `json:"additional_chat_count,omitempty"`

	// (Optional) The number of Telegram Stars that were split between giveaway winners; for Telegram Star giveaways only
	PrizeStarCount *int `json:"prize_star_count,omitempty"`

	// (Optional) The number of months the Telegram Premium subscription won from the giveaway will be active for; for Telegram Premium giveaways only
	PremiumSubscriptionMonthCount *int `json:"premium_subscription_month_count,omitempty"`

	// (Optional) Number of undistributed prizes
	UnclaimedPrizeCount *int `json:"unclaimed_prize_count,omitempty"`

	// (Optional) True, if only users who had joined the chats after the giveaway started were eligible to win
	OnlyNewMembers *bool `json:"only_new_members,omitempty"`

	// (Optional) True, if the giveaway was canceled because the payment for it was refunded
	WasRefunded *bool `json:"was_refunded,omitempty"`

	// (Optional) Description of additional giveaway prize
	PrizeDescription *string `json:"prize_description,omitempty"`
}
