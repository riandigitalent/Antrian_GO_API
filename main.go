package main

import (
	"github.com/gin-gonic/gin"
	"github.com/riandigitalent/Antrian_GO_API/app/controller"
)

func main(){
	router := gin.Default()

	router.LoadHTMLGlob("views/*")
	//setting endpoint API
	router.GET("/", controller.HomePage)
	router.GET("/antrian/status", controller.GetAntrianHandler)
	router.POST("/antrian/add", controller.AddAntrianHandler)
	router.PUT("/antrian/edit/id/:idAntrian", controller.UpdateAntrianHandler)
	router.DELETE("/antrian/id/:idAntrian/delete", controller.DeleteAntrianHandler)
	router.GET("/antrian/list",controller.PageAntrianHandler)
	router.Run(":8080")
}





