CREATE TABLE billing_schedules (
    id SERIAL PRIMARY KEY,
    loan_id INT NOT NULL,
    week_number INT NOT NULL,
    due_date DATE NOT NULL,
    amount_due DECIMAL(12,2) NOT NULL,
    status VARCHAR(50) CHECK (status IN ('Pending', 'Paid', 'Missed')) NOT NULL DEFAULT 'Pending',
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (loan_id) REFERENCES Loans(id)
);