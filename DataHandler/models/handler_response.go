package models

type SuccessMessage struct {
	Success		bool		`json:"success"`
	Message		string		`json:"message,omitempty"`
}
