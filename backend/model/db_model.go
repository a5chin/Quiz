package db_model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Question struct {
	gorm.Model
	year       int    `gorm:"not null"`
	genre      string `gorm:"not null"`
	question   string `gorm:"not null"`
	answer     string `gorm:"not null"`
	commentary string `gorm:"not null"`
}

func DBGetAll() []Question {
	db, err := gorm.Open("sqlite3", "data/questions.db")
	if err != nil {
		panic("You can't open DB (dbGetAll())")
	}
	defer db.Close()
	var questions []Question
	db.Order("created_at desc").Find(&questions)
	return questions
}
