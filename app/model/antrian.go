package model

import (
	"fmt"
	"log"
	"strings"
	"firebase.google.com/go/db"
)

type Antrian struct {
	ID string `json:"id"`
	STATUS bool `json:"status"`
}

//model
func GetAntrian()(bool,error, []map[string]interface{}){
	var data []map[string]interface{}
	ref := client.NewRef("antrian")
	if err := ref.Get(ctx,&data); err != nil {
		log.Fatalln("error baca data dari DB:", err)
		return false,err,nil
	}

	return true, nil, data
}

//AddAntrian
func AddAntrian() (bool,error){
	_, _, dataAntrian := GetAntrian()
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

func UpdateAntrian(idAntrian string) (bool,error){
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

func DeleteAntrian(idAntrian string)(bool,error){
	ref := client.NewRef("antrian")
	id := strings.Split(idAntrian,"-")
	childRef := ref.Child(id[1])

	if err:= childRef.Delete(ctx); err != nil{
		log.Fatal(err)
		return false,err
	}
	return  true,nil
}