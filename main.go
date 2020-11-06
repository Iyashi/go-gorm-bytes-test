package main

import (
	"fmt"
	"os"

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

// this will work
type ModelRaw struct {
  Bytes   []byte
}

type Hash []byte

// this will NOT work
type ModelTyped struct {
  Bytes Hash
}

func main() {
	// Open database
	dialector := mysql.Open(fmt.Sprintf("%s:%s@(%s:%v)/%s?charset=utf8mb4", user, password, host, port, dbname))
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Migrate database
	if err := db.AutoMigrate(&ModelRaw{}, &ModelTyped{}); err != nil {
		panic(err)
	}

  // this works
  mdl1 := &ModelRaw{Bytes: []byte("random-bytes")}
  if err := db.Debug().Create(mdl1).Error; err != nil {
    panic(err)
  }

  // error here
  mdl2 := &ModelTyped{Bytes: Hash("random-bytes")}
  if err := db.Debug().Create(mdl2).Error; err != nil {
    panic(err)
  }
}
