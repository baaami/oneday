package post

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/baaami/oneday/common"
	"github.com/baaami/oneday/db"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	r.GET("/read", GetPost)
	r.POST("/write", PostPost)
	r.PATCH("/update", ReplacePost)
	r.DELETE("/delete", DeletePost)
}

/*
	글 획득
*/
func GetPost(c *gin.Context) {
	id, err := strconv.ParseUint(c.Query("id"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	var _post common.Post
	_post.Id, _post.Title, _post.Body = db.SelectPost(id)

	c.JSON(http.StatusOK, _post)
}

/*
	글 등록
*/
func PostPost(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		log.Fatal(err)
	}

	var _post common.Post
	// json to common.Post (golang struct)
	json.Unmarshal([]byte(value), &_post)

	fmt.Printf("%v", _post)

	id := db.InsertPost(_post.Title, _post.Body)
	_post.Id = uint64(id)

	c.JSON(http.StatusOK, _post)
}

/*
	글 업데이트
*/
func ReplacePost(c *gin.Context) {
	id, err := strconv.ParseUint(c.Query("id"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		log.Fatal(err)
	}

	// json to map
	var _post common.Post
	json.Unmarshal([]byte(value), &_post)

	db.ReplacePost(id, _post.Title, _post.Body)

	c.JSON(http.StatusOK, nil)
}

func DeletePost(c *gin.Context) {
	id, err := strconv.ParseUint(c.Query("id"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	db.DeletePost(id)

	c.JSON(http.StatusOK, nil)
}
