package controller

import (
	"app/internal/database"
	"app/internal/model"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProductByID(context *gin.Context) {
	productID := context.Param("productId")

	var product model.Product
	err := database.Db.QueryRow(
		`SELECT product_id, name, description, interest_rate FROM vtb_products WHERE product_id = $1`, productID,
	).Scan(&product.ProductID, &product.Name, &product.Description, &product.InterestRate)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			context.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve product"})
		}
		return
	}

	context.JSON(http.StatusOK, product)
}

func CreateCustomerLead(context *gin.Context) {
	var leadRequest model.CustomerLeadRequest
	if err := context.BindJSON(&leadRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	var leadID string
	err := database.Db.QueryRow(
		`INSERT INTO vtb_customer_leads (customer_name, contact_info) VALUES ($1, $2) RETURNING lead_id`,
		leadRequest.CustomerName, leadRequest.ContactInfo,
	).Scan(&leadID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create customer lead"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"leadId": leadID})
}

func GetCustomerLeadByID(context *gin.Context) {
	leadID := context.Param("customerLeadId")

	var lead model.CustomerLeadResponse
	err := database.Db.QueryRow(
		`SELECT lead_id, customer_name, contact_info, status FROM vtb_customer_leads WHERE lead_id = $1`, leadID,
	).Scan(&lead.LeadID, &lead.CustomerName, &lead.ContactInfo, &lead.Status)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			context.JSON(http.StatusNotFound, gin.H{"error": "Lead not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve lead"})
		}
		return
	}

	context.JSON(http.StatusOK, lead)
}

func DeleteCustomerLead(context *gin.Context) {
	leadID := context.Param("customerLeadId")

	result, err := database.Db.Exec(`DELETE FROM vtb_customer_leads WHERE lead_id = $1`, leadID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete customer lead"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		context.JSON(http.StatusNotFound, gin.H{"error": "Lead not found"})
		return
	}

	context.JSON(http.StatusNoContent, nil)
}

func CreateProductOffer(context *gin.Context) {
	var offerRequest model.ProductOfferRequest
	if err := context.BindJSON(&offerRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	var offerID string
	err := database.Db.QueryRow(
		`INSERT INTO vtb_product_offers (product_id, lead_id, offer_details) VALUES ($1, $2, $3) RETURNING offer_id`,
		offerRequest.ProductID, offerRequest.LeadID, offerRequest.OfferDetails,
	).Scan(&offerID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product offer"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"offerId": offerID})
}

func GetProductOffers(context *gin.Context) {
	rows, err := database.Db.Query(`SELECT offer_id, product_id, lead_id, offer_details FROM vtb_product_offers`)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve product offers"})
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var offers []model.ProductOffer
	for rows.Next() {
		var offer model.ProductOffer
		if err := rows.Scan(&offer.OfferID, &offer.ProductID, &offer.LeadID, &offer.OfferDetails); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse offer data"})
			return
		}
		offers = append(offers, offer)
	}

	context.JSON(http.StatusOK, gin.H{"offers": offers})
}

func GetProductOfferByID(context *gin.Context) {
	offerID := context.Param("offerId")

	var offer model.ProductOffer
	err := database.Db.QueryRow(
		`SELECT offer_id, product_id, lead_id, offer_details FROM vtb_product_offers WHERE offer_id = $1`, offerID,
	).Scan(&offer.OfferID, &offer.ProductID, &offer.LeadID, &offer.OfferDetails)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			context.JSON(http.StatusNotFound, gin.H{"error": "Offer not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve offer"})
		}
		return
	}

	context.JSON(http.StatusOK, offer)
}

func DeleteProductOffer(context *gin.Context) {
	offerID := context.Param("offerId")

	result, err := database.Db.Exec(`DELETE FROM vtb_product_offers WHERE offer_id = $1`, offerID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product offer"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		context.JSON(http.StatusNotFound, gin.H{"error": "Offer not found"})
		return
	}

	context.JSON(http.StatusNoContent, nil)
}

func CreateProductOfferConsent(context *gin.Context) {
	var consentRequest model.ProductOfferConsentRequest
	if err := context.BindJSON(&consentRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	var consentID string
	err := database.Db.QueryRow(
		`INSERT INTO vtb_product_offer_consents (offer_id, consent_details) VALUES ($1, $2) RETURNING consent_id`,
		consentRequest.OfferID, consentRequest.ConsentDetails,
	).Scan(&consentID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product offer consent"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"consentId": consentID})
}

func GetProductOfferConsentByID(context *gin.Context) {
	consentID := context.Param("consentId")

	var consent model.ProductOfferConsent
	err := database.Db.QueryRow(
		`SELECT consent_id, offer_id, consent_details FROM vtb_product_offer_consents WHERE consent_id = $1`, consentID,
	).Scan(&consent.ConsentID, &consent.OfferID, &consent.ConsentDetails)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			context.JSON(http.StatusNotFound, gin.H{"error": "Consent not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve consent"})
		}
		return
	}

	context.JSON(http.StatusOK, consent)
}

func DeleteProductOfferConsent(context *gin.Context) {
	consentID := context.Param("consentId")

	result, err := database.Db.Exec(`DELETE FROM vtb_product_offer_consents WHERE consent_id = $1`, consentID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product offer consent"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		context.JSON(http.StatusNotFound, gin.H{"error": "Consent not found"})
		return
	}

	context.JSON(http.StatusNoContent, nil)
}

func CreateProductApplication(context *gin.Context) {
	var applicationRequest model.ProductApplicationRequest
	if err := context.BindJSON(&applicationRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	var applicationID string
	err := database.Db.QueryRow(
		`INSERT INTO vtb_product_applications (product_id, applicant_id, application_details) VALUES ($1, $2, $3) RETURNING application_id`,
		applicationRequest.ProductID, applicationRequest.ApplicantID, applicationRequest.ApplicationDetails,
	).Scan(&applicationID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product application"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"applicationId": applicationID})
}

func GetProductApplications(context *gin.Context) {
	rows, err := database.Db.Query(`SELECT application_id, product_id, applicant_id, status FROM vtb_product_applications`)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve product applications"})
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var applications []model.ProductApplication
	for rows.Next() {
		var application model.ProductApplication
		if err := rows.Scan(&application.ApplicationID, &application.ProductID, &application.ApplicantID, &application.Status); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse application data"})
			return
		}
		applications = append(applications, application)
	}

	context.JSON(http.StatusOK, gin.H{"applications": applications})
}

func GetProductApplicationByID(context *gin.Context) {
	applicationID := context.Param("productApplicationId")

	var application model.ProductApplication
	err := database.Db.QueryRow(
		`SELECT application_id, product_id, applicant_id, application_details, status FROM vtb_product_applications WHERE application_id = $1`, applicationID,
	).Scan(&application.ApplicationID, &application.ProductID, &application.ApplicantID, &application.ApplicationDetails, &application.Status)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			context.JSON(http.StatusNotFound, gin.H{"error": "Application not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve application"})
		}
		return
	}

	context.JSON(http.StatusOK, application)
}

func DeleteProductApplication(context *gin.Context) {
	applicationID := context.Param("productApplicationId")

	result, err := database.Db.Exec(`DELETE FROM vtb_product_applications WHERE application_id = $1`, applicationID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product application"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		context.JSON(http.StatusNotFound, gin.H{"error": "Application not found"})
		return
	}

	context.JSON(http.StatusNoContent, nil)
}
