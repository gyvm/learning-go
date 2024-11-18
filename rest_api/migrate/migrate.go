package main

import (
	"fmt"
	"log"
	"rest_api/db"
	"rest_api/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrate")
	defer db.CloseDB(dbConn)
	err := dbConn.AutoMigrate(&model.User{}, &model.Task{})
	if err != nil {
		log.Fatalln(err)
	}
}
