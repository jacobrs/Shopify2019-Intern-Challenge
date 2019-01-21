package models

import (
	"database/sql"
	"fmt"
	"strconv"
)

// CreateCart creates a new cart with an item inside
func CreateCart(productID string, db *sql.DB) ([]Product, error) {
	var id int
	itemID, err := LockItem(productID, db)
	if err == nil {
		result := db.QueryRow("INSERT INTO carts (item_id) VALUES ($1) RETURNING id", itemID)
		err := result.Scan(&(id))
		if err == nil {
			return GetCart(strconv.Itoa(id), db), nil
		}
		return []Product{}, err
	}
	return []Product{}, err
}

// AddToCart adds item of product x to cart y
func AddToCart(productID string, cartID string, db *sql.DB) ([]Product, error) {
	itemID, err := LockItem(productID, db)
	if err == nil {
		_, err := db.Exec("INSERT INTO carts (id, item_id) VALUES ($1, $2)", cartID, itemID)
		if err == nil {
			return GetCart(cartID, db), nil
		}
		return []Product{}, err
	}
	return []Product{}, err
}

// LockItem finds itemID for product and marks it unavailable
func LockItem(productID string, db *sql.DB) (int, error) {
	var id int
	result := db.QueryRow("SELECT item_id FROM inventory WHERE available = true AND product_id = $1", productID)
	err := result.Scan(&(id))
	_, err = db.Exec("UPDATE inventory SET available = false WHERE item_id = $1", id)
	return id, err
}

// GetCart fetches a cart from the database
func GetCart(cartID string, db *sql.DB) []Product {
	query := fmt.Sprintf("SELECT p.product_id, p.title, p.price, COUNT(c.item_id) as cart_count FROM inventory i JOIN carts c ON (i.item_id = c.item_id) JOIN products p ON (i.product_id = p.product_id) WHERE c.id = %s GROUP BY p.product_id", cartID)
	return ParseProducts(query, db)
}

// CheckoutCart currently simply deletes the cart and keeps the items unavailable
func CheckoutCart(cartID string, db *sql.DB) error {
	_, err := db.Exec("DELETE FROM carts WHERE id = $1", cartID)
	return err
}

// ReleaseCart currently simply deletes the cart and resets the items to available
func ReleaseCart(cartID string, db *sql.DB) error {
	_, err := db.Exec("UPDATE inventory SET available = true WHERE item_id IN (SELECT item_id FROM carts WHERE id = $1)", cartID)
	_, err = db.Exec("DELETE FROM carts WHERE id = $1", cartID)
	return err
}
