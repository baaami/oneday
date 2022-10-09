package main

import (
	"github.com/baaami/oneday/api"
	"github.com/baaami/oneday/db"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

const port string = ":4000"

func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())
	api.SetRouter(router)
	router.Run(port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080"

	db.CloseDB()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, Origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, PATCH, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
