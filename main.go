package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/data/list",ForGet)
	r.POST("/data/add",ForPost)
}
func ForPost(c *gin.Context)  {
	fmt.Println("helo")
	c.JSON(200,"succes")
}
func ForGet(c *gin.Context)  {
	fmt.Println("privet")
	c.JSON(200,"succes2")
}