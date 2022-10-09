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
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Methods", "GET, DELETE, POST, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
