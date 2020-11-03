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

var client *db.Client
var ctx context.Context

func init(){
	ctx = context.Background()
	conf := &firebase.Config{
		DatabaseURL:"https://antriango.firebaseio.com/",
	}
	//ambil service account key dari json file content
	opt := option.WithCredentialsFile("firebasekunci.json")

	//inisiasi aplikasi dengan service account, dan kasih admin leverl
	app,err := firebase.NewApp(ctx,conf,opt)
	if err != nil{
		log.Fatalln("error inisiasi app :", err)
	}
	client,err =app.Database(ctx)
	if err !=nil {
		log.Fatalln("error inisiasi database client", err)
	}
}

type Antrian struct {
	ID string `json:"id"`
	STATUS bool `json:"status"`
}

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


//model
func getAntrian()(bool,error, []map[string]interface{}){
	var data []map[string]interface{}
	ref := client.NewRef("antrian")
	if err := ref.Get(ctx,&data); err != nil {
		log.Fatalln("error baca data dari DB:", err)
		return false,err,nil
	}

	return true, nil, data
}



func addAntrian() (bool,error){
	_, _, dataAntrian := getAntrian()
	var IDs string
	var antrianRef *db.Ref
	ref := client.NewRef("antrian")


	if dataAntrian==nil {
		IDs = fmt.Sprintf("B-0")
		antrianRef=ref.Child("0")
	}else {
		IDs = fmt.Sprintf("B-%d", len(dataAntrian))
		antrianRef = ref.Child(fmt.Sprintf("%d",len(dataAntrian)))
	}
	antrian :=  Antrian{
		ID: IDs,
		STATUS: false,
	}
	if err := antrianRef.Set(ctx,antrian); err !=nil{
		log.Fatal(err)
		return false,err
	}
	return true,nil
}

func updateAntrian(idAntrian string) (bool,error){
	ref := client.NewRef("antrian")
	id := strings.Split(idAntrian,"-")
	childRef := ref.Child(id[1])

	antrian :=  Antrian{
		ID: idAntrian,
		STATUS: true,
	}
	if err := childRef.Set(ctx,antrian); err !=nil{
		log.Fatal(err)
		return false,err
	}
	return true,nil
}

func deleteAntrian(idAntrian string)(bool,error){
	ref := client.NewRef("antrian")
	id := strings.Split(idAntrian,"-")
	childRef := ref.Child(id[1])

	if err:= childRef.Delete(ctx); err != nil{
		log.Fatal(err)
		return false,err
	}
	return  true,nil
}


//handler
func GetAntrianHandler(c *gin.Context){
	flag, err, data := getAntrian()
	if flag{
		c.JSON(http.StatusOK,map[string]interface{}{
			"status" : "sukses, berikut daftar antrian",
			"data": data,
		})
	} else {
		c.JSON(http.StatusBadGateway,map[string]interface{}{
			"status" : "Gagal",
			"error" : err,
		})
	}
}

func AddAntrianHandler(c *gin.Context){
	flag, err := addAntrian()
	if flag{
		c.JSON(http.StatusOK,map[string]interface{}{
			"status" : "Antrian Berhasil di tambahkan",
		})
	} else {
		c.JSON(http.StatusBadGateway,map[string]interface{}{
			"status" : "Gagal",
			"error" : err,
		})
	}
}
func UpdateAntrianHandler(c *gin.Context){
idAntrian := c.Param("idAntrian")
flag, err := updateAntrian(idAntrian)
	if flag{
		c.JSON(http.StatusOK,map[string]interface{}{
			"status" : "Antrian Berhasil di ubah",
		})
	} else {
		c.JSON(http.StatusBadGateway,map[string]interface{}{
			"status" : "Gagal",
			"error" : err,
		})
	}
}
func DeleteAntrianHandler(c *gin.Context){
	idAntrian := c.Param("idAntrian")
	flag, err := deleteAntrian(idAntrian)
	if flag{
		c.JSON(http.StatusOK,map[string]interface{}{
			"status" : "Antrian Berhasil di di delete",
		})
	} else {
		c.JSON(http.StatusBadGateway,map[string]interface{}{
			"status" : "Gagal",
			"error" : err,
		})
	}
}


func HomePage(c *gin.Context){
	c.JSON(http.StatusOK,map[string]interface{}{
		"body":" Welcome BRO ",
	})
	return
}

func PageAntrianHandler(c *gin.Context) {
	flag, err, result := getAntrian()
	var currentAntrian map[string]interface{}

	for _, item := range result {
		if item != nil {
			currentAntrian = item
			break
		}
	}

	if flag && len(result) > 0 {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"antrian": currentAntrian["id"],
		})
	} else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed",
			"error":  err,
		})
	}
}