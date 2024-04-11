package response

import "github.com/gin-gonic/gin"

type Response struct {
	Status  Status `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type Status struct {
	Code      int  `json:"code"`
	IsSuccess bool `json:"isSuccess"`
}

func Success(ctx *gin.Context, code int, message string, data any) {
	ctx.JSON(code, Response{
		Status: Status{
			Code:      code,
			IsSuccess: true,
		},
		Message: message,
		Data:    data,
	})
}

func Error(ctx *gin.Context, code int, message string, err error) {
	ctx.JSON(code, Response{
		Status: Status{
			Code:      code,
			IsSuccess: false,
		},
		Message: message,
		Data:    err.Error(),
	})
}
