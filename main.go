package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "admin:123456789@tcp(testdb.cyrgqq4frcl9.us-east-2.rds.amazonaws.com:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&Memo{})
}

func main() {
	r := gin.Default()

	r.GET("/health", HealthCheck)
	r.GET("/memos", GetMemos)
	r.GET("/memos/:id", GetMemo)
	r.POST("/memos", CreateMemo)
	r.PUT("/memos/:id", UpdateMemo)
	r.DELETE("/memos/:id", DeleteMemo)

	r.Run(":8080")
}
