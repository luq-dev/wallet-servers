CREATE TYPE transaction_status AS ENUM ('PENDING', 'COMPLETE');
CREATE TYPE user_role AS ENUM ('user','admin');

CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        fullname VARCHAR(255) NOT NULL,
        email VARCHAR(255) UNIQUE NOT NULL,
        password VARCHAR(255) NOT NULL,
        role user_role NOT NULL,
        phone_number VARCHAR(20),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS accounts (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id),
    account_name VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS balances (
    id SERIAL PRIMARY KEY,
    account_id INTEGER NOT NULL REFERENCES accounts(id),
    currency VARCHAR(3) NOT NULL,
    amount DECIMAL(18, 2) NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(account_id, currency)
);


CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    from_account_id INTEGER NOT NULL REFERENCES accounts(id),
    to_account_id INTEGER NOT NULL REFERENCES accounts(id),
    currency VARCHAR(3) NOT NULL,
    amount DECIMAL(18, 2) NOT NULL,
    description TEXT,
    status transaction_status NOT NULL,                            -- transaction_state ('PENDING', 'COMPLETE')
    exchange_rate DECIMAL(10, 6),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS account_types (
    id SERIAL PRIMARY KEY,
    type_name VARCHAR(50) NOT NULL,
    currency VARCHAR(3)
);

CREATE INDEX IF NOT EXISTS idx_transactions_from ON transactions(from_account_id);
CREATE INDEX IF NOT EXISTS idx_transactions_to ON transactions(to_account_id);
CREATE INDEX IF NOT EXISTS idx_balances_account ON balances(account_id);
