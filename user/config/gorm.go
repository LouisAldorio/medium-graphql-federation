package config

import (
	"fmt"
	"log"
	"os"
	"time"
	"database/sql"

	mysql "gorm.io/driver/mysql"
	gorm "gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func ConnectGorm()  (*gorm.DB, *sql.DB) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true", 
		os.Getenv("DB_USERNAME"), 
		os.Getenv("DB_PASSWORD"), 
		os.Getenv("DB_HOST"), 
		os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:         initLog(),
		NamingStrategy: initNamingStrategy(),
	})

	if err != nil {
		fmt.Println(err)
		panic("Fail To Connect Database")
	}

	sql, _ := db.DB()
	sql.SetMaxOpenConns(120)

	return db, sql
}

func initLog() logger.Interface {
	newLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		Colorful:      true,
		LogLevel:      logger.Info,
		SlowThreshold: time.Second,
	})
	return newLogger
}

func initNamingStrategy() *schema.NamingStrategy {
	return &schema.NamingStrategy{
		SingularTable: true,
		TablePrefix:   "",
	}
}
