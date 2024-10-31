package model

// VRPControlParameters представляет параметры управления для ППДС
type VRPControlParameters struct {
	ValidityPeriod struct {
		StartDateTime string `json:"startDateTime"`
		EndDateTime   string `json:"endDateTime"`
	} `json:"validityPeriod"`
	MaximumIndividualAmount struct {
		Amount   string `json:"amount"`
		Currency string `json:"currency"`
	} `json:"maximumIndividualAmount"`
}

// VRPInitiation представляет данные инициации ППДС
type VRPInitiation struct {
	InstructionIdentification string `json:"instructionIdentification"`
	EndToEndIdentification    string `json:"endToEndIdentification"`
	InstructedAmount          struct {
		Amount   string `json:"amount"`
		Currency string `json:"currency"`
	} `json:"instructedAmount"`
	CreditorAccount struct {
		SchemeName     string `json:"schemeName"`
		Identification string `json:"identification"`
		Name           string `json:"name"`
	} `json:"creditorAccount"`
}

// VRPInstruction представляет инструкции для ППДС
type VRPInstruction struct {
	RemittanceInformation struct {
		Unstructured string `json:"unstructured"`
		Reference    string `json:"reference"`
	} `json:"remittanceInformation"`
}

// VRPConsentRequest представляет запрос на создание согласия
type VRPConsentRequest struct {
	Data struct {
		ControlParameters VRPControlParameters `json:"ControlParameters"`
		Initiation        VRPInitiation        `json:"Initiation"`
	} `json:"Data"`
	Risk Risk `json:"Risk"`
}

// VRPConsentResponse представляет ответ при создании согласия
type VRPConsentResponse struct {
	ConsentID string            `json:"consentId"`
	Data      VRPConsentRequest `json:"data"`
	Status    string            `json:"status"`
	Risk      string            `json:"risk"`
	CreatedAt string            `json:"createdAt"`
}

// VRPFundsConfirmationRequest представляет запрос на подтверждение средств
type VRPFundsConfirmationRequest struct {
	Data struct {
		ConsentID        string `json:"consentId"`
		Reference        string `json:"reference"`
		InstructedAmount string `json:"InstructedAmount"`
	} `json:"Data"`
}

// VRPRequest представляет запрос на инициацию платежа
type VRPRequest struct {
	Data struct {
		ConsentID   string         `json:"consentId"`
		Initiation  VRPInitiation  `json:"Initiation"`
		Instruction VRPInstruction `json:"Instruction"`
	} `json:"Data"`
	Risk Risk `json:"Risk"`
}

// VRPResponse представляет ответ при инициации платежа
type VRPResponse struct {
	PaymentID   string         `json:"paymentId"`
	ConsentID   string         `json:"consentId"`
	Initiation  VRPInitiation  `json:"Initiation"`
	Instruction VRPInstruction `json:"Instruction"`
	Status      string         `json:"status"`
	CreatedAt   string         `json:"createdAt"`
}

// VRPPaymentDetailsResponse представляет детальную информацию о платеже по ППДС
type VRPPaymentDetailsResponse struct {
	PaymentID   string         `json:"paymentId"`
	Initiation  VRPInitiation  `json:"initiation"`
	Instruction VRPInstruction `json:"instruction"`
	Details     PaymentDetails `json:"details"`
}

// PaymentDetails представляет дополнительные детали платежа
type PaymentDetails struct {
	Status            string `json:"status"`
	ExecutionDateTime string `json:"executionDateTime"`
	SettlementAmount  struct {
		Amount   string `json:"amount"`
		Currency string `json:"currency"`
	} `json:"settlementAmount"`
	Charges []Charge `json:"charges"`
}

// Charge представляет информацию о сборах, связанных с платежом
type Charge struct {
	Amount struct {
		Amount   string `json:"amount"`
		Currency string `json:"currency"`
	} `json:"amount"`
	Type string `json:"type"`
}
