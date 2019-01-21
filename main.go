package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jacobrs/Shopify2019-Intern-Challenge/routers"
	_ "github.com/lib/pq"
)

func getDatabase() *sql.DB {
	databaseInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"127.0.0.1",
		5432,
		"postgres",
		"postgres",
		"shopify")
	db, _ := sql.Open("postgres", databaseInfo)
	return db
}

func main() {

	router := gin.Default()
	database := getDatabase()
	routers.AddProductRoutes(router, database)

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	var dbConnectionErr error
	dbConnectionStatus := "OK"

	go func(err *error, db *sql.DB) {
		for true {
			dbConnectionErr = db.Ping()
			time.Sleep(10 * time.Second)
		}
	}(&dbConnectionErr, database)

	router.GET("/health", func(c *gin.Context) {
		returnCode := 200
		if dbConnectionErr != nil {
			returnCode = 503
			dbConnectionStatus = "FAIL"
		} else {
			dbConnectionStatus = "OK"
		}
		c.JSON(returnCode, gin.H{
			"status": "OK",
			"postgres_connection": gin.H{
				"status": dbConnectionStatus,
				"error":  fmt.Sprintf("%s", dbConnectionErr),
			},
		})
	})

	router.Run()

}
