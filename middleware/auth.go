package middleware

import (
	"context"
	"fmt"
	"go-gin-todo/lib"
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
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		ctx.Abort()
		return
	}
	user, err := client.GetUser(context.Background(), token.UID)
	if err != nil {
		panic(err)
	}

	fmt.Print(user.Email, user.UID, user.PhotoURL, user.DisplayName)

	ctx.Next()
}
