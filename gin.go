package middleware

import (
	"github.com/gin-gonic/gin"
)

func GinHandle(handle ServeHTTP) gin.HandlerFunc {
	return func(c *gin.Context) {
		if handle(c.Writer,c.Request){
			c.Next()
		}else{
			c.Abort()
		}
	}
}