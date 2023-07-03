package main

import (
	"database/sql"
// 	"encoding/json"
	"log"

	 "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// MySQLデータベースへの接続情報
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/go_mysql")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Ginルーターの作成
	r := gin.Default()

	// /peopleエンドポイントへのGETリクエストを処理するハンドラ
	r.GET("/people", func(c *gin.Context) {
		// データのクエリ
		rows, err := db.Query("SELECT * FROM test")
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		// 結果を格納するスライス
		people := []Person{}

		// 結果をスキャンしてスライスに追加
		for rows.Next() {
			var person Person
			err := rows.Scan(&person.ID, &person.Name, &person.Age)
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
			people = append(people, person)
		}
		if err := rows.Err(); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// JSONデータをレスポンスとして返す
		c.JSON(200, people)
	})

	// サーバーの開始
	r.Run(":8000")
}
