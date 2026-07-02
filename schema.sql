CREATE TYPE transaction_state AS ENUM ('PENDING', 'COMPLETE');
CREATE TYPE user_role AS ENUM ('user','admin');

-- TABLES

CREATE TABLE IF NOT EXISTS users (
        id BIGSERIAL PRIMARY KEY,
        fullname VARCHAR(255) NOT NULL,
        email VARCHAR(255) UNIQUE NOT NULL,
        password VARCHAR(255) NOT NULL,
        user_type user_role NOT NULL,
        phone_number VARCHAR(20),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS accounts (
    id BIGSERIAL PRIMARY KEY,
    account_number VARCHAR(255) UNIQUE NOT NULL,
    user_id INTEGER NOT NULL REFERENCES users(id),
    account_type INTEGER NOT NULL REFERENCES account_types(id),
    currency VARCHAR(3),
    account_name VARCHAR(255),
    PIN VARCHAR(4),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS balances (
    id BIGSERIAL PRIMARY KEY,
    account_id INTEGER NOT NULL REFERENCES accounts(id),
    currency VARCHAR(3) NOT NULL,
    amount DECIMAL(18, 2) NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(account_id, currency)
);

CREATE TABLE IF NOT EXISTS transactions (
    id BIGSERIAL PRIMARY KEY,
    transaction_id VARCHAR(16),     -- need a keying system
    from_account_number VARCHAR(255),
    to_account_number VARCHAR(255),
    currency VARCHAR(3) NOT NULL,
    amount DECIMAL(18, 2) NOT NULL,
    destination_bank VARCHAR(255),
    details TEXT,
    transaction_status transaction_state NOT NULL,                            -- transaction_state ('PENDING', 'COMPLETE')
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
)PARTITION BY RANGE (created_at);


CREATE TABLE IF NOT EXISTS messages (
    id BIGSERIAL PRIMARY KEY,
    receiver VARCHAR(255),  -- user Email
    message_content TEXT NOT NULL,
    is_read BOOLEAN DEFAULT FALSE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS account_types (
    id SERIAL PRIMARY KEY,
    type_name VARCHAR(50) NOT NULL
);


CREATE TABLE IF NOT EXISTS exchange_rates {
    business_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    currency VARCHAR(3) NOT NULL,
    rate DECIMAL(18, 2) NOT NULL, -- AGAINST USD
}

-- VIEWS


CREATE VIEW user_account_details AS 
        SELECT a.user_id, a.account_id, a.account_name, t.type_name as account_type 
        FROM accounts a 
        JOIN account_types t 
        ON a.account_type = t.type_name;

-- INDICES

CREATE INDEX IF NOT EXISTS idx_transactions_from ON transactions(from_account_number);
CREATE INDEX IF NOT EXISTS idx_transactions_to ON transactions(to_account_number);
CREATE INDEX IF NOT EXISTS idx_balances_account ON balances(account_id);

-- FUNCTIONS

