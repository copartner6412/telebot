package telegram

// Document represents a general file (as opposed to photos, voice messages and audio files).
//
// See "Document" https://core.telegram.org/bots/api#document
type Document struct {
	// (Required) Identifier for this file, which can be used to download or reuse the file.
	FileID string `json:"file_id"`

	// (Required) Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// (Optional) Document thumbnail as defined by the sender.
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`

	// (Optional) Original filename as defined by the sender.
	FileName *string `json:"file_name,omitempty"`

	// (Optional) MIME type of the file as defined by the sender.
	MimeType *string `json:"mime_type,omitempty"`

	// (Optional) File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this value.
	FileSize *int `json:"file_size,omitempty"`
}
