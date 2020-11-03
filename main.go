package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"

)

//var data []Antrian






func main(){
	router := gin.Default()
	//setting endpoint API
	router.GET("/", HomePage)
	router.GET("/antrian/status", GetAntrianHandler)
	router.POST("/antrian/add", AddAntrianHandler)
	router.PUT("/antrian/edit/id/:idAntrian", UpdateAntrianHandler)
	router.DELETE("/antrian/id/:idAntrian/delete", DeleteAntrianHandler)
	router.GET("/antrian/list",PageAntrianHandler)
	router.Run(":8080")
}





