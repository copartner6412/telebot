package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// DeleteWebhookRequest represents a request to remove webhook integration.
//
// See "deleteWebhook" https://core.telegram.org/bots/api#deletewebhook
type DeleteWebhookRequest struct {
	// (Optional) Pass True to drop all pending updates.
	DropPendingUpdates bool `json:"drop_pending_updates,omitempty"`
}

// DeleteWebhook removes webhook integration if the bot decides to switch back to getUpdates.
//
// See "deleteWebhook" https://core.telegram.org/bots/api#deletewebhook
func (b *Bot) DeleteWebhook(request DeleteWebhookRequest) error {
	requestPayload := new(bytes.Buffer)

	if err := json.NewEncoder(requestPayload).Encode(request); err != nil {
		return fmt.Errorf("error encoding request payload: %w", err)
	}

	httpResponse, err := b.sendRequest("POST", "application/json", "deleteWebhook", requestPayload)
	if err != nil {
		return err
	}
	defer httpResponse.Body.Close()

	var response struct {
		Ok          bool               `json:"ok"`
		Result      bool               `json:"result"`
		Description string             `json:"description"`
		ErrorCode   int                `json:"error_code"`
		Parameters  ResponseParameters `json:"parameters"`
	}

	if err := json.NewDecoder(httpResponse.Body).Decode(&response); err != nil {
		return fmt.Errorf("error decoding response: %w", err)
	}

	if httpResponse.StatusCode != http.StatusOK || !response.Ok {
		return fmt.Errorf("HTTP status %s, Telegram code %d, Telegram API error: %s",
			httpResponse.Status, response.ErrorCode, response.Description)
	}

	return nil
}
