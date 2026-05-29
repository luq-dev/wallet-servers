CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        fullname VARCHAR(255) NOT NULL,
        email VARCHAR(255) UNIQUE NOT NULL,
        password VARCHAR(255) NOT NULL,
        phone_number VARCHAR(20),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

    CREATE TABLE IF NOT EXISTS accounts (
        id SERIAL PRIMARY KEY,
        user_id INTEGER NOT NULL REFERENCES users(id),
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
        status ENUM('PENDING', 'COMPLETED') NOT NULL,
        exchange_rate DECIMAL(10, 6),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

    CREATE INDEX IF NOT EXISTS idx_transactions_from ON transactions(from_account_id);
    CREATE INDEX IF NOT EXISTS idx_transactions_to ON transactions(to_account_id);
    CREATE INDEX IF NOT EXISTS idx_balances_account ON balances(account_id);