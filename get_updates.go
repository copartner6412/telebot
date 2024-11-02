package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetUpdatesRequest represents a request to receive incoming updates using long polling.
//
// See "getUpdates" https://core.telegram.org/bots/api#getupdates
type GetUpdatesRequest struct {
	// (Optional) Identifier of the first update to be returned. Must be greater by one than the highest among the identifiers of previously received updates.
	// By default, updates starting with the earliest unconfirmed update are returned. An update is considered confirmed as soon as getUpdates is called
	// with an offset higher than its update_id. The negative offset can be specified to retrieve updates starting from -offset update from the end of the updates queue.
	// All previous updates will be forgotten.
	Offset *int `json:"offset,omitempty"`

	// (Optional) Limits the number of updates to be retrieved. Values between 1-100 are accepted. Defaults to 100.
	Limit *int `json:"limit,omitempty"`

	// (Optional) Timeout in seconds for long polling. Defaults to 0, i.e. usual short polling. Should be positive, short polling should be used for testing purposes only.
	Timeout *int `json:"timeout,omitempty"`

	// (Optional) A JSON-serialized list of the update types you want your bot to receive. For example, specify ["message", "edited_channel_post", "callback_query"]
	// to only receive updates of these types. See Update for a complete list of available update types. Specify an empty list to receive all update types
	// except chat_member, message_reaction, and message_reaction_count (default). If not specified, the previous setting will be used.
	//
	// Please note that this parameter doesn't affect updates created before the call to the getUpdates, so unwanted updates may be received for a short period of time.
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

// GetUpdates sends a request to the Telegram API to retrieve incoming updates using long polling.
//
// This method will not work if an outgoing webhook is set up.
//
// In order to avoid getting duplicate updates, recalculate offset after each server response.
//
// See "getUpdates" https://core.telegram.org/bots/api#getupdates
func (b *Bot) getUpdates(request GetUpdatesRequest) ([]Update, error) {
	requestPayload := new(bytes.Buffer)
	if err := json.NewEncoder(requestPayload).Encode(request); err != nil {
		return nil, fmt.Errorf("error encoding request payload: %w", err)
	}

	httpResponse, err := b.sendRequest("POST", "application/json", "getUpdates", requestPayload)
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()

	var response = struct {
		Ok          bool               `json:"ok"`
		Result      []Update           `json:"result"`
		Description string             `json:"description"`
		ErrorCode   int                `json:"error_code"`
		Parameters  ResponseParameters `json:"parameters"`
	}{}

	if err := json.NewDecoder(httpResponse.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	if httpResponse.StatusCode != http.StatusOK || !response.Ok {
		return nil, fmt.Errorf("HTTP status %s, Telegram code  %d, Telegram API error: %s", httpResponse.Status, response.ErrorCode, response.Description)
	}

	return response.Result, nil
}
