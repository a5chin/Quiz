package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

// TODO: fix type error
type Question struct {
	gorm.Model
	year       int    `gorm:"not null"`
	genre      string `gorm:"not null"`
	question   string `gorm:"not null"`
	answer     string `gorm:"not null"`
	commentary string `gorm:"not null"`
}

func dbGetAll() []Question {
	db, err := gorm.Open("sqlite3", "data/questions.db")
	if err != nil {
		panic("You can't open DB (dbGetAll())")
	}
	fmt.Print(db)
	defer db.Close()
	var questions []Question
	db.Order("created_at desc").Find(&questions)
	return questions
}

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		questins := dbGetAll()
		c.JSON(200, questins)
	})
	r.Run()
}
