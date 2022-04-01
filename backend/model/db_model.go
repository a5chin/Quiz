package db_model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Question struct {
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

	questions := []Question{}
	db.Find(&questions)

	// return None
	return questions
}
