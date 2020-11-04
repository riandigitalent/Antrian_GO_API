package main

import (
	"github.com/gin-gonic/gin"
	"github.com/riandigitalent/Antrian_GO_API/App/Controller"
	)

func main(){
	router := gin.Default()

	router.LoadHTMLGlob("views/*")
	//setting endpoint API
	router.GET("/", HomePage)
	router.GET("/antrian/status", GetAntrianHandler)
	router.POST("/antrian/add", AddAntrianHandler)
	router.PUT("/antrian/edit/id/:idAntrian", UpdateAntrianHandler)
	router.DELETE("/antrian/id/:idAntrian/delete", DeleteAntrianHandler)
	router.GET("/antrian/list",PageAntrianHandler)
	router.Run(":8080")
}





