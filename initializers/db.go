package initializers

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(config *DBConfig) {
	var err error
	var dsn string

	switch config.Environment {
	case "local-win":
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tashkent", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)
	case "host-dev":
		dsn = config.ExternalDBURL
	case "prod":
		dsn = config.InternalDBURL
	}

	//if config.Environment == "local-win" {
	//	dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)
	//} else if config.Environment == "prod" {
	//	dsn = config.InternalDBURL
	//}

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("? Connected Successfully to the Database")
}
