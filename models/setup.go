package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB instnce
var DB *gorm.DB

// ConnectDataBase ...
func ConnectDataBase() {
	database, err := gorm.Open(sqlite.Open("./imfk.db"), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&MembershipFunc{})

	DB = database
}
