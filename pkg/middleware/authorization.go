package middleware

import (
	"errors"
	"memorabilia/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (m *middleware) OnlyAdmin(ctx *gin.Context) {
	userId, err := m.jwtAuth.GetLoginUser(ctx)
	if err != nil {
		response.Error(ctx, http.StatusForbidden, "failed to get userId", err)
		ctx.Abort()
		return
	}

	user, err := m.service.UserService.GetUserById(userId)
	if err != nil {
		response.Error(ctx, http.StatusForbidden, "failed to get login user", err)
		ctx.Abort()
		return
	}

	if user.RoleID != 1 {
		response.Error(ctx, http.StatusForbidden, "insufficient privilege", errors.New("only admin can access this endpoint"))
		ctx.Abort()
		return
	}

	ctx.Next()
}
