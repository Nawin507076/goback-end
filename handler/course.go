package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/nawin14/course/db"
	"gitlab.com/nawin14/course/model"
	"gitlab.com/nawin14/course/util"
)

func Error(c *gin.Context, status int, err error) {
	log.Println(err)
	c.JSON(status, gin.H{
		"message": err.Error(),
	})
}

func ListCourses(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		courses, err := db.GetAllCourse()
		if err != nil {
			util.SendError(c, http.StatusInternalServerError, err)
			return
		}
		c.IndentedJSON(http.StatusOK, courses)

	}
}

func GetCourse(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			util.SendError(c, http.StatusNotFound, err)

			return
		}
		course, err := db.GetCourse(uint(id))
		if err != nil {
			util.SendError(c, http.StatusBadRequest, err)
			return
		}
		c.IndentedJSON(http.StatusOK, course)
	}

}

func CreateCourse(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := new(model.Course)
		if err := c.BindJSON(req); err != nil {
			util.SendError(c, http.StatusBadRequest, err)
			return
		}

		if err := db.CreateCourse(req); err != nil {
			util.SendError(c, http.StatusInternalServerError, err)
			return
		}
		c.IndentedJSON(http.StatusOK, req)
	}
}
