# Shopify Summer 2019 Internship Challenge

### Endpoint Documentation

| Route | Type | Example Payload | Example Response |
|---|---|---|---|
| /health | GET | None | {"status":"OK"} |
| /v1/products | GET | None | Array[ProductModel] |
| /v1/products/:id | GET | None | [ProductModel] |
| /v1/products | POST | [PostProductModel] | [ProductModel] |
| /v1/products | DELETE | None | None |

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
