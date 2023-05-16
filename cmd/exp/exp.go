package main

import (
	stdctx "context"
	"fmt"

	"github.com/tombracho/ch-lenslocked/context"
	"github.com/tombracho/ch-lenslocked/models"
)

type ctxKey string

const (
	favoriteColorKey ctxKey = "favorite-color"
)

func main() {
	ctx := stdctx.Background()

	user := models.User{
		Email: "artyomsonyx@gmail.com",
	}
	ctx = context.WithUser(ctx, &user)

	retrievedUser := context.User(ctx)
	fmt.Println(retrievedUser.Email)
}
