package conf

import (
	"fmt"
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
	// MYSQL_DSN="echarge:evc2023@tcp(103.176.179.98:3307)/singo?charset=utf8&parseTime=True&loc=Local"
	mysqlConn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s",
		os.Getenv("MYSQL_USERNAME"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
		"utf8",
		"True",
	)
	// fmt.Println("mysqlConn", mysqlConn)
	model.Database(mysqlConn)
	cache.Redis()
}
