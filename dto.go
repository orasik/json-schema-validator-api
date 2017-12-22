package main

// ResponseMessage This is the template for response message
// that will be json encoded
type ResponseMessage struct {
	SuccessStatus bool              `json:"success"`
	Errors        map[string]string `json:"errors"`
}
