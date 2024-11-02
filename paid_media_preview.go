package telegram

// PaidMediaPreview the paid media isn't available before the payment.
//
// See "PaidMediaPreview" https://core.telegram.org/bots/api#paidmediapreview
type PaidMediaPreview struct {
	// (Required) Type of the paid media, always “preview”.
	Type string `json:"type"`

	// (Optional) Media width as defined by the sender.
	Width *int `json:"width,omitempty"`

	// (Optional) Media height as defined by the sender.
	Height *int `json:"height,omitempty"`

	// (Optional) Duration of the media in seconds as defined by the sender.
	Duration *int `json:"duration,omitempty"`
}
