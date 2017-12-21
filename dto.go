package main

type ResponseMessage struct {
	SuccessStatus bool              `json:"success"`
	Errors        map[string]string `json:"errors"`
}