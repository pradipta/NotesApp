package handlers

import (
	"NotesApp/dbconfig"
	"NotesApp/models"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

var err error

func Welcome(c *gin.Context) {
	c.String(200, "Welcome!")
}

func CreateNote(c *gin.Context) {
	var newNote models.Note
	err = c.BindJSON(&newNote)
	if err != nil {
		panic(err)
	}

	t := time.Now()
	newNote.CreatedAt = t
	newNote.UpdatedAt = newNote.CreatedAt
	dbconfig.DB.Create(&newNote)
	fmt.Printf("%v", dbconfig.DB)
	c.JSON(200, newNote)
}

func GetNotes(c *gin.Context) {
	c.String(200, "Here are your notes")
}
