package models

import (
	"database/sql"
	"fmt"
)

type Session struct {
	ID     int
	UserID int
	// Token is only set when creating a new session. When looking up a session
	// this will be left empty, as we only store the hash of a session token
	// in our database and we cannot reverse it into a raw token.
	Token     string
	TokenHash string
}

type SessionService struct {
	DB *sql.DB
	TM *TokenManager
}

// Create will create a new session for the user provided. The session token
// will be returned as the Token field on the Session type, but only the hashed
// session token is stored in the database.
func (ss *SessionService) Create(userID int) (*Session, error) {
	//If TokenManager isn`t ctreated, create it
	if ss.TM == nil {
		ss.TM = &TokenManager{}
	}
	token, err := ss.TM.New()
	if err != nil {
		return nil, fmt.Errorf("session token create: %w", err)
	}
	session := Session{
		UserID:    userID,
		Token:     token,
		TokenHash: ss.TM.Hash(token),
	}
	row := ss.DB.QueryRow(`
		INSERT INTO sessions (user_id, token_hash)
		VALUES ($1, $2) ON CONFLICT (user_id) DO
		UPDATE
		SET token_hash = $2
		RETURNING id;`, session.UserID, session.TokenHash)
	err = row.Scan(&session.ID)
	if err != nil {
		return nil, fmt.Errorf("session token create: %w", err)
	}
	return &session, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	if ss.TM == nil {
		ss.TM = &TokenManager{}
	}
	tokenHash := ss.TM.Hash(token)
	row := ss.DB.QueryRow(`
	SELECT email,
		username,
		user_id,
		password_hash 
	FROM users
	    JOIN sessions ON users.id = sessions.user_id
	    WHERE token_hash=$1;`, tokenHash)
	user := User{}
	err := row.Scan(&user.Email, &user.Username, &user.ID, &user.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("user session token: %w", err)
	}
	return &user, nil
}

func (ss *SessionService) Delete(token string) error {
	if ss.TM == nil {
		ss.TM = &TokenManager{}
	}
	tokenHash := ss.TM.Hash(token)
	_, err := ss.DB.Exec(`
		DELETE FROM sessions
		WHERE token_hash = $1`, tokenHash)
	if err != nil {
		return fmt.Errorf("delete token: %w", err)
	}
	return nil
}
