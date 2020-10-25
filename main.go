package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database config
var (
	host     = os.Getenv("MYSQL_HOST")
	port     = os.Getenv("MYSQL_PORT")
	user     = os.Getenv("MYSQL_USER")
	password = os.Getenv("MYSQL_PASSWORD")
	dbname   = os.Getenv("MYSQL_DATABASE")
)

type PasswordHash []byte

// Model describes the model that should be inserted
type Model struct {
	Password PasswordHash
}

func main() {
	// Open database
	dialector := mysql.Open(fmt.Sprintf("%s:%s@(%s:%v)/%s?charset=utf8mb4", user, password, host, port, dbname))
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Migrate database
	if err := db.AutoMigrate(&Model{}); err != nil {
		panic(err)
	}

	// Insert raw bytes
	mdl := &Model{
		Password: PasswordHash("random-bytes"),
	}
	if err := db.Create(mdl).Error; err != nil {
		panic(err)
	}

	// Insert bcrypt hash
	hash, err := bcrypt.GenerateFromPassword([]byte("random-bytes"), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	mdl = &Model{
		Password: PasswordHash(hash),
	}
	if err := db.Create(mdl).Error; err != nil {
		panic(err)
	}
}
