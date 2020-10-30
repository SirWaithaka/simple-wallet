package errors

// ApiErrorResponse properties returned by the REST api for requests that
// are not successful
type ApiErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}
