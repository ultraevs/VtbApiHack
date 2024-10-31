package model

import "time"

type ConsentRequest struct {
	Data struct {
		Initiation struct {
			InstructionIdentification string  `json:"instructionIdentification"`
			EndToEndIdentification    string  `json:"endToEndIdentification"`
			InstructedAmount          Amount  `json:"instructedAmount"`
			CreditorAccount           Account `json:"creditorAccount"`
		} `json:"initiation"`
	} `json:"data"`
	Risk Risk `json:"risk"`
}

type Amount struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type Account struct {
	SchemeName     string `json:"schemeName"`
	Identification string `json:"identification"`
}

type Risk struct {
	PaymentContextCode string `json:"paymentContextCode"`
}

type ConsentResponse struct {
	ConsentID string    `json:"consentId"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}

type PaymentRequest struct {
	Data struct {
		ConsentID  string `json:"consentId"`
		Initiation struct {
			InstructionIdentification string  `json:"instructionIdentification"`
			EndToEndIdentification    string  `json:"endToEndIdentification"`
			InstructedAmount          Amount  `json:"instructedAmount"`
			CreditorAccount           Account `json:"creditorAccount"`
		} `json:"initiation"`
	} `json:"data"`
	Risk Risk `json:"risk"`
}

type PaymentResponse struct {
	PaymentID string    `json:"paymentId"`
	ConsentID string    `json:"consentId"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}
