package main

import (
	"backend/model"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/questions", func(c *gin.Context) {
		questions := model.GetAll()
		c.JSON(200, questions)
	})
	r.GET("/:tag/:num", func(c *gin.Context) {
		tag := c.Param("tag")
		num := c.Param("num")
		question := model.GetBy(tag, num)
		c.JSON(200, question)
	})

	r.Run()
}
