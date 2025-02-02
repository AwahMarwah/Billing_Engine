# Billing_Engine  

This is a test project to simulate the billing process. We will build it using **Golang** with **GIN**, **PostgreSQL** (with **GORM**).  

## Features  
The engine provides the ability to:  

- Generate a **loan schedule** for a given loan (when and how much should be paid).  
- Track the **outstanding amount** for a given loan.  
- Determine whether the **borrower is delinquent** or not.  

## Setup  

### 1. Run Database Migration  
Before running the application, ensure the database schema is properly set up by running the migration command:  

```sh
make run_db_migrate_up
```

### 2. Import Postman Collection
To simplify API testing, you can import the Postman collection provided in the documents folder. Follow these steps:

1. Open Postman.
2. Click on Import in the top-left corner.
3. Select the Postman Collection JSON file located in the documents folder.
4. Click Import to load the API collection.
5. Use the pre-configured requests to test the API endpoints.
