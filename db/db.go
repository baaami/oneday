package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

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

		db.SetConnMaxLifetime(time.Minute * 3)
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)
	}
	return db
}

func createOnedayTable() {
	query := `CREATE TABLE IF NOT EXISTS post(id int primary key auto_increment, title text,  
		body text, created_at datetime default CURRENT_TIMESTAMP, updated_at datetime default CURRENT_TIMESTAMP)`
	_, err := DB().Exec(query)
	if err != nil {
		panic(err)
	}
}

// table, key를 paramter로 받아서 특정 타입의 slice를 return
func GetValuesbyKeyAndTable(key, table string, limit int) []string {
	var value string
	var values []string

	query := fmt.Sprintf("SELECT %s FROM %s LIMIT %d", key, table, limit)
	rows, err := DB().Query(query)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err := rows.Scan(&value)
		if err != nil {
			log.Fatal(err)
		}
		values = append(values, value)
	}

	return values
}

func CloseDB() {
	defer DB().Close()
}
