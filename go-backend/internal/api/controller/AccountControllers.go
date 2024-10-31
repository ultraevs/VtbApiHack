package controller

import (
	"app/internal/database"
	"app/internal/model"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateAccountAccessConsent создает согласие на доступ к счету
func CreateAccountAccessConsent(context *gin.Context) {
	var consentRequest model.ConsentRequest
	if err := context.BindJSON(&consentRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	var consentID string
	err := database.Db.QueryRow(
		`INSERT INTO vtb_account_consents (data, risk) VALUES ($1, $2) RETURNING consent_id`,
		consentRequest.Data, consentRequest.Risk,
	).Scan(&consentID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create account consent"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"consentId": consentID})
}

func GetAccountAccessConsent(context *gin.Context) {
	consentID := context.Param("consentId")

	var consent model.AccountConsentResponse
	err := database.Db.QueryRow(
		`SELECT consent_id, data, risk, status, created_at FROM vtb_account_consents WHERE consent_id = $1`, consentID,
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

func DeleteAccountAccessConsent(context *gin.Context) {
	consentID := context.Param("consentId")

	result, err := database.Db.Exec(`DELETE FROM vtb_account_consents WHERE consent_id = $1`, consentID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete consent"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		context.JSON(http.StatusNotFound, gin.H{"error": "Consent not found"})
		return
	}

	context.JSON(http.StatusNoContent, nil)
}

func GetAccountAccessRetrievalGrant(context *gin.Context) {
	consentID := context.Param("consentId")

	var grant model.RetrievalGrantResponse
	err := database.Db.QueryRow(
		`SELECT retrieval_grant_id, data FROM vtb_retrieval_grants WHERE consent_id = $1`, consentID,
	).Scan(&grant.ConsentID, &grant.Data)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			context.JSON(http.StatusNotFound, gin.H{"error": "Retrieval grant not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve grant"})
		}
		return
	}

	context.JSON(http.StatusOK, grant)
}

func GetAccounts(context *gin.Context) {
	rows, err := database.Db.Query(`SELECT account_id, status, currency FROM vtb_accounts`)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve accounts"})
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var accounts []model.AccountResponse
	for rows.Next() {
		var account model.AccountResponse
		if err := rows.Scan(&account.AccountID, &account.Status, &account.Currency); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse account data"})
			return
		}
		accounts = append(accounts, account)
	}

	context.JSON(http.StatusOK, gin.H{"accounts": accounts})
}

func GetAccount(context *gin.Context) {
	accountID := context.Param("accountId")

	var account model.AccountDetailResponse
	err := database.Db.QueryRow(
		`SELECT account_id, status, currency, details FROM vtb_accounts WHERE account_id = $1`, accountID,
	).Scan(&account.AccountID, &account.Status, &account.Currency, &account.Details)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			context.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve account"})
		}
		return
	}

	context.JSON(http.StatusOK, account)
}

func GetAccountBalance(context *gin.Context) {
	accountID := context.Param("accountId")

	var balance model.BalanceResponse
	err := database.Db.QueryRow(
		`SELECT account_id, amount, currency FROM vtb_balances WHERE account_id = $1`, accountID,
	).Scan(&balance.AccountID, &balance.Amount, &balance.Currency)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			context.JSON(http.StatusNotFound, gin.H{"error": "Balance not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve balance"})
		}
		return
	}

	context.JSON(http.StatusOK, balance)
}

func GetBalances(context *gin.Context) {
	rows, err := database.Db.Query(`SELECT account_id, amount, currency FROM vtb_balances`)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve balances"})
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var balances []model.BalanceResponse
	for rows.Next() {
		var balance model.BalanceResponse
		if err := rows.Scan(&balance.AccountID, &balance.Amount, &balance.Currency); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse balance data"})
			return
		}
		balances = append(balances, balance)
	}

	context.JSON(http.StatusOK, gin.H{"balances": balances})
}

func GetAccountTransactions(context *gin.Context) {
	accountID := context.Param("accountId")

	rows, err := database.Db.Query(`SELECT transaction_id, amount, currency, date FROM vtb_transactions WHERE account_id = $1`, accountID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve transactions"})
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var transactions []model.TransactionResponse
	for rows.Next() {
		var transaction model.TransactionResponse
		if err := rows.Scan(&transaction.TransactionID, &transaction.Amount, &transaction.Currency, &transaction.Date); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse transaction data"})
			return
		}
		transactions = append(transactions, transaction)
	}

	context.JSON(http.StatusOK, gin.H{"transactions": transactions})
}

func GetTransactions(context *gin.Context) {
	rows, err := database.Db.Query(`SELECT transaction_id, account_id, amount, currency, date FROM vtb_transactions`)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve transactions"})
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var transactions []model.TransactionResponse
	for rows.Next() {
		var transaction model.TransactionResponse
		if err := rows.Scan(&transaction.TransactionID, &transaction.AccountID, &transaction.Amount, &transaction.Currency, &transaction.Date); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse transaction data"})
			return
		}
		transactions = append(transactions, transaction)
	}

	context.JSON(http.StatusOK, gin.H{"transactions": transactions})
}

func CreateStatement(context *gin.Context) {
	accountID := context.Param("accountId")

	var statementRequest model.StatementRequest
	if err := context.BindJSON(&statementRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	var statementID string
	err := database.Db.QueryRow(
		`INSERT INTO vtb_statements (account_id, from_date, to_date) VALUES ($1, $2, $3) RETURNING statement_id`,
		accountID, statementRequest.FromDate, statementRequest.ToDate,
	).Scan(&statementID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create statement"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"statementId": statementID})
}

func GetStatementById(context *gin.Context) {
	accountID := context.Param("accountId")
	statementID := context.Param("statementId")

	var statement model.StatementResponse
	err := database.Db.QueryRow(
		`SELECT statement_id, account_id, from_date, to_date FROM vtb_statements WHERE statement_id = $1 AND account_id = $2`,
		statementID, accountID,
	).Scan(&statement.StatementID, &statement.AccountID, &statement.FromDate, &statement.ToDate)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			context.JSON(http.StatusNotFound, gin.H{"error": "Statement not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve statement"})
		}
		return
	}

	context.JSON(http.StatusOK, statement)
}

func GetAllStatements(context *gin.Context) {
	rows, err := database.Db.Query(`SELECT statement_id, account_id, from_date, to_date FROM vtb_statements`)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve statements"})
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var statements []model.StatementResponse
	for rows.Next() {
		var statement model.StatementResponse
		if err := rows.Scan(&statement.StatementID, &statement.AccountID, &statement.FromDate, &statement.ToDate); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse statement data"})
			return
		}
		statements = append(statements, statement)
	}

	context.JSON(http.StatusOK, gin.H{"statements": statements})
}
