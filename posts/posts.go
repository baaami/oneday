package posts

import (
	"net/http"

	"github.com/baaami/oneday/db"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	r.GET("/", getPosts)
}

/*
	글 목록 획득
*/
func getPosts(c *gin.Context) {
	// select
	posts := db.GetValuesbyKeyAndTable("title", "post", 5)
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}
