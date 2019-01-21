# Shopify Summer 2019 Internship Challenge

### Running the Service + Database with Docker

#### Starting the database
`cd database`  
`docker build -t shopify_db .`  
`docker run --name shopify_db_inst -p 5432:5432 -i shopify_db`  

#### Starting the service with Docker
`docker build -t shopify_service .`  
`docker run --name shopify_service_inst -p 3000:3000 -d -i shopify_service`  

#### Starting the service outside docker with Go
Install Dep on Mac: `brew install dep`  
Install Dep on Other: `curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh`  

Get dependencies: `dep ensure`  
Run service: `go build -o shop && ./shop`  

Access service at `localhost:3000`  

### Endpoint Documentation

| Route | Type | Example Payload | Example Response |
|---|---|---|---|
| /health | GET | None | {"status":"OK"} |
| /v1/products | GET | None | Array[ProductModel] |
| /v1/products?onlyAvailable=true | GET | None | Array[ProductModel] |
| /v1/products/:id | GET | None | [ProductModel] |
| /v1/products | POST | [PostProductModel] | [ProductModel] |
| /v1/products | DELETE | None | None |
| /v1/carts/:cartId | GET | None | Array[ProductModel] |
| /v1/carts/:cartId?productId=# | POST | None | Array[ProductModel] |
| /v1/carts?productId=# | POST | None | Array[ProductModel] |
| /v1/carts/:cartId | DELETE | None | None |
| /v1/checkout/:cartId | POST | None | None |

### Payload Definitions
#### ProductModel
```json
{
    "id": 1,
    "title": "Regular T-Shirt",
    "price": 19.99,
    "inventoryCount": 1
}
```

#### PostProductModel
```json
{
    "title": "Regular Crewneck",
    "price": 34.99
}
```
