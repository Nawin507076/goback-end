package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gitlab.com/nawin14/course/db"
	"gitlab.com/nawin14/course/handler"
	"gitlab.com/nawin14/course/util"
)

func RequireUser(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get Header
		header := c.GetHeader("Authorization")
		header = strings.TrimSpace(header)
		main := len("Bearer ")
		if len(header) <= main {
			util.SendError(c, http.StatusUnauthorized, errors.New("token is require"))
			return
		}
		// if not found -> error, http status code401
		token := header[main:]
		claims, err := handler.VerifyToken(token)
		if err != nil {
			util.SendError(c, http.StatusUnauthorized, err)
			return
		}
		user, err := db.GetUserByID(claims.UserID)
		if user == nil || err != nil {
			util.SendError(c, http.StatusUnauthorized, err)
			return
		}
		handler.SetUser(c, user)
	}
}
