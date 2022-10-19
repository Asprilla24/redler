package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// JSONLogMiddleware logs a gin HTTP request in JSON format, with some additional custom key/values
func JSONLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Process Request
		c.Next()

		entry := logrus.WithFields(logrus.Fields{
			"method":     c.Request.Method,
			"path":       c.Request.RequestURI,
			"status":     c.Writer.Status(),
			"referrer":   c.Request.Referer(),
			"request_id": c.Writer.Header().Get("Request-Id"),
			// "api_version": util.ApiVersion,
		})

		if c.Writer.Status() >= 500 {
			entry.Error(c.Errors.String())
		} else {
			entry.Info("")
		}
	}
}
