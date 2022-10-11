package api

import (
	"github.com/baaami/oneday/post"
	"github.com/baaami/oneday/posts"
	"github.com/gin-gonic/gin"
)

func SetRouter(router *gin.Engine) {
	api := router.Group("/api")

	post.Router(api.Group("/post"))
	posts.Router(api.Group("/posts"))
}
