package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_DATABASE"),
	}
	return &dbConfig
}
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
func NewDB(params ...string) *gorm.DB {
	var err error
	conString := DbURL(BuildDBConfig())

	DB, err = gorm.Open(mysql.Open(conString), &gorm.Config{})

	// sqlDB, err := DB.DB()

	// sqlDB.SetMaxIdleConns(10)
	// sqlDB.SetMaxOpenConns(100)
	// sqlDB.SetConnMaxLifetime(time.Hour)

	// defer sqlDB.Close()

	if err != nil {
		log.Panic(err)
	}
	return DB
}

func GetDBInstance() *gorm.DB {
	return DB
}
