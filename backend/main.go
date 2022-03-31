package main

import (
	db_model "backend/model"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

// TODO: fix type error

func main() {
	r := gin.Default()
	r.GET("/questions", func(c *gin.Context) {
		questins := db_model.DBGetAll()
		c.JSON(200, questins)
	})
	r.Run()
}
