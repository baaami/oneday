package image

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/baaami/oneday/db"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	r.POST("/upload", uploadHandler)
}

// 싱글 파일 업로드 핸들러
func uploadHandler(c *gin.Context) {
	postId, err := strconv.ParseUint(c.Query("id"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	// Single file
	file, _ := c.FormFile("image")

	path, _ := os.Getwd()
	path = filepath.Join(path, "/public/images", file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		panic(err)
	}

	db.UploadImage(postId, path)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded! (with %d)", file.Filename, postId))
}
