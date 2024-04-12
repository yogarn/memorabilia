package middleware

import "github.com/gin-gonic/gin"

type Interface interface {
	Timeout() gin.HandlerFunc
}

type middleware struct{}

func Init() Interface {
	return &middleware{}
}
