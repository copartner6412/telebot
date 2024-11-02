package telegram

// VideoChatParticipantsInvited represents a service message about new members invited to a video chat.
//
// See "VideoChatParticipantsInvited" https://core.telegram.org/bots/api#videochatparticipantsinvited
type VideoChatParticipantsInvited struct {
	// (Required) New members that were invited to the video chat
	Users []User `json:"users"`
}
