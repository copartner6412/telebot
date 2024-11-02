package telegram

// PaidMediaInfo describes the paid media added to a message.
//
// See "PaidMediaInfo" https://core.telegram.org/bots/api#paidmediainfo
type PaidMediaInfo struct {
	// (Required) The number of Telegram Stars that must be paid to buy access to the media.
	StarCount int `json:"star_count"`

	// (Required) Information about the paid media.
	PaidMedia []PaidMedia `json:"paid_media"`
}
