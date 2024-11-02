package telegram

// ForumTopicCreated represents a service message about a new forum topic created in the chat.
//
// See "ForumTopicCreated" https://core.telegram.org/bots/api#forumtopiccreated
type ForumTopicCreated struct {
	// (Required) Name of the topic.
	Name string `json:"name"`

	// (Required) Color of the topic icon in RGB format.
	IconColor int `json:"icon_color"`

	// (Optional) Unique identifier of the custom emoji shown as the topic icon.
	IconCustomEmojiID *string `json:"icon_custom_emoji_id,omitempty"`
}
