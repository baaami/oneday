package posts

import (
	"net/http"

	"github.com/baaami/oneday/db"
	"github.com/baaami/oneday/post"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	r.GET("/", getPosts)
}

/*
	글 목록 획득
*/
func getPosts(c *gin.Context) {
	var posts []post.Post

	// select
	MapPosts := db.SelectPost()
	for key, value := range MapPosts {
		var post post.Post
		post.Title = key
		post.Body = value

		posts = append(posts, post)
	}

	c.JSON(http.StatusOK, posts)
}
