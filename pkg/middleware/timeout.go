package middleware

import (
	"errors"
	"log"
	"memorabilia/pkg/response"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

func (m *middleware) Timeout() gin.HandlerFunc {
	timeLimit, err := strconv.Atoi(os.Getenv("TIME_OUT_LIMIT"))
	if err != nil {
		log.Fatalf("can't convert time out limit: %v", err)
		return nil
	}

	return timeout.New(
		timeout.WithTimeout(time.Duration(timeLimit)*time.Second),
		timeout.WithHandler(func(ctx *gin.Context) {
			ctx.Next()
		}),
		timeout.WithResponse(func(ctx *gin.Context) {
			response.Error(ctx, http.StatusRequestTimeout, "timeout", errors.New("request timeout"))
		}),
	)
}
