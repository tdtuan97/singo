package middleware

import (
	"regexp"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Cors Cross-domain configuration
func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie"}
	if gin.Mode() == gin.ReleaseMode {
		// Production environment requires cross-domain domain name, otherwise 403
		config.AllowOrigins = []string{"http://www.example.com"}
	} else {
		// Test environment fuzzy matching local request
		config.AllowOriginFunc = func(origin string) bool {
			if regexp.MustCompile(`^http://127\.0\.0\.1:\d+$`).MatchString(origin) {
				return true
			}
			if regexp.MustCompile(`^http://localhost:\d+$`).MatchString(origin) {
				return true
			}
			return false
		}
	}
	config.AllowCredentials = true
	return cors.New(config)
}
