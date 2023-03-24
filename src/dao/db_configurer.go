package dao

import (
	"auth-service/src/config"
	"auth-service/src/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Configure(config *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("? Connected Successfully to the Database")
	//to fix 'Function uuid_generate_v4() does not exist'
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	// Migrate the schema
	db.AutoMigrate(&models.User{})
	return db
}
