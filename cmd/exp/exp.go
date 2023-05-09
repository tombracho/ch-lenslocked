package main

import (
	"fmt"

	"github.com/tombracho/ch-lenslocked/models"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

type Order struct {
	ID          uint `gorm:"primaryKey"`
	UserID      uint
	Amount      int
	Description string
}

type User struct {
	ID     uint `gorm:"primaryKey"`
	Name   string
	Email  string
	Orders []Order
}

func main() {
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected!")
	us := models.UserService{
		DB: db,
	}
	user, err := us.Create("katrinmail@gmail.com", "21312012KMG", "KMG")
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
}
