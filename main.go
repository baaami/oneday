package main

import (
	"github.com/baaami/oneday/db"
	"github.com/baaami/oneday/post"
	"github.com/baaami/oneday/posts"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

const port string = ":4000"

func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())

	// post routing
	router.GET("/post", post.GetPost)
	router.POST("/post", post.PostPost)
	router.PATCH("/post", post.ReplacePost)
	router.DELETE("/post", post.DeletePost)

	// posts routing
	router.GET("/posts", posts.GetPosts)

	// api.SetRouter(router)
	router.Run(port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080"

	db.CloseDB()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Methods", "GET, DELETE, POST, PATCH, OPTIONS")
		c.Header("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
