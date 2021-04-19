package utils

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jroopa/offers/pkg/models"
)

var db *gorm.DB
var err error

func ConnectDB() *gorm.DB {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Print("Error loading env file")
	// }

	// username := os.Getenv("databaseUser")
	// password := os.Getenv("databasePassword")
	// port := os.Getenv("databasePort")
	// databaseName := os.Getenv("databaseName")
	// databaseHost := os.Getenv("databaseHost")

	db, err = gorm.Open("mysql", "root:secret@tcp(127.0.0.1:3307)/snippetbox?parseTime=true")
	if err != nil {
		log.Fatal("DB connection error")
	}
	err = db.DB().Ping()
	if err != nil {
		log.Fatal(err)

	}
	// if no error. Ping is successful
	fmt.Println("Ping to database successful, connection is still alive")

	// Migrate the schema
	db.AutoMigrate(
		&models.User{},
		&models.Offers{},
	)

	fmt.Println("Successfully connected!")
	return db
}
