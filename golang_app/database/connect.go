package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"
)

// ConnectDB connect to db
func ConnectDB() {

	// Postgres Connection
	db_url := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USERNAME"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DATABASE"))
	log.Println(db_url)
	_, err := sql.Open("postgres", db_url)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	//drv := entsql.OpenDB(dialect.Postgres, db)
	log.Println("Connection Opened to Database")
	// var err error
	// p := config.Config("DB_PORT")
	// port, err := strconv.ParseUint(p, 10, 32)

	// if err != nil {
	// 	panic("failed to parse database port")
	// }

	// dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
	// DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// if err != nil {
	// 	panic("failed to connect database")
	// }

	// fmt.Println("Connection Opened to Database")
	// DB.AutoMigrate(&model.Product{}, &model.User{})
	// fmt.Println("Database Migrated")
}
