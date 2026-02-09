package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func ErrorHandling() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		defer func(){
			if err :=recover(); err != nil{
				log.Printf("Panic occured: %v",err)
			}
		}()
		ctx.Next()
	}
}