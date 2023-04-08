package main

import (
	"errors"
	"fmt"
	"log"
)

func Connect() error {
	return errors.New("connection faild")
}

func CreateUser() error {
	err := Connect()
	if err != nil {
		return fmt.Errorf("create user: %w", err)
	}
	return nil
}

func CreaeteOrg() error {
	err := CreateUser()
	if err != nil {
		return fmt.Errorf("create org: %w", err)
	}
	return nil
}

func main() {
	err := CreateUser()
	if err != nil {
		log.Println(err)
	}
	err = CreaeteOrg()
	if err != nil {
		log.Println(err)
	}
}
