CREATE TABLE IF NOT EXISTS account_types (
    id SERIAL PRIMARY KEY,
    type_name VARCHAR(50) NOT NULL,
    currency VARCHAR(3)
);

ALTER TABLE accounts ADD COLUMN IF NOT EXISTS account_type INTEGER REFERENCES account_types(id) ON DELETE RESTRICT; 
