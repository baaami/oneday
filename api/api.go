package api

import (
	"github.com/baaami/oneday/post"
	"github.com/baaami/oneday/posts"
	"github.com/gin-gonic/gin"
)

func SetRouter(router *gin.Engine) {
	post.Router(router.Group("/post"))
	posts.Router(router.Group("/posts"))
}
