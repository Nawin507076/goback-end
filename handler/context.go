package handler

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/nawin14/course/model"
)

const key string = "user"

func SetUser(c *gin.Context, user *model.User) {
		c.Set(key, user)
}

func User(c *gin.Context) *model.User {
	user, ok := c.Value(key).(*model.User)
	if !ok {
		return nil 
	}
	return user
}