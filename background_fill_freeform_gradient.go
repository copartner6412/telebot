package telegram

// BackgroundFillFreeformGradient represents the background filled with a freeform gradient.
//
// See "BackgroundFillFreeformGradient" https://core.telegram.org/bots/api#backgroundfillfreeformgradient
type BackgroundFillFreeformGradient struct {
	// (Required) Type of the background fill, always “freeform_gradient”.
	Type string `json:"type"`

	// (Required) A list of the 3 or 4 base colors that are used to generate the freeform gradient in the RGB24 format.
	Colors []int `json:"colors"`
}
