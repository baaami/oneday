package posts

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	r.GET("/", getPosts)
}

/*
	글 목록 획득
*/
func getPosts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"posts": "get",
	})
}
