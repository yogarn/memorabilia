package middleware

import (
	"errors"
	"memorabilia/pkg/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (m *middleware) AuthenticateUser(ctx *gin.Context) {
	bearer := ctx.GetHeader("Authorization")

	if bearer == "" {
		response.Error(ctx, http.StatusUnauthorized, "empty token", errors.New("empty token at Authorization header"))
		ctx.Abort()
		return
	}

	token := strings.Split(bearer, " ")[1]
	userID, err := m.jwtAuth.ValidateToken(token)
	if err != nil {
		response.Error(ctx, http.StatusUnauthorized, "invalid token", err)
		ctx.Abort()
		return
	}

	user, err := m.service.UserService.GetUserById(userID)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to get user details", err)
		ctx.Abort()
		return
	}

	ctx.Set("user", user)
	ctx.Next()
}
