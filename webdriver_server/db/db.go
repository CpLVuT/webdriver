package databases

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"webdriver_server/models"
)

var DB *gorm.DB

func InitDB(user string, pass string, addr string, port string, dbName string) *gorm.DB {
	// Connect to DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, addr, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// Migrate DB
	Migrate(db)
	DB = db
	return db
}

func Migrate(db *gorm.DB) {
	models.ManuscriptMigrate(db)
	models.UserMigrate(db)
}
