package telegram

// ChatBoostAdded represents a service message about a user boosting a chat.
//
// See "ChatBoostAdded" https://core.telegram.org/bots/api#chatboostadded
type ChatBoostAdded struct {
	// (Required) Number of boosts added by the user.
	BoostCount int `json:"boost_count"`
}
