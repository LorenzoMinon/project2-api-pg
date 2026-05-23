# project2-api

A REST API built with Go and PostgreSQL to manage products. Data persists in a real database.

## Requirements

- Go 1.25+
- Docker

## Setup

Start the PostgreSQL container:

```bash
docker run --name project2-db \
  -e POSTGRES_USER=admin \
  -e POSTGRES_PASSWORD=admin \
  -e POSTGRES_DB=project2 \
  -p 5433:5432 \
  -d postgres
```

Create the products table:

```bash
docker exec -it project2-db psql -U admin -d project2
```

```sql
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    description TEXT,
    stock INT NOT NULL DEFAULT 0
);
```

## Run

```bash
go run main.go
```

Server starts on `http://localhost:8080`.

## Endpoints

### GET /products
Returns all products.

### GET /products/{id}
Returns a single product by ID.
Returns 404 if not found.

### POST /products
Creates a new product.

**Body:**
```json
{
    "name": "Jeans",
    "price": 100.00,
    "description": "Levis 505",
    "stock": 5
}
```

### PUT /products/{id}
Updates the stock of a product.

**Body:**
```json
{
    "stock": 20
}
```

### DELETE /products/{id}
Deletes a product by ID.
Returns 204 on success, 404 if not found.