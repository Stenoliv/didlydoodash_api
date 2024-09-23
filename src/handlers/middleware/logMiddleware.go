package middleware

import (
	"bytes"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type LoggerBody struct {
	gin.ResponseWriter
	body bytes.Buffer
}

func (g *LoggerBody) Write(b []byte) (int, error) {
	g.body.Write(b)
	return g.ResponseWriter.Write(b)
}

func RequestLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Starting time request
		startTime := time.Now()

		// Processing request
		c.Next()

		// End Time request
		endTime := time.Now()

		// execution time
		latencyTime := endTime.Sub(startTime)

		// Request method
		reqMethod := c.Request.Method

		// Request route
		reqUri := c.Request.RequestURI

		// status code
		statusCode := c.Writer.Status()

		// Request IP
		clientIP := c.ClientIP()

		log.WithFields(log.Fields{
			"METHOD":    reqMethod,
			"URI":       reqUri,
			"STATUS":    statusCode,
			"LATENCY":   latencyTime,
			"CLIENT_IP": clientIP,
		}).Info("HTTP REQUEST")

		c.Next()
	}
}
