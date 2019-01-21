package models

import (
	"database/sql"
	"strconv"
)

// Product represents a marketplace product
type Product struct {
	ID             int64   `json:"id"`
	Title          string  `json:"title"`
	Price          float64 `json:"price"`
	InventoryCount int64   `json:"inventoryCount"`
}

// ProductCreationPayload represents a create payload for a product
type ProductCreationPayload struct {
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

// GetAllProducts fetches all products from the database
func GetAllProducts(db *sql.DB) []Product {
	return parseProducts("SELECT product_id, title, price, inventory_count FROM products_view", db)
}

// GetAllAvailableProducts fetches all products from the database that have at least 1 item in inventory
func GetAllAvailableProducts(db *sql.DB) []Product {
	return parseProducts("SELECT product_id, title, price, inventory_count FROM products_view WHERE inventory_count > 0", db)
}

func parseProducts(query string, db *sql.DB) []Product {
	products := []Product{}
	productRows, err := db.Query(query)
	if err == nil {
		defer productRows.Close()
		for productRows.Next() {
			product := Product{}
			if err := productRows.Scan(&(product.ID), &(product.Title), &(product.Price), &(product.InventoryCount)); err != nil {
				return products
			}
			products = append(products, product)
		}
	}
	return products
}

// GetProduct returns a specific product with a specified ID
func GetProduct(id string, db *sql.DB) (Product, error) {
	p := Product{}
	result := db.QueryRow("SELECT product_id, title, price, inventory_count FROM products_view WHERE product_id = $1", id)
	err := result.Scan(&(p.ID), &(p.Title), &(p.Price), &(p.InventoryCount))
	return p, err
}

// CreateProduct creates a new product in database
func CreateProduct(product ProductCreationPayload, db *sql.DB) (Product, error) {
	var id int
	result := db.QueryRow("INSERT INTO products (title, price) VALUES ($1, $2) RETURNING product_id", product.Title, product.Price)
	err := result.Scan(&(id))
	if err == nil {
		return GetProduct(strconv.Itoa(id), db)
	}
	return Product{}, err
}

// DeleteProduct deletes a product in the database
func DeleteProduct(id string, db *sql.DB) error {
	_, err := db.Exec("DELETE FROM products WHERE product_id = $1", id)
	return err
}
