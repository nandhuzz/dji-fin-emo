CREATE TABLE transactions (
    transaction_id SERIAL PRIMARY KEY,
    amount NUMERIC(15, 2) NOT NULL,
    transaction_type VARCHAR(50) NOT NULL, -- e.g., income, expense, transfer
    account_id INT NOT NULL,
    user_id INT NOT NULL,
    transaction_date DATE NOT NULL,
    transaction_time TIME NOT NULL,
    registration TIMESTAMP DEFAULT NOW(),
    description TEXT,
    labels TEXT[], -- PostgreSQL array of text for list of labels
    payee VARCHAR(255),
    category_id INT
);