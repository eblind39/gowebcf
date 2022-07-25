package main

import (
	"context"
	"fmt"
	"github.com/eblind39/gowebcf/go-webarch/session03/session"
)

func main() {
	ctx := context.Background()
	ctx = session.WithUserID(ctx, 1)

	userID := session.GetUserID(ctx)
	if userID == nil {
		fmt.Println("Not logged in")
		return
	}
	fmt.Println(*userID)
}
