package main

import (
	"NotesApp/dbconfig"
	"NotesApp/models"
	"NotesApp/routers"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var Database models.Note
var err error

//var DatabaseStruct models.Database

func main() {

	if err != nil {
		fmt.Println(err)
	}

	defer dbconfig.DB.Close()
	//	dbconfig.DB.AutoMigrate(models.Note{})

	r := routers.GetRouter()
	r.Run("0.0.0.0:4001")
}
