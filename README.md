# Shopify Summer 2019 Internship Challenge

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
