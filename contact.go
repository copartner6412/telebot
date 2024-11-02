package telegram

// Contact represents a phone contact.
//
// See "Contact" https://core.telegram.org/bots/api#contact
type Contact struct {
	// (Required) Contact's phone number.
	PhoneNumber string `json:"phone_number"`

	// (Required) Contact's first name.
	FirstName string `json:"first_name"`

	// (Optional) Contact's last name.
	LastName *string `json:"last_name,omitempty"`

	// (Optional) Contact's user identifier in Telegram. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier.
	UserID *int64 `json:"user_id,omitempty"`

	// (Optional) Additional data about the contact in the form of a vCard.
	VCard *string `json:"vcard,omitempty"`
}
