package conf

import (
	"os"
	"singo/cache"
	"singo/model"
	"singo/util"

	"github.com/joho/godotenv"
)

// Init Initialize configuration
func Init() {
	// Load environment variables from local file
	godotenv.Load()

	// Set log level
	util.BuildLogger(os.Getenv("LOG_LEVEL"))

	// Load translation file
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		util.Log().Panic("Failed to load translation file", err)
	}

	// Connect to database
	model.Database(os.Getenv("MYSQL_DSN"))
	cache.Redis()
}
