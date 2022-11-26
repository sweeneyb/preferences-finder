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

	err = fsclient.DeleteCollection("second", ctx)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print("Deleted")

}
