# Go REST API Application

## Introduction

This application is a REST API server built with Go (Golang) that manages users and financial transactions. It allows users to upload transactions via CSV files, store data in a PostgreSQL database, and generate transaction summaries to be sent via email. The application is designed to be modular and extensible.

## Features

- **Transaction Upload**: Allows users to upload transactions via CSV files.
- **User Management**: Stores user information in a PostgreSQL database.
- **Summary Generation**: Calculates transaction summaries and sends reports via email.

## Requirements

- **Go**: This application is built with Go 1.22.3. Make sure you have Go installed if you want to run the code directly.

## Installation

### Clone the Repository

First, clone the repository:

```bash
git clone https://github.com/juanrod1028/Transactions
cd Transactions
```
- In the main.go file, replace the empty fields for COMPANY_EMAIL and COMPANY_EMAIL_PASS.
- Adjust the host, port, user, password, and dbname settings to connect to PostgreSQL.
Then, run the make file:
```bash
make run
```
# EndPoints
## POST: transactions
This endpoint requires:
- A CSV file like the following
- ![image](https://github.com/user-attachments/assets/f01b0894-9baf-4467-b48f-467338808b86)
- The email of the person to whom the email will be sent and their identification.
![image](https://github.com/user-attachments/assets/8a39a225-ff28-42ff-9a95-507e959af3ae)

## GET: user/transactions/{id}
-This endpoint requires the ID of the person who previously registered their transactions.
![image](https://github.com/user-attachments/assets/c0f42be2-4f97-4aad-bf30-91744e710c93)
