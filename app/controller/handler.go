package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/riandigitalent/Antrian_GO_API/app/model"
	"net/http"
)

//handler
func GetAntrianHandler(c *gin.Context){
	flag, err, data := model.GetAntrian()
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
	flag, err := model.AddAntrian()
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
	flag, err := model.UpdateAntrian(idAntrian)
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
	flag, err := model.DeleteAntrian(idAntrian)
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
	flag, err, result := model.GetAntrian()
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
