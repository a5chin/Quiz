package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

const db_file string = "data/questions.db"

type Question struct {
	ID         int    `gorm:"primary_key"`
	Year       int    `gorm:"not null"`
	Genre      string `gorm:"not null"`
	Question   string `gorm:"not null"`
	Answer     string `gorm:"not null"`
	Commentary string `gorm:"not null"`
}

func ReadDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", db_file)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	return db
}

func GetAll() []Question {
	db := ReadDB()
	defer db.Close()

	questions := []Question{}
	db.Find(&questions)

	return questions
}

func GetBy(tag string, num string) []Question {
	db := ReadDB()
	defer db.Close()

	questions := []Question{}
	db.Where(fmt.Sprintf("%s = ?", tag), num).Find(&questions)

	return questions
}
