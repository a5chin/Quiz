package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Question struct {
	ID         int    `gorm:"primary_key"`
	Year       int    `gorm:"not null"`
	Genre      string `gorm:"not null"`
	Question   string `gorm:"not null"`
	Answer     string `gorm:"not null"`
	Commentary string `gorm:"not null"`
}

func GetAll() []Question {
	db, err := gorm.Open("sqlite3", "data/questions.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	questions := []Question{}
	db.Find(&questions)

	return questions
}

func GetBy(tag string, num string) []Question {
	db, err := gorm.Open("sqlite3", "data/questions.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	questions := []Question{}
	db.Where(fmt.Sprintf("%s = ?", tag), num).Find(&questions)

	return questions
}
