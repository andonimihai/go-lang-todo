package helper

import (
	"go-gin-todo/entity"

	"github.com/gin-gonic/gin"
)

func GetLoggedInUser(ctx *gin.Context) entity.User {
	user := ctx.MustGet("user").(entity.User)
	return user
}
