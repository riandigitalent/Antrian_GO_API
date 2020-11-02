package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var data []Antrian

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
	router.Run(":8080")
}


//model
func getAntrian()(bool, []Antrian,error){
	return true, data, nil
}
func addAntrian() (bool,error){
	_, dataAntrian, _ :=getAntrian()
	var IDs string

	if dataAntrian==nil {
		IDs = fmt.Sprintf("B-0")
	}else {
		IDs = fmt.Sprintf("B-%d", len(dataAntrian))
	}
	data= append(data, Antrian{
		ID: IDs,
		STATUS: false,
	})
	return true,nil
}

func updateAntrian(idAntrian string) (bool,error){
	for i := range data {
		if data[i].ID == idAntrian{
			data[i].STATUS =true
		break
		}
	}
	return true,nil
}

func deleteAntrian(idAntrian string)(bool,error){
	for i := range data {
		if data[i].ID == idAntrian{
			data = append(data[:i],data[i+1:]...)
		}
	}
	return true,nil
}


//handler
func GetAntrianHandler(c *gin.Context){
	flag, data, err := getAntrian()
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
			"status" : "Antrian Berhasil di tambahkan",
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
			"status" : "Antrian Berhasil di tambahkan",
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

