package telegram

// PaidMediaPhoto the paid media is a photo.
//
// See "PaidMediaPhoto" https://core.telegram.org/bots/api#paidmediaphoto
type PaidMediaPhoto struct {
	// (Required) Type of the paid media, always “photo”.
	Type string `json:"type"`

	// (Required) The photo.
	Photo []PhotoSize `json:"photo"`
}
