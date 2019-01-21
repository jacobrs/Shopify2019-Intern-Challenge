package routers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/jacobrs/Shopify2019-Intern-Challenge/models"
)

// AddCartRoutes adds cart routes to a gin engine
func AddCartRoutes(router *gin.Engine, db *sql.DB) *gin.Engine {

	// add item to existing cart
	router.POST("/v1/carts/:cartId", func(ctx *gin.Context) {
		productID := ctx.Request.URL.Query().Get("productId")
		cartID := ctx.Param("cartId")
		products, err := models.AddToCart(productID, cartID, db)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(200, products)
		}
	})

	// create a new cart and add an item to it
	router.POST("/v1/carts", func(ctx *gin.Context) {
		productID := ctx.Request.URL.Query().Get("productId")
		products, err := models.CreateCart(productID, db)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(200, products)
		}
	})

	// get items in cart
	router.GET("/v1/carts/:cartId", func(ctx *gin.Context) {
		cartID := ctx.Param("cartId")
		products := models.GetCart(cartID, db)
		ctx.JSON(200, products)
	})

	// checkout cart
	router.POST("/v1/checkouts/:cartId", func(ctx *gin.Context) {
		id := ctx.Param("cartId")
		err := models.CheckoutCart(id, db)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
		} else {
			ctx.Status(204)
		}
	})

	// release cart setting all items as available
	router.DELETE("/v1/carts/:cartId", func(ctx *gin.Context) {
		id := ctx.Param("cartId")
		err := models.ReleaseCart(id, db)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
		} else {
			ctx.Status(204)
		}
	})

	return router
}
