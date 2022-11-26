package main

import (
	"context"
	"fmt"
	"log"

	pflib "go-admin/pflib"
)

func main() {
	// Use a service account
	ctx := context.Background()
	fsclient, err := pflib.NewClient(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Bugger off")

	err = fsclient.DeleteCollection("third", ctx)
	if err != nil {
		log.Fatalln(err)
	}
}
