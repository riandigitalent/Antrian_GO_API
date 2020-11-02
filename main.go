package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	router := gin.Default()
	router.GET("/", ambilapakek)
	router.Run(":8080")
}

func ambilapakek(c *gin.Context){
	c.JSON(http.StatusOK,map[string]interface{}{
		"body":" Sukses BRO ",
	})
	return
}

