package telegram

// ForumTopicEdited represents a service message about an edited forum topic.
//
// See "ForumTopicEdited" https://core.telegram.org/bots/api#forumtopicedited
type ForumTopicEdited struct {
	// (Optional) New name of the topic, if it was edited.
	Name *string `json:"name,omitempty"`

	// (Optional) New identifier of the custom emoji shown as the topic icon, if it was edited; an empty string if the icon was removed.
	IconCustomEmojiID *string `json:"icon_custom_emoji_id,omitempty"`
}
