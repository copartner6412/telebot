package telegram

// SharedUser contains information about a user that was shared with the bot using a KeyboardButtonRequestUsers button.
//
// See "SharedUser" https://core.telegram.org/bots/api#shareduser
type SharedUser struct {
	// (Required) Identifier of the shared user. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so 64-bit integers or double-precision float types are safe for storing these identifiers. The bot may not have access to the user and could be unable to use this identifier, unless the user is already known to the bot by some other means.
	UserID int64 `json:"user_id"`

	// (Optional) First name of the user, if the name was requested by the bot.
	FirstName *string `json:"first_name,omitempty"`

	// (Optional) Last name of the user, if the name was requested by the bot.
	LastName *string `json:"last_name,omitempty"`

	// (Optional) Username of the user, if the username was requested by the bot.
	Username *string `json:"username,omitempty"`

	// (Optional) Available sizes of the chat photo, if the photo was requested by the bot.
	Photo *[]PhotoSize `json:"photo,omitempty"`
}
