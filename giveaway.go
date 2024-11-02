package telegram

// Giveaway represents a message about a scheduled giveaway.
//
// See "Giveaway" https://core.telegram.org/bots/api#giveaway
type Giveaway struct {
	// (Required) The list of chats which the user must join to participate in the giveaway
	Chats []Chat `json:"chats"`

	// (Required) Point in time (Unix timestamp) when winners of the giveaway will be selected
	WinnersSelectionDate int64 `json:"winners_selection_date"`

	// (Required) The number of users which are supposed to be selected as winners of the giveaway
	WinnerCount int `json:"winner_count"`

	// (Optional) True, if only users who join the chats after the giveaway started should be eligible to win
	OnlyNewMembers *bool `json:"only_new_members,omitempty"`

	// (Optional) True, if the list of giveaway winners will be visible to everyone
	HasPublicWinners *bool `json:"has_public_winners,omitempty"`

	// (Optional) Description of additional giveaway prize
	PrizeDescription *string `json:"prize_description,omitempty"`

	// (Optional) A list of two-letter ISO 3166-1 alpha-2 country codes indicating the countries from which eligible users for the giveaway must come. If empty, then all users can participate in the giveaway. Users with a phone number that was bought on Fragment can always participate in giveaways.
	CountryCodes *[]string `json:"country_codes,omitempty"`

	// (Optional) The number of Telegram Stars to be split between giveaway winners; for Telegram Star giveaways only
	PrizeStarCount *int `json:"prize_star_count,omitempty"`

	// (Optional) The number of months the Telegram Premium subscription won from the giveaway will be active for; for Telegram Premium giveaways only
	PremiumSubscriptionMonthCount *int `json:"premium_subscription_month_count,omitempty"`
}
