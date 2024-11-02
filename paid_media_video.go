package telegram

// PaidMediaVideo the paid media is a video.
//
// See "PaidMediaVideo" https://core.telegram.org/bots/api#paidmediavideo
type PaidMediaVideo struct {
	// (Required) Type of the paid media, always “video”.
	Type string `json:"type"`

	// (Required) The video.
	Video Video `json:"video"`
}
