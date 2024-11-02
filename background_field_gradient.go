package telegram

// BackgroundFillGradient represents the background filled with a gradient.
//
// See "BackgroundFillGradient" https://core.telegram.org/bots/api#backgroundfillgradient
type BackgroundFillGradient struct {
	// (Required) Type of the background fill, always “gradient”.
	Type string `json:"type"`

	// (Required) Top color of the gradient in the RGB24 format.
	TopColor int `json:"top_color"`

	// (Required) Bottom color of the gradient in the RGB24 format.
	BottomColor int `json:"bottom_color"`

	// (Required) Clockwise rotation angle of the background fill in degrees; 0-359.
	RotationAngle int `json:"rotation_angle"`
}
