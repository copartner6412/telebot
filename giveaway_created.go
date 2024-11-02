package telegram

// GiveawayCreated represents a service message about the creation of a scheduled giveaway.
//
// See "GiveawayCreated" https://core.telegram.org/bots/api#giveawaycreated
type GiveawayCreated struct {
	// (Optional) The number of Telegram Stars to be split between giveaway winners; for Telegram Star giveaways only
	PrizeStarCount *int `json:"prize_star_count,omitempty"`
}
