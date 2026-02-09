package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// logger
func LoggerMiddleware()gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		ctx.Next()

		duration := time.Since(start)
		log.Printf("%s %s %d %v",ctx.Request.Method,ctx.Request.RequestURI,ctx.Writer.Status(),duration)
	}
}