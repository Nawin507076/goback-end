package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/nawin14/course/db"
	"gitlab.com/nawin14/course/util"
)

type ClassReq struct {
	ID        uint      `json:"id"`
	CourseId  uint      `json:"course_id"`
	TrainerID uint      `json:"trainer_id"`
	Start     time.Time `json:"start"`
	End       time.Time `json:"end"`
	Seats     int       `json:"seats"`
}

func CreateClass(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := new(ClassReq)
		if err := c.BindJSON(req); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			util.SendError(c, http.StatusBadRequest, err)
			return
		}
		course, err := db.GetCourse(req.CourseId)
		if err != nil {
			util.SendError(c, http.StatusNotFound, err)
			return
		}
		class, err := course.CreateClass(req.Start, req.End)
		if err != nil {
			util.SendError(c, http.StatusBadRequest, err)
			return
		}
		if err := class.SetSeats(req.Seats); err != nil {
			util.SendError(c, http.StatusBadRequest, err)
			return
		}
		class.Trainer.ID = req.TrainerID

		if err := db.CreateClass(class); err != nil {
			util.SendError(c, http.StatusInternalServerError, err)
			return
		}

		c.Status(http.StatusOK)
	}

}
