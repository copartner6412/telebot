package telegram

// BackgroundFillSolid represents the background filled using the selected color.
//
// See "BackgroundFillSolid" https://core.telegram.org/bots/api#backgroundfillsolid
type BackgroundFillSolid struct {
	// (Required) Type of the background fill, always “solid”.
	Type string `json:"type"`

	// (Required) The color of the background fill in the RGB24 format.
	Color int `json:"color"`
}
