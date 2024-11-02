package telegram

// MessageEntity represents one special entity in a text message.
// For example, hashtags, usernames, URLs, etc.
//
// See "MessageEntity" https://core.telegram.org/bots/api#messageentity
type MessageEntity struct {
	// (Required) Type of the entity. Currently, can be “mention” (@username), “hashtag” (#hashtag or #hashtag@chatusername),
	// “cashtag” ($USD or $USD@chatusername), “bot_command” (/start@jobs_bot), “url” (https://telegram.org),
	// “email” (do-not-reply@telegram.org), “phone_number” (+1-212-555-0123), “bold” (bold text),
	// “italic” (italic text), “underline” (underlined text), “strikethrough” (strikethrough text),
	// “spoiler” (spoiler message), “blockquote” (block quotation), “expandable_blockquote” (collapsed-by-default block quotation),
	// “code” (monowidth string), “pre” (monowidth block), “text_link” (for clickable text URLs),
	// “text_mention” (for users without usernames), “custom_emoji” (for inline custom emoji stickers).
	Type string `json:"type"`

	// (Required) Offset in UTF-16 code units to the start of the entity.
	Offset int `json:"offset"`

	// (Required) Length of the entity in UTF-16 code units.
	Length int `json:"length"`

	// (Optional) For “text_link” only, URL that will be opened after user taps on the text.
	URL *string `json:"url,omitempty"`

	// (Optional) For “text_mention” only, the mentioned user.
	User *User `json:"user,omitempty"`

	// (Optional) For “pre” only, the programming language of the entity text.
	Language *string `json:"language,omitempty"`

	// (Optional) For “custom_emoji” only, unique identifier of the custom emoji.
	// Use getCustomEmojiStickers to get full information about the sticker.
	CustomEmojiID *string `json:"custom_emoji_id,omitempty"`
}
