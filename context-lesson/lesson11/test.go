package lesson11

import (
	"context"
	"fmt"
)

type contextKey string

const requestIdKey = contextKey("id")
const userKey = contextKey("user")

func Test() {
	work := func() int {
		return 42
	}

	ctx := context.WithValue(context.Background(), requestIdKey, 1234)
	ctx = context.WithValue(ctx, userKey, "admin")

	res := execute(ctx, work)
	fmt.Println(res)

	ctx = context.Background()
	res = execute(ctx, work)
	fmt.Println(res)
}

func execute(ctx context.Context, fn func() int) int {
	reqId := ctx.Value(requestIdKey)

	if reqId != nil {
		fmt.Printf("Request ID = %d \n", reqId)
	} else {
		fmt.Println("Request ID unknown")
	}

	user := ctx.Value(userKey)

	if user != nil {
		fmt.Printf("User is %s \n", user)
	} else {
		fmt.Println("Request user unknown")
	}

	return fn()
}
