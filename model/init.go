package model

import (
	"log"
	"os"
	"singo/util"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB Database connection singleton
var DB *gorm.DB

// Database Initialize mysql connection in middleware
func Database(connString string) {
	// Initialize GORM log configuration
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level(Remember to change it according to your needs)
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{
		Logger: newLogger,
	})
	// Error
	if connString == "" || err != nil {
		util.Log().Error("mysql lost: %v", err)
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		util.Log().Error("mysql lost: %v", err)
		panic(err)
	}

	// Set connection pool
	// Idle
	sqlDB.SetMaxIdleConns(10)
	// Open
	sqlDB.SetMaxOpenConns(20)
	DB = db

	migration()
}
