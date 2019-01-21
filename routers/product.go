package routers

import (
	"database/sql"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jacobrs/Shopify2019-Intern-Challenge/models"
)

// AddProductRoutes adds product routes to a gin engine
func AddProductRoutes(router *gin.Engine, db *sql.DB) *gin.Engine {

	// get all products
	router.GET("/v1/products", func(ctx *gin.Context) {
		filterAvailable := ctx.Request.URL.Query().Get("onlyAvailable")
		if strings.ToLower(filterAvailable) != "true" {
			ctx.JSON(200, models.GetAllProducts(db))
		} else {
			ctx.JSON(200, models.GetAllAvailableProducts(db))
		}
	})

	// get specific product
	router.GET("/v1/products/:productId", func(ctx *gin.Context) {
		id := ctx.Param("productId")
		product, err := models.GetProduct(id, db)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(200, product)
		}
	})

	// create new product
	router.POST("/v1/products", func(ctx *gin.Context) {
		var product models.ProductCreationPayload
		err := ctx.BindJSON(&product)
		if err != nil {
			ctx.Status(400)
			return
		}
		createdProduct, err := models.CreateProduct(product, db)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(200, createdProduct)
		}
	})

	// delete a specific product
	router.DELETE("/v1/products/:productId", func(ctx *gin.Context) {
		id := ctx.Param("productId")
		err := models.DeleteProduct(id, db)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
		} else {
			ctx.Status(200)
		}
	})

	return router
}
