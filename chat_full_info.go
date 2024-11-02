package telegram

// ChatFullInfo contains full information about a chat.
//
// This object holds detailed information about various chat types in Telegram.
//
// See "ChatFullInfo" https://core.telegram.org/bots/api#chatfullinfo
type ChatFullInfo struct {
	// (Required) Unique identifier for this chat. This number may have more than 32 significant bits and some
	// programming languages may have difficulty/silent defects in interpreting it. But it has at most 52
	// significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this
	// identifier.
	ID int64 `json:"id"`

	// (Required) Type of the chat, can be either “private”, “group”, “supergroup” or “channel”.
	Type string `json:"type"`

	// (Optional) Title, for supergroups, channels, and group chats.
	Title *string `json:"title,omitempty"`

	// (Optional) Username, for private chats, supergroups, and channels if available.
	Username *string `json:"username,omitempty"`

	// (Optional) First name of the other party in a private chat.
	FirstName *string `json:"first_name,omitempty"`

	// (Optional) Last name of the other party in a private chat.
	LastName *string `json:"last_name,omitempty"`

	// (Optional) True, if the supergroup chat is a forum (has topics enabled).
	IsForum *bool `json:"is_forum,omitempty"`

	// (Required) Identifier of the accent color for the chat name and backgrounds of the chat photo, reply
	// header, and link preview. See accent colors for more details.
	AccentColorID int `json:"accent_color_id"`

	// (Required) The maximum number of reactions that can be set on a message in the chat.
	MaxReactionCount int `json:"max_reaction_count"`

	// (Optional) Chat photo.
	Photo *ChatPhoto `json:"photo,omitempty"`

	// (Optional) If non-empty, the list of all active chat usernames; for private chats, supergroups, and channels.
	ActiveUsernames []string `json:"active_usernames,omitempty"`

	// (Optional) For private chats, the date of birth of the user.
	Birthdate *Birthdate `json:"birthdate,omitempty"`

	// (Optional) For private chats with business accounts, the intro of the business.
	BusinessIntro *BusinessIntro `json:"business_intro,omitempty"`

	// (Optional) For private chats with business accounts, the location of the business.
	BusinessLocation *BusinessLocation `json:"business_location,omitempty"`

	// (Optional) For private chats with business accounts, the opening hours of the business.
	BusinessOpeningHours *BusinessOpeningHours `json:"business_opening_hours,omitempty"`

	// (Optional) For private chats, the personal channel of the user.
	PersonalChat *Chat `json:"personal_chat,omitempty"`

	// (Optional) List of available reactions allowed in the chat. If omitted, then all emoji reactions are allowed.
	AvailableReactions []ReactionType `json:"available_reactions,omitempty"`

	// (Optional) Custom emoji identifier of the emoji chosen by the chat for the reply header and link preview background.
	BackgroundCustomEmojiID *string `json:"background_custom_emoji_id,omitempty"`

	// (Optional) Identifier of the accent color for the chat's profile background. See profile accent colors for
	// more details.
	ProfileAccentColorID *int `json:"profile_accent_color_id,omitempty"`

	// (Optional) Custom emoji identifier of the emoji chosen by the chat for its profile background.
	ProfileBackgroundCustomEmojiID *string `json:"profile_background_custom_emoji_id,omitempty"`

	// (Optional) Custom emoji identifier of the emoji status of the chat or the other party in a private chat.
	EmojiStatusCustomEmojiID *string `json:"emoji_status_custom_emoji_id,omitempty"`

	// (Optional) Expiration date of the emoji status of the chat or the other party in a private chat, in Unix
	// time, if any.
	EmojiStatusExpirationDate *int `json:"emoji_status_expiration_date,omitempty"`

	// (Optional) Bio of the other party in a private chat.
	Bio *string `json:"bio,omitempty"`

	// (Optional) True, if privacy settings of the other party in the private chat allows to use
	// tg://user?id=<user_id> links only in chats with the user.
	HasPrivateForwards *bool `json:"has_private_forwards,omitempty"`

	// (Optional) True, if the privacy settings of the other party restrict sending voice and video note messages
	// in the private chat.
	HasRestrictedVoiceAndVideoMessages *bool `json:"has_restricted_voice_and_video_messages,omitempty"`

	// (Optional) True, if users need to join the supergroup before they can send messages.
	JoinToSendMessages *bool `json:"join_to_send_messages,omitempty"`

	// (Optional) True, if all users directly joining the supergroup without using an invite link need to be
	// approved by supergroup administrators.
	JoinByRequest *bool `json:"join_by_request,omitempty"`

	// (Optional) Description, for groups, supergroups, and channel chats.
	Description *string `json:"description,omitempty"`

	// (Optional) Primary invite link, for groups, supergroups, and channel chats.
	InviteLink *string `json:"invite_link,omitempty"`

	// (Optional) The most recent pinned message (by sending date).
	PinnedMessage *Message `json:"pinned_message,omitempty"`

	// (Optional) Default chat member permissions, for groups and supergroups.
	Permissions *ChatPermissions `json:"permissions,omitempty"`

	// (Optional) True, if paid media messages can be sent or forwarded to the channel chat. The field is available
	// only for channel chats.
	CanSendPaidMedia *bool `json:"can_send_paid_media,omitempty"`

	// (Optional) For supergroups, the minimum allowed delay between consecutive messages sent by each unprivileged
	// user; in seconds.
	SlowModeDelay *int `json:"slow_mode_delay,omitempty"`

	// (Optional) For supergroups, the minimum number of boosts that a non-administrator user needs to add in order
	// to ignore slow mode and chat permissions.
	UnrestrictBoostCount *int `json:"unrestrict_boost_count,omitempty"`

	// (Optional) The time after which all messages sent to the chat will be automatically deleted; in seconds.
	MessageAutoDeleteTime *int `json:"message_auto_delete_time,omitempty"`

	// (Optional) True, if aggressive anti-spam checks are enabled in the supergroup. The field is only available
	// to chat administrators.
	HasAggressiveAntiSpamEnabled *bool `json:"has_aggressive_anti_spam_enabled,omitempty"`

	// (Optional) True, if non-administrators can only get the list of bots and administrators in the chat.
	HasHiddenMembers *bool `json:"has_hidden_members,omitempty"`

	// (Optional) True, if messages from the chat can't be forwarded to other chats.
	HasProtectedContent *bool `json:"has_protected_content,omitempty"`

	// (Optional) True, if new chat members will have access to old messages; available only to chat administrators.
	HasVisibleHistory *bool `json:"has_visible_history,omitempty"`

	// (Optional) For supergroups, name of the group sticker set.
	StickerSetName *string `json:"sticker_set_name,omitempty"`

	// (Optional) True, if the bot can change the group sticker set.
	CanSetStickerSet *bool `json:"can_set_sticker_set,omitempty"`

	// (Optional) For supergroups, the name of the group's custom emoji sticker set. Custom emoji from this set can
	// be used by all users and bots in the group.
	CustomEmojiStickerSetName *string `json:"custom_emoji_sticker_set_name,omitempty"`

	// (Optional) Unique identifier for the linked chat, i.e. the discussion group identifier for a channel and vice
	// versa; for supergroups and channel chats. This identifier may be greater than 32 bits and some programming
	// languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed
	// 64-bit integer or double-precision float type are safe for storing this identifier.
	LinkedChatID *int64 `json:"linked_chat_id,omitempty"`

	// (Optional) For supergroups, the location to which the supergroup is connected.
	Location *ChatLocation `json:"location,omitempty"`
}

// Total number of fields in the ChatFullInfo struct: 55
