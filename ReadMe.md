# Inventory Management API

This is a simple API built with Go. It lets you register, log in, create inventory items, restock items, and view restock history. JWT is used for authentication.

## How to Run

1. Clone the repository.
2. Run `go get` to install dependencies.
3. Run `go run main.go`.
4. The SQLite database file `inventory_mg.db` is created automatically.

## API Endpoints

### Register

Create a new admin account
```bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"email": "admin@example.com", "password": "pass123"}'
```

### Login

Login to the admin account
```bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"email": "admin@example.com", "password": "pass123"}'
```

### Create Inventory Item (JWT auth is required)

```bash
curl -X POST http://localhost:8080/inventory/ \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer {your token}" \
  -d '{"name": "Widget", "description": "A useful widget", "quantity": 50}'
```

### Restock an Item (JWT auth is required))

Restock an existing item. The validation is that the amount must be between 10 and 1000, and an item can only be restocked 3 times in 24 hours.
```bash
curl -X POST http://localhost:8080/inventory/1/restock \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer {your token}" \
  -d '{"amount": 100}'
```

### View Restock History (JWT auth is required))

```bash
curl -X GET http://localhost:8080/inventory/1/restocks \
  -H "Authorization: Bearer {your token}"
```
