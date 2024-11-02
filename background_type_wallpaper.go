package telegram

// BackgroundTypeWallpaper represents the background as a wallpaper in the JPEG format.
//
// See "BackgroundTypeWallpaper" https://core.telegram.org/bots/api#backgroundtypewallpaper
type BackgroundTypeWallpaper struct {
	// (Required) Type of the background, always “wallpaper”.
	Type string `json:"type"`

	// (Required) Document with the wallpaper.
	Document Document `json:"document"`

	// (Optional) Dimming of the background in dark themes, as a percentage; 0-100.
	DarkThemeDimming *int `json:"dark_theme_dimming,omitempty"`

	// (Optional) True, if the wallpaper is downscaled to fit in a 450x450 square and then box-blurred with radius 12.
	IsBlurred *bool `json:"is_blurred,omitempty"`

	// (Optional) True, if the background moves slightly when the device is tilted.
	IsMoving *bool `json:"is_moving,omitempty"`
}
