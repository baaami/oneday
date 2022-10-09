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
	c.JSON(http.StatusOK, db.SelectPosts())
}
