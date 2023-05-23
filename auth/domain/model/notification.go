package model

type VerificationEmailNotification struct {
	Recipient        string `json:"recipient"`
	VerificationCode string `json:"code"`
}
