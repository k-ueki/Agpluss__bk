package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

type DB struct {
	Host    string
	User    string
	Port    int
	Pass    string
	DBName  string
	Protcol string
}

func NewDBStruct() *DB {
	if err := godotenv.Load(fmt.Sprintf("../.env")); err != nil {
		log.Fatal("Error loading .env file")
	}
	test := os.Getenv("name")
	fmt.Println("test,", test)

	return &DB{
		MS:      "",
		Host:    "",
		User:    "",
		Port:    0,
		Pass:    "",
		DBName:  "",
		Protcil: "",
	}
}

func (d *DB) Connect() *gorm.DB {
	DSN := fmt.Sprintf("%s:@%s(%s:%d)/%s", d.User, d.Pass, d.Protcol, d.Host, d.Port, d.DBName)

	db, err := gorm.Open(fmt.Printf("%s,%s", d.MS, DSN))
	if err != nil {
		return nil
	}

	return db
}

func main() {}
