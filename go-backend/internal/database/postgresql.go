package database

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
	"strconv"
)

var Db *sql.DB

func ConnectDatabase() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error occurred while loading .env file, please check")
	}
	// Read environment variables from .env file
	host := os.Getenv("HOST")
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	user := os.Getenv("POSTGRES_USER")
	dbname := os.Getenv("POSTGRES_DB")
	pass := os.Getenv("POSTGRES_PASSWORD")
	psqlSetup := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		user, pass, host, port, dbname)
	db, errSql := sql.Open("postgres", psqlSetup)
	if errSql != nil {
		fmt.Println("There is an error while connecting to the database: ", errSql)
		panic(errSql)
	} else {
		Db = db
		fmt.Println("Successfully connected to the database!")
	}

	createTablesQuery := `
	CREATE TABLE IF NOT EXISTS vtb_users (
		id SERIAL PRIMARY KEY,  
		name VARCHAR(255) NOT NULL,  
		email VARCHAR(255) NOT NULL, 
		password VARCHAR(255) NOT NULL,
		CONSTRAINT unique_name UNIQUE (name)
	);
	CREATE TABLE vtb_payment_consents (
    	consent_id SERIAL PRIMARY KEY,
    	data JSONB NOT NULL,
    	risk JSONB NOT NULL,
    	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	CREATE TABLE vtb_payments (
	    payment_id SERIAL PRIMARY KEY,
	    consent_id INTEGER REFERENCES vtb_payment_consents(consent_id),
	    data JSONB NOT NULL,
	    risk JSONB NOT NULL,
	    status VARCHAR(50) DEFAULT 'Created',
	    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	CREATE TABLE vtb_account_consents (
	    consent_id SERIAL PRIMARY KEY,
	    data JSONB NOT NULL,
	    risk JSONB,
	    status VARCHAR(50) DEFAULT 'AwaitingAuthorization',
	    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	CREATE TABLE vtb_accounts (
	    account_id SERIAL PRIMARY KEY,
	    status VARCHAR(50) NOT NULL,
	    currency VARCHAR(10) NOT NULL,
	    details JSONB
	);
	CREATE TABLE vtb_balances (
	    balance_id SERIAL PRIMARY KEY,
	    account_id INTEGER REFERENCES vtb_accounts(account_id),
	    amount NUMERIC(15, 2) NOT NULL,
	    currency VARCHAR(10) NOT NULL,
	    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	CREATE TABLE vtb_transactions (
	    transaction_id SERIAL PRIMARY KEY,
	    account_id INTEGER REFERENCES vtb_accounts(account_id),
	    amount NUMERIC(15, 2) NOT NULL,
	    currency VARCHAR(10) NOT NULL,
	    date TIMESTAMP NOT NULL,
	    transaction_reference VARCHAR(100)
	);
	CREATE TABLE vtb_statements (
	    statement_id SERIAL PRIMARY KEY,
	    account_id INTEGER REFERENCES vtb_accounts(account_id),
	    from_date TIMESTAMP NOT NULL,
	    to_date TIMESTAMP NOT NULL,
	    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	CREATE TABLE vtb_vrp_payments_details (
	    payment_id INTEGER REFERENCES vtb_vrp_payments(payment_id),
	    initiation JSONB,
	    instruction JSONB,
	    details JSONB,
	    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	CREATE TABLE vtb_vrp_consents (
	    consent_id SERIAL PRIMARY KEY,
	    data JSONB NOT NULL,
	    risk JSONB,
	    status VARCHAR(50) DEFAULT 'AwaitingAuthorization',
	    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	
	CREATE TABLE vtb_vrp_funds_confirmations (
	    confirmation_id SERIAL PRIMARY KEY,
	    consent_id INTEGER REFERENCES vtb_vrp_consents(consent_id),
	    reference VARCHAR(35),
	    amount NUMERIC(15, 2),
	    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	
	CREATE TABLE vtb_vrp_payments (
	    payment_id SERIAL PRIMARY KEY,
	    consent_id INTEGER REFERENCES vtb_vrp_consents(consent_id),
	    initiation JSONB,
	    instruction JSONB,
	    risk JSONB,
	    status VARCHAR(50) DEFAULT 'Pending',
	    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	CREATE TABLE vtb_products (
	    product_id SERIAL PRIMARY KEY,
	    name VARCHAR(255) NOT NULL,
	    description TEXT,
	    interest_rate NUMERIC(5, 2) NOT NULL
	);
	CREATE TABLE vtb_customer_leads (
	    lead_id SERIAL PRIMARY KEY,
	    customer_name VARCHAR(255) NOT NULL,
	    contact_info TEXT NOT NULL,
	    status VARCHAR(50) DEFAULT 'Pending'
	);
	CREATE TABLE vtb_product_offers (
	    offer_id SERIAL PRIMARY KEY,
	    product_id INTEGER REFERENCES vtb_products(product_id),
	    lead_id INTEGER REFERENCES vtb_customer_leads(lead_id),
	    offer_details TEXT NOT NULL,
	    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	CREATE TABLE vtb_product_offer_consents (
	    consent_id SERIAL PRIMARY KEY,
	    offer_id INTEGER REFERENCES vtb_product_offers(offer_id),
	    consent_details TEXT NOT NULL,
	    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	CREATE TABLE vtb_product_applications (
	    application_id SERIAL PRIMARY KEY,
	    product_id INTEGER REFERENCES vtb_products(product_id),
	    applicant_id INTEGER,
	    application_details TEXT NOT NULL,
	    status VARCHAR(50) DEFAULT 'Pending',
	    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err = Db.Exec(createTablesQuery)
	if err != nil {
		fmt.Println("An error occurred while creating the tables:", err)
		panic(err)
	} else {
		fmt.Println("Tables have been created successfully or already exist")
	}
}
