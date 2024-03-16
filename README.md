# Golang Restful API with MongoDB Database
This is hosted on https://golang-mongodb-rest.onrender.com

## Quick Start
1. Clone

```bash
git clone https://github.com/tobiasprima/Golang_MongoDB_REST.git
```

2. Install Dependency

```bash
go mod tidy
```

3. To test with Mongodb locally
-create new `.env` file and add:

```bash
DATABASE_URI = {YOUR_DATABASE_URI}
```

-run on localhost

```bash
go run main.go
```

4. Test Endpoints
on postman

### get products
```bash
http://localhost:8080/products
```

### post products
```bash
http://localhost:8080/products
```
with body
```bash
{
    "name": "{PRODUCT_NAME}",
    "category": "{PRODUCT_CATEGORY}",
    "price": {PRODUCT_PRICE},
    "stock": {PRODUCT_STOCK}
}
```

### patch products' price or stock
```bash
http://localhost:8080/products/{ID}   // copy ID from get/products method
```
with body
```bash
{
   "stock": {PRODUCT_STOCK}
}
```
or
```bash
{
   "price": {PRODUCT_PRICE}
}
```

### get product by id
```bash
http://localhost:8080/products/{ID}   // copy ID from get/products method
```
