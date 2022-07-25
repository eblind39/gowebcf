package main

import (
	"context"
	"fmt"
	"github.com/eblind39/gowebcf/go-webarch/session03/session"
	"github.com/eblind39/gowebcf/go-webarch/session03/session02"
)

func main() {
	// ________________direct_sol_________________
	ctx := context.Background()
	ctx = context.WithValue(ctx, "userID", 12345)
	userId, ok := ctx.Value("userID").(int)
	if !ok {
		panic("Error on get userID")
	}

	fmt.Println("userID: ", userId)

	// _________________session___________________

	ctx = context.Background()
	ctx = session.WithUserID(ctx, 1)

	userID := session.GetUserID(ctx)
	if userID == nil {
		fmt.Println("Not logged in")
		return
	}
	fmt.Println(*userID)

	// _______________session02___________________
	ctx = context.Background()
	ctx = session02.SetUsetID(ctx, 1)
	ctx = session02.SetAdminAccess(ctx, true)

	uID := session02.GetUserID(ctx)
	isAdmin := session02.GetAdmin(ctx)

	fmt.Printf("User %d is an admin %t", uID, isAdmin)
}
