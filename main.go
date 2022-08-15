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

	api.SetRouter(router)
	router.Run(port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080"

	db.CloseDB()
}
