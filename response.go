package telegram

// If 'ok' equals True, the request was successful and the result of the query can be found in the 'result' field.
//
// In case of an unsuccessful request, 'ok' equals false and the error is explained in the 'description'.
//
// An Integer 'error_code' field is also returned, but its contents are subject to change in the future.
//
// Some errors may also have an optional field 'parameters' of the type
//
// See "Making requests" https://core.telegram.org/bots/api#making-requests
type Response struct {
	Ok          bool               `json:"ok"`
	Result      any                `json:"result"`
	Description string             `json:"description"`
	ErrorCode   int                `json:"error_code"`
	Parameters  ResponseParameters `json:"parameters"`
}
