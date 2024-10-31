package controller

import (
	"app/internal/database"
	"app/internal/model"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateVRPConsent создает согласие на инициирование ППДС
func CreateVRPConsent(context *gin.Context) {
	var consentRequest model.VRPConsentRequest
	if err := context.BindJSON(&consentRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	var consentID string
	err := database.Db.QueryRow(
		`INSERT INTO vtb_vrp_consents (data, risk) VALUES ($1, $2) RETURNING consent_id`,
		consentRequest.Data, consentRequest.Risk,
	).Scan(&consentID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create VRP consent"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"consentId": consentID})
}

func GetVRPConsent(context *gin.Context) {
	consentID := context.Param("consentId")

	var consent model.VRPConsentResponse
	err := database.Db.QueryRow(
		`SELECT consent_id, data, risk, status, created_at FROM vtb_vrp_consents WHERE consent_id = $1`, consentID,
	).Scan(&consent.ConsentID, &consent.Data, &consent.Risk, &consent.Status, &consent.CreatedAt)

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

func DeleteVRPConsent(context *gin.Context) {
	consentID := context.Param("consentId")

	result, err := database.Db.Exec(`DELETE FROM vtb_vrp_consents WHERE consent_id = $1`, consentID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete VRP consent"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		context.JSON(http.StatusNotFound, gin.H{"error": "Consent not found"})
		return
	}

	context.JSON(http.StatusNoContent, nil)
}

func CreateVRPFundsConfirmation(context *gin.Context) {
	consentID := context.Param("consentId")

	var fundsRequest model.VRPFundsConfirmationRequest
	if err := context.BindJSON(&fundsRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	var confirmationID string
	err := database.Db.QueryRow(
		`INSERT INTO vtb_vrp_funds_confirmations (consent_id, reference, amount) VALUES ($1, $2, $3) RETURNING confirmation_id`,
		consentID, fundsRequest.Data.Reference, fundsRequest.Data.InstructedAmount,
	).Scan(&confirmationID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to confirm funds"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"confirmationId": confirmationID})
}

func CreateVRPPayment(context *gin.Context) {
	var paymentRequest model.VRPRequest
	if err := context.BindJSON(&paymentRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	var paymentID string
	err := database.Db.QueryRow(
		`INSERT INTO vtb_vrp_payments (consent_id, initiation, instruction, risk) VALUES ($1, $2, $3, $4) RETURNING payment_id`,
		paymentRequest.Data.ConsentID, paymentRequest.Data.Initiation, paymentRequest.Data.Instruction, paymentRequest.Risk,
	).Scan(&paymentID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initiate payment"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"paymentId": paymentID})
}

func GetVRPPayment(context *gin.Context) {
	paymentID := context.Param("VRPId")

	var payment model.VRPResponse
	err := database.Db.QueryRow(
		`SELECT payment_id, consent_id, initiation, instruction, status, created_at FROM vtb_vrp_payments WHERE payment_id = $1`,
		paymentID,
	).Scan(&payment.PaymentID, &payment.ConsentID, &payment.Initiation, &payment.Instruction, &payment.Status, &payment.CreatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			context.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve payment"})
		}
		return
	}

	context.JSON(http.StatusOK, payment)
}

func GetVRPPaymentDetails(context *gin.Context) {
	paymentID := context.Param("VRPId")

	var paymentDetails model.VRPPaymentDetailsResponse
	err := database.Db.QueryRow(
		`SELECT payment_id, initiation, instruction, details FROM vtb_vrp_payments_details WHERE payment_id = $1`,
		paymentID,
	).Scan(&paymentDetails.PaymentID, &paymentDetails.Initiation, &paymentDetails.Instruction, &paymentDetails.Details)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			context.JSON(http.StatusNotFound, gin.H{"error": "Payment details not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve payment details"})
		}
		return
	}

	context.JSON(http.StatusOK, paymentDetails)
}
