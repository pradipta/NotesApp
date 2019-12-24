package routers

import (
	"NotesApp/handlers"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine{
	r := gin.Default()

	//r.Use();

	r.GET("", handlers.Welcome)
	r.POST("/v1/notes/create", handlers.CreateNote)
	r.GET("/v1/notes/get", handlers.GetNotes)

	return r
}
