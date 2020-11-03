package model

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
	"log"
)

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

