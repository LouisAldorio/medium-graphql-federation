package config

import (
    "database/sql"
	"os"
	"fmt"
    
    _ "github.com/go-sql-driver/mysql"
    "github.com/golang-migrate/migrate"
    "github.com/golang-migrate/migrate/database/mysql"
    _ "github.com/golang-migrate/migrate/source/file"
)


func Migrate() *migrate.Migrate {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true", 
		os.Getenv("DB_USERNAME"), 
		os.Getenv("DB_PASSWORD"), 
		os.Getenv("DB_HOST"), 
		os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))

	fmt.Println(dsn)

	db, err := sql.Open("mysql", dsn)
    
    if err != nil {
        panic(err)
    }
    driver, _ := mysql.WithInstance(db, &mysql.Config{})
    m, err := migrate.NewWithDatabaseInstance(
        "file://migrations/",
        "mysql", 
        driver,
    )

    if err != nil {
        fmt.Println(err)
    }

    return m
}

func MigrateUp() {
    migrator := Migrate()
    if err := migrator.Up(); err != nil {
        fmt.Println(err)
		panic(err)
    }
}

func MigrateDown() {
    migrator := Migrate()
    if err := migrator.Down(); err != nil {
        fmt.Println(err)
        panic(err)
    }
}
