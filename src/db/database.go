package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func ConnectToDatabase() (*gorm.DB, error) {
	godotenv.Load(".env")
	DB_HOST := os.Getenv("DB_HOST")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")
	DB_SSL := os.Getenv("DB_SSL")
	DB_TIMEZONE := os.Getenv("DB_TIMEZONE")

	DSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=%s timezone=%s", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT, DB_SSL, DB_TIMEZONE)

	DB, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  DSN,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
	})
	if err != nil {
		fmt.Println("Cannot connect to database! ")
		log.Fatal("Connection error: ", err)
	}
	fmt.Println("Connected to the database!")

	return DB, nil
}

func Init() error {
	db, err := ConnectToDatabase()
	if err != nil {
		return err
	}
	DB = db
	return nil
}

func DropType(typestr string) bool {
	// Check if type exists
	var count int64
	query := fmt.Sprintf("SELECT COUNT(*) FROM pg_type WHERE typname = '%s';", typestr)
	if err := DB.Raw(query).Count(&count).Error; err != nil {
		fmt.Printf("Error checking type %s existence: %v\n", typestr, err)
		return false
	}
	if count == 0 {
		fmt.Printf("Type %s does not exist\n", typestr)
		return true
	}

	// Drop type if it exists
	query = fmt.Sprintf("DROP TYPE %s;", typestr)
	if err := DB.Exec(query).Error; err != nil {
		fmt.Printf("Error dropping type %s: %v\n", typestr, err)
		return false
	}
	fmt.Printf("Type %s dropped successfully\n", typestr)
	return true
}

func CreateType(typestr string, values string) bool {
	// Check if type already exists
	var count int64
	query := fmt.Sprintf("SELECT COUNT(*) FROM pg_type WHERE typname = '%s';", typestr)
	if err := DB.Raw(query).Count(&count).Error; err != nil {
		fmt.Printf("Error checking type %s existence: %v\n", typestr, err)
		return false
	}
	if count > 0 {
		fmt.Printf("Type %s already exists\n", typestr)
		return true
	}

	// Create type if it doesn't exist
	query = fmt.Sprintf("CREATE TYPE %s AS %s;", typestr, values)
	if err := DB.Exec(query).Error; err != nil {
		fmt.Printf("Error creating type %s: %v\n", typestr, err)
		return false
	}
	fmt.Printf("Type %s created successfully\n", typestr)
	return true
}
