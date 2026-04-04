package middlewares

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggingMiddleware logs basic info about requests and errors to a daily log file
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Process request
		c.Next()

		// Only log if status code is 500
		if c.Writer.Status() < 500 {
			return
		}

		// Prepare log message
		duration := time.Since(start)
		statusCode := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path
		logMsg := fmt.Sprintf("[%s] %s %s %d %s\n", time.Now().Format(time.RFC3339), method, path, statusCode, duration)

		// If there are errors in the context, append them too
		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				logMsg += fmt.Sprintf("ERROR: %s\n", e.Error())
			}
		}

		// Determine log file name: Log-YYYY-MM-DD.txt
		logFileName := fmt.Sprintf("Log-%s.txt", time.Now().Format("2006-01-02"))

		// Open file in append mode or create if it doesn't exist
		f, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			// fallback to console if file can't be opened
			fmt.Println("Failed to open log file:", err)
			fmt.Print(logMsg)
			return
		}
		defer f.Close()

		// Write log message
		if _, err := f.WriteString(logMsg); err != nil {
			fmt.Println("Failed to write log:", err)
		}
	}
}
