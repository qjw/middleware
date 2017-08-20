package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ServeHTTP func(http.ResponseWriter, *http.Request)()
func GinHandle(handle ServeHTTP) gin.HandlerFunc {
	return func(c *gin.Context) {
		handle(c.Writer,c.Request)
		c.Next()
	}
}