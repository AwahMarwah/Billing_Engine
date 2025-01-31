CREATE TABLE Borrower (
    borrower_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255),
    phone_number VARCHAR(50),
    address TEXT
);

CREATE TABLE Loan (
    loan_id SERIAL PRIMARY KEY,
    borrower_id INT NOT NULL,
    principal_amount DECIMAL(12, 2) NOT NULL,
    interest_rate DECIMAL(5, 2) NOT NULL,
    loan_duration_weeks INT NOT NULL,
    total_amount_due DECIMAL(12, 2) NOT NULL,
    outstanding_balance DECIMAL(12, 2) NOT NULL,
    status VARCHAR(50) CHECK (status IN ('Active', 'Closed')) NOT NULL,
    FOREIGN KEY (borrower_id) REFERENCES Borrower(borrower_id)
);

CREATE TABLE Repayment (
    repayment_id SERIAL PRIMARY KEY,
    loan_id INT NOT NULL,
    payment_date DATE NOT NULL,
    amount_paid DECIMAL(12, 2) NOT NULL,
    payment_status VARCHAR(50) CHECK (payment_status IN ('Paid', 'Missed')) NOT NULL,
    week_number INT NOT NULL,
    FOREIGN KEY (loan_id) REFERENCES Loan(loan_id)
);

CREATE TABLE PaymentHistory (
    payment_history_id SERIAL PRIMARY KEY,
    repayment_id INT NOT NULL,
    payment_date DATE NOT NULL,
    amount_paid DECIMAL(12, 2) NOT NULL,
    status VARCHAR(50) CHECK (status IN ('Paid', 'Missed')) NOT NULL,
    FOREIGN KEY (repayment_id) REFERENCES Repayment(repayment_id)
);

CREATE TABLE Delinquency (
    delinquency_id SERIAL PRIMARY KEY,
    loan_id INT NOT NULL,
    delinquent BOOLEAN NOT NULL,
    date_checked DATE NOT NULL,
    FOREIGN KEY (loan_id) REFERENCES Loan(loan_id)
);
