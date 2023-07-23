package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/nawin14/course/db"
	"gitlab.com/nawin14/course/util"
)

type Enrollmenreq struct {
	ClassID  uint `json:"class_id"`
}

func EnrollClass(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := new(Enrollmenreq)
		if err := c.BindJSON(req); err != nil {
			util.SendError(c, http.StatusBadRequest, err)
			return
		}
		class, err := db.GetClass(req.ClassID)
		if err != nil {
			util.SendError(c, http.StatusBadRequest, err)
			return
		}
		student, err := db.GetStudent(User(c).ID)
		if err != nil {
			util.SendError(c, http.StatusBadRequest, err)
			return
		}
		if err := class.AddStudent(*student); err != nil {
			util.SendError(c, http.StatusBadRequest, err)
			return
		}
		if err := db.CreateClassstudent(student.ID, class.ID); err != nil {
			util.SendError(c, http.StatusInternalServerError, err)
			return
		}
		c.Status(http.StatusOK)
	}
}
