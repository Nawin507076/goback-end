package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Course struct {
	ID          string
	Name        string
	Description string
}

var courses = []Course{
	{ID: "1", Name: "nawin kaewlong", Description: "learngo"},
	{ID: "2", Name: "nawin kaewlong", Description: "learngo"},
}

func main() {
	r := gin.Default()

	r.GET("/courses", ListCourses)
	r.GET("/courses/:id", getCourse)

	r.Run(":8080")
}

func ListCourses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, courses)
}

func getCourse(c *gin.Context) {
	id := c.Param("id")
	for _, course := range courses {
		if course.ID == id {
			c.IndentedJSON(http.StatusOK, course)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "course not found",
})
}
