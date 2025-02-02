# Billing_Engine
This is Test Project to simulate billing process, we will build it using golang with GIN, postgresql (with GORM),
This is a test project to simulate the billing process. We will build it using Golang with GIN, PostgreSQL (with GORM).

Features
The engine provides the ability to:

Generate a loan schedule for a given loan (when and how much should be paid).
Track the outstanding amount for a given loan.
Determine whether the borrower is delinquent or not.
Setup
1. Run Database Migration
Before running the application, ensure the database schema is properly set up by running the migration command:

make run_db_migrate_up
This command will execute all the necessary migrations to create the required tables in PostgreSQL.
