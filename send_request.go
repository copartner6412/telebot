package telegram

import (
	"fmt"
	"io"
	"net/http"
)

// sendRequest sends an HTTP request to the Telegram API.
func (b *Bot) sendRequest(httpMethod, contentTypeHeader, telegramMethod string, requestPayload io.ReadWriter) (*http.Response, error) {
	url := fmt.Sprintf(telegramEndpoint, b.Token, telegramMethod)

	httpRequest, err := http.NewRequest(httpMethod, url, requestPayload)
	if err != nil {
		return nil, fmt.Errorf("error creating new %s request to %s: %w", httpMethod, telegramMethod, err)
	}

	httpRequest.Header.Set("Content-Type", contentTypeHeader)

	httpResponse, err := b.client.Do(httpRequest)
	if err != nil {
		return nil, fmt.Errorf("error sending %s request to %s: %w", httpMethod, telegramMethod, err)
	}

	return httpResponse, nil
}
