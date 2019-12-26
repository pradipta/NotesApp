package handlers

import (
	"NotesApp/dbconfig"
	"NotesApp/models"
	"time"

	"github.com/gin-gonic/gin"
)

func Welcome(c *gin.Context) {
	c.String(200, "Welcome!")
}

func CreateNote(c *gin.Context) {
	var newNote models.Note
	c.BindJSON(&newNote)
	t := time.Now()
	newNote.CreatedAt = t.Format(time.UnixDate)
	newNote.UpdatedAt = newNote.CreatedAt

	dbconfig.DB.Create(&newNote)
	c.JSON(200, newNote)
}

func GetNotes(c *gin.Context) {
	var notes []models.Note
	if err := dbconfig.DB.Find(&notes).Error; err != nil {
		c.AbortWithError(404, err)
	}

	c.JSON(200, notes)

}
