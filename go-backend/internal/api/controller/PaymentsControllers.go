package controller

import (
	"app/internal/database"
	"app/internal/model"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// CreatePaymentConsent Создание ресурса согласия на проведение платежа
// @Summary Создать согласие на платеж
// @Description Создает ресурс согласия на стороне ППУ для проведения платежа без предварительной идентификации пользователя.
// @Accept json
// @Produce json
// @Param request body model.ConsentRequest true "Запрос на создание согласия на платеж"
// @Success 201 {object} model.CodeResponse "Согласие создано"
// @Failure 400 {object} model.ErrorResponse "Некорректный запрос"
// @Failure 500 {object} model.ErrorResponse "Ошибка сервера"
// @Tags Payments
// @Router /v1/payment-consents [post]
func CreatePaymentConsent(context *gin.Context) {
	// Структура для парсинга тела запроса
	var request model.ConsentRequest
	if err := context.BindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Вставка согласия в базу данных
	var consentID string
	err := database.Db.QueryRow(
		`INSERT INTO vtb_payment_consents (data, risk) VALUES ($1, $2) RETURNING consent_id`,
		request.Data, request.Risk,
	).Scan(&consentID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create consent"})
		return
	}

	// Формируем ответ
	response := model.ConsentResponse{
		ConsentID: consentID,
		Status:    "Created",
		CreatedAt: time.Now(),
	}

	context.JSON(http.StatusCreated, response)
}

// GetPaymentConsent Получение ресурса согласия
// @Summary Получить согласие на платеж
// @Description Возвращает детали согласия на платеж по его идентификатору
// @Accept json
// @Produce json
// @Param consentId path string true "ID согласия"
// @Success 200 {object} model.ConsentResponse "Детали согласия"
// @Failure 400 {object} model.ErrorResponse "Некорректный запрос"
// @Failure 404 {object} model.ErrorResponse "Согласие не найдено"
// @Tags Payments
// @Router /v1/payment-consents/{consentId} [get]
func GetPaymentConsent(context *gin.Context) {
	consentID := context.Param("consentId")

	var consent model.ConsentResponse
	err := database.Db.QueryRow(
		`SELECT consent_id, created_at, status FROM vtb_payment_consents WHERE consent_id = $1`,
		consentID,
	).Scan(&consent.ConsentID, &consent.CreatedAt, &consent.Status)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			context.JSON(http.StatusNotFound, gin.H{"error": "Consent not found"})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve consent"})
		return
	}

	context.JSON(http.StatusOK, consent)
}

// CreatePayment Инициирование платежа
// @Summary Инициировать платеж
// @Description Создает новый платеж на основании согласия
// @Accept json
// @Produce json
// @Param request body model.PaymentRequest true "Запрос на создание платежа"
// @Success 201 {object} model.PaymentResponse "Платеж успешно инициирован"
// @Failure 400 {object} model.ErrorResponse "Некорректный запрос"
// @Failure 500 {object} model.ErrorResponse "Ошибка сервера"
// @Tags Payments
// @Router /v1/payments [post]
func CreatePayment(context *gin.Context) {
	var request model.PaymentRequest
	if err := context.BindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Проверка существования согласия
	var consentExists bool
	err := database.Db.QueryRow(
		`SELECT EXISTS(SELECT 1 FROM vtb_payment_consents WHERE consent_id = $1)`,
		request.Data.ConsentID,
	).Scan(&consentExists)

	if err != nil || !consentExists {
		context.JSON(http.StatusNotFound, gin.H{"error": "Consent not found"})
		return
	}

	// Вставка платежа в базу данных
	var paymentID string
	err = database.Db.QueryRow(
		`INSERT INTO vtb_payments (consent_id, data, risk) VALUES ($1, $2, $3) RETURNING payment_id`,
		request.Data.ConsentID, request.Data, request.Risk,
	).Scan(&paymentID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initiate payment"})
		return
	}

	response := model.PaymentResponse{
		PaymentID: paymentID,
		Status:    "Created",
		CreatedAt: time.Now(),
	}

	context.JSON(http.StatusCreated, response)
}

// GetPayment Получение статуса платежа
// @Summary Получить статус платежа
// @Description Возвращает статус платежа по его идентификатору
// @Accept json
// @Produce json
// @Param paymentId path string true "ID платежа"
// @Success 200 {object} model.PaymentResponse "Статус платежа"
// @Failure 404 {object} model.ErrorResponse "Платеж не найден"
// @Tags Payments
// @Router /v1/payments/{paymentId} [get]
func GetPayment(context *gin.Context) {
	paymentID := context.Param("paymentId")

	var payment model.PaymentResponse
	err := database.Db.QueryRow(
		`SELECT payment_id, consent_id, created_at, status FROM vtb_payments WHERE payment_id = $1`,
		paymentID,
	).Scan(&payment.PaymentID, &payment.ConsentID, &payment.CreatedAt, &payment.Status)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			context.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve payment"})
		return
	}

	context.JSON(http.StatusOK, payment)
}
