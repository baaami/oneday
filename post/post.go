package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	r.GET("/", getPost)
	r.POST("/", postPost)
	r.PATCH("/", updatePost)
	r.DELETE("/", deletePost)
}

/*
	글 획득
*/
func getPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"post": "get",
	})
}

/*
	글 등록
*/
func postPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"post": "post",
	})
}

/*
	글 업데이트
*/
func updatePost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"post": "update",
	})
}

func deletePost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"post": "delete",
	})
}
