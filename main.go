package main

import (
	"NotesApp/dbconfig"
	"NotesApp/routers"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var err error

//var DatabaseStruct models.Database

func main() {

	dbconfig.NewConnection("root", "", "127.0.0.1")
	//fmt.Println("Here \n\n\n")
	//fmt.Println(models.Database{})
	if err != nil {
		fmt.Println(err)
	}

	defer dbconfig.DB.Close()
	//	dbconfig.DB.AutoMigrate(models.Note{})

	r := routers.GetRouter()
	r.Run(":4000")
}
