package model

// Product представляет информацию о кредитном продукте
type Product struct {
	ProductID    string  `json:"productId"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	InterestRate float64 `json:"interestRate"`
}

// CustomerLeadRequest представляет запрос на создание лидогенерации
type CustomerLeadRequest struct {
	CustomerName string `json:"customerName"`
	ContactInfo  string `json:"contactInfo"`
}

// CustomerLeadResponse представляет ответ на запрос о лидогенерации
type CustomerLeadResponse struct {
	LeadID       string `json:"leadId"`
	CustomerName string `json:"customerName"`
	ContactInfo  string `json:"contactInfo"`
	Status       string `json:"status"`
}

// ProductOffer представляет предложение по продукту
type ProductOffer struct {
	OfferID      string `json:"offerId"`
	ProductID    string `json:"productId"`
	LeadID       string `json:"leadId"`
	OfferDetails string `json:"offerDetails"`
}

// ProductOfferRequest представляет запрос на создание предложения по продукту
type ProductOfferRequest struct {
	ProductID    string `json:"productId"`
	LeadID       string `json:"leadId"`
	OfferDetails string `json:"offerDetails"`
}

// ProductOffer представляет предложение по продукту
type ProductOffer struct {
	OfferID      string `json:"offerId"`
	ProductID    string `json:"productId"`
	LeadID       string `json:"leadId"`
	OfferDetails string `json:"offerDetails"`
}

// ProductOfferConsentRequest представляет запрос на создание согласия на предложение по продукту
type ProductOfferConsentRequest struct {
	OfferID        string `json:"offerId"`
	ConsentDetails string `json:"consentDetails"`
}

// ProductOfferConsent представляет согласие на предложение по продукту
type ProductOfferConsent struct {
	ConsentID      string `json:"consentId"`
	OfferID        string `json:"offerId"`
	ConsentDetails string `json:"consentDetails"`
}

// ProductApplicationRequest представляет запрос на создание заявления по продукту
type ProductApplicationRequest struct {
	ProductID          string `json:"productId"`
	ApplicantID        string `json:"applicantId"`
	ApplicationDetails string `json:"applicationDetails"`
}

// ProductApplication представляет информацию о заявлении по продукту
type ProductApplication struct {
	ApplicationID      string `json:"applicationId"`
	ProductID          string `json:"productId"`
	ApplicantID        string `json:"applicantId"`
	ApplicationDetails string `json:"applicationDetails"`
	Status             string `json:"status"`
}
