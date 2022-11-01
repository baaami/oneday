package db

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/baaami/oneday/common"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func DB() *sql.DB {
	if db == nil {
		// init db
		dbPointer, err := sql.Open("mysql", "baami:2719@/oneday")
		db = dbPointer
		if err != nil {
			panic(err)
		}
		createOnedayTable()
		createImageTable()

		db.SetConnMaxLifetime(time.Minute * 3)
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)
	}
	return db
}

func createOnedayTable() {
	// Post Table
	query := `CREATE TABLE IF NOT EXISTS post(
		id INT primary key auto_increment, 
		title TEXT,  
		body TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP, 
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP)
		`
	_, err := DB().Exec(query)
	if err != nil {
		panic(err)
	}
}

func createImageTable() {
	// image Table
	query := `CREATE TABLE IF NOT EXISTS image(
			id INT primary key auto_increment, 
			postId INT,  
			imagePath TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP) 
			`
	_, err := DB().Exec(query)
	if err != nil {
		panic(err)
	}
}

// table, key를 paramter로 받아서 특정 타입의 slice를 return
func SelectPosts() []common.Post {
	var id uint64
	var title, body string
	var posts []common.Post
	var post common.Post

	// query := fmt.Sprintf("SELECT * FROM post")
	query := "SELECT id, title, body FROM post"
	rows, err := DB().Query(query)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err := rows.Scan(&id, &title, &body)
		if err != nil {
			log.Fatal(err)
		}

		post.Id = id
		post.Title = title
		post.Body = body

		posts = append(posts, post)
	}

	return posts
}

func SelectPost(_id uint64) (uint64, string, string) {
	var id uint64
	var title, body string

	rows, err := DB().Query("SELECT id, title, body FROM post WHERE id = ?", _id)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err := rows.Scan(&id, &title, &body)
		if err != nil {
			log.Fatal(err)
		}
	}

	return id, title, body
}

func InsertPost(title, body string) int64 {
	query := "INSERT INTO post (title, body) VALUES(?, ?)"
	result, err := DB().Exec(query, title, body)
	if err != nil {
		log.Fatal(err)
	}

	nlastId, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return nlastId
}

func ReplacePost(id uint64, title, body string) {
	query := "UPDATE post SET title=?, body=? WHERE id=?"
	_, err := DB().Exec(query, title, body, strconv.Itoa(int(id)))
	if err != nil {
		log.Fatal(err)
	}
}

func DeletePost(id uint64) {
	result, err := DB().Exec("DELETE FROM post WHERE id=?", strconv.Itoa(int(id)))
	if err != nil {
		log.Fatal(err)
	}

	nRows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Delete Post Count :", nRows)
}

func UploadImage(postId uint64, imagePath string) {
	query := "INSERT INTO image (postId, imagePath) VALUES(?, ?)"
	_, err := DB().Exec(query, postId, imagePath)
	if err != nil {
		log.Fatal(err)
	}
}

func CloseDB() {
	defer DB().Close()
}
