package models

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/tombracho/ch-lenslocked/rand"
)

type TokenManager struct {
	BytesPerToken int
}

const (
	//The minimum number of bytes to be used for each session token
	MinBytePerToken = 32
)

func (tm TokenManager) New() (token string, err error) {
	if tm.BytesPerToken < MinBytePerToken {
		tm.BytesPerToken = MinBytePerToken
	}
	token, err = rand.String(tm.BytesPerToken)
	if err != nil {
		return "", fmt.Errorf("token manager new: %w", err)
	}
	return token, nil
}

func (tm TokenManager) Hash(token string) string {
	tokenHash := sha256.Sum256([]byte(token))
	//base64 encode the data into a string
	return base64.URLEncoding.EncodeToString(tokenHash[:])
}
