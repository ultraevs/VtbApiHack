package model

import "encoding/json"

// AccountConsentRequest представляет запрос на создание согласия
type AccountConsentRequest struct {
	Data struct {
		Permissions             []string `json:"permissions"`
		ExpirationDateTime      string   `json:"expirationDateTime"`
		TransactionFromDateTime string   `json:"transactionFromDateTime"`
		TransactionToDateTime   string   `json:"transactionToDateTime"`
	} `json:"data"`
	Risk struct{} `json:"risk"`
}

// AccountConsentResponse представляет ответ с деталями согласия для API по аккаунтам
type AccountConsentResponse struct {
	ConsentID string                `json:"consentId"`
	Data      AccountConsentRequest `json:"data"`
	Status    string                `json:"status"`
	Risk      string                `json:"risk"`
	CreatedAt string                `json:"createdAt"`
}

// AccountResponse представляет базовую информацию о счете
type AccountResponse struct {
	AccountID string `json:"accountId"`
	Status    string `json:"status"`
	Currency  string `json:"currency"`
}

// AccountDetailResponse содержит детализированную информацию о счете
type AccountDetailResponse struct {
	AccountID string `json:"accountId"`
	Status    string `json:"status"`
	Currency  string `json:"currency"`
	Details   struct {
		SchemeName     string `json:"schemeName"`
		Identification string `json:"identification"`
		Name           string `json:"name"`
	} `json:"details"`
}

// BalanceResponse содержит информацию о балансе счета
type BalanceResponse struct {
	AccountID string `json:"accountId"`
	Amount    string `json:"amount"`
	Currency  string `json:"currency"`
}

// TransactionResponse содержит информацию о транзакции
type TransactionResponse struct {
	TransactionID string `json:"transactionId"`
	AccountID     string `json:"accountId"`
	Amount        string `json:"amount"`
	Currency      string `json:"currency"`
	Date          string `json:"date"`
}

// StatementRequest представляет запрос на создание выписки
type StatementRequest struct {
	FromDate string `json:"fromDate"`
	ToDate   string `json:"toDate"`
}

// StatementResponse содержит информацию о выписке
type StatementResponse struct {
	StatementID string `json:"statementId"`
	AccountID   string `json:"accountId"`
	FromDate    string `json:"fromDate"`
	ToDate      string `json:"toDate"`
}

// RetrievalGrantResponse представляет ответ для ресурса извлечения согласия
type RetrievalGrantResponse struct {
	ConsentID string          `json:"consentId"`
	Data      json.RawMessage `json:"data"`
}
