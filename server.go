package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
	"togglDecks/common"
	"togglDecks/db"
	"togglDecks/logging"
	"togglDecks/routers"
)

func LoggingMiddleware(c *gin.Context) {
	logging.Logger.Info("Request received",
		zap.String("method", c.Request.Method),
		zap.String("path", c.Request.URL.Path),
		zap.String("query", c.Request.URL.RawQuery),
		zap.String("ip", c.ClientIP()),
	)

	c.Next()
}

func main() {
	common.LoadEnvFile()
	db.MigrateDb()

	HOST := os.Getenv("HOST")
	PORT := os.Getenv("PORT")

	router := routers.SetupMainRouter()
	router.Use(LoggingMiddleware)

	// unhandled error, also should set trusted proxies - see https://github.com/gin-gonic/gin/blob/v1.9.1/gin.go#L427
	_ = router.Run(fmt.Sprintf("%s:%s", HOST, PORT))
}
