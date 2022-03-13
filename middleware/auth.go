package middleware

import (
	"context"
	"go-gin-todo/entity"
	"go-gin-todo/lib"
	"go-gin-todo/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context) {

	client := lib.FirebaseClient

	authorizationToken := ctx.GetHeader("Authorization")
	idToken := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))
	if idToken == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Id token not available"})
		ctx.Abort()
		return
	}

	//verify token
	token, err := client.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	user, err := client.GetUser(context.Background(), token.UID)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	dbUser, err := service.UpsertUser(entity.UpsertUser{
		Name:   user.DisplayName,
		Email:  user.Email,
		UserId: user.UID,
		Avatar: user.PhotoURL,
	})

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.Set("user", dbUser)
	ctx.Next()
}
