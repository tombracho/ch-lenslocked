package models

import (
	"database/sql"
	"fmt"
	"net/mail"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int
	Email        string
	Username     string
	PasswordHash string
}

//type NewUser struct {
//	Email        string
//	Username     string
//	PasswordHash string
//}

type UserService struct {
	DB *sql.DB
}

//func (us *UserService) Create(user *User) error {
//	//TODO
//	return nil
//}
//
//func (us *UserService) Create(user *NewUser) (*User, error) {
//	//TODO
//	return nil, nil
//}

func (us *UserService) Create(email, password, username string) (*User, error) {
	email = strings.ToLower(email)
	hashedBytes, err := bcrypt.GenerateFromPassword(
		[]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}
	passwordHash := string(hashedBytes)

	user := User{
		Email:        email,
		Username:     username,
		PasswordHash: passwordHash,
	}

	row := us.DB.QueryRow(`
		INSERT INTO users (email, username, password_hash)
		VALUES ($1, $2, $3) RETURNING id`, email, username, passwordHash)
	err = row.Scan(&user.ID)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}
	return &user, nil
}

func (us *UserService) Authenticate(username, password string) (*User, error) {
	var queryRow string
	var user User
	_, err := mail.ParseAddress(username)
	if err != nil {
		queryRow = `
			SELECT email, password_hash, id FROM users
			WHERE username=$1
		`
		user.Username = username
	} else {
		username = strings.ToLower(username)
		queryRow = `
			SELECT username, password_hash, id FROM users
			WHERE email=$1
		`
		user.Email = username
	}

	row := us.DB.QueryRow(queryRow, username)
	if len(user.Email) == 0 {
		err = row.Scan(&user.Email, &user.PasswordHash, &user.ID)
	} else {
		err = row.Scan(&user.Username, &user.PasswordHash, &user.ID)
	}
	if err != nil {
		return nil, fmt.Errorf("auth user: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("auth user: %w", err)
	}
	return &user, nil
}
