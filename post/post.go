package post

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/baaami/oneday/db"
	"github.com/gin-gonic/gin"
)

type Post struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

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
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}

	// json to map
	var _post Post
	json.Unmarshal([]byte(value), &_post)

	db.InsertPost(_post.Title, _post.Body)

	c.JSON(http.StatusOK, _post)
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
