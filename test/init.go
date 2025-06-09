package test

import (
	"os"
	"singo/cache"
	"singo/conf"
	"singo/model"
	"singo/server"
	"singo/util"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	s *gin.Engine
)

func init() {
	// Read configuration from configuration file
	confInit()
	// API
	s = server.NewRouter()
}

// Init Initialize configuration items
func confInit() {
	// Read environment variables from local file
	godotenv.Load()

	// Set log level
	util.BuildLogger(os.Getenv("LOG_LEVEL"))

	// Read translation file
	if err := conf.LoadLocales("../conf/locales/zh-cn.yaml"); err != nil {
		util.Log().Panic("Failed to load translation file", err)
	}

	// Connect to database
	model.Database(os.Getenv("MYSQL_DSN"))
	cache.Redis()
}
