package telegram

// VideoNote represents a video message (available in Telegram apps as of v.4.0).
//
// See "VideoNote" https://core.telegram.org/bots/api#videonote
type VideoNote struct {
	// (Required) Identifier for this file, which can be used to download or reuse the file.
	FileID string `json:"file_id"`

	// (Required) Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// (Required) Video width and height (diameter of the video message) as defined by the sender.
	Length int `json:"length"`

	// (Required) Duration of the video in seconds as defined by the sender.
	Duration int `json:"duration"`

	// (Optional) Video thumbnail.
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`

	// (Optional) File size in bytes.
	FileSize *int `json:"file_size,omitempty"`
}
