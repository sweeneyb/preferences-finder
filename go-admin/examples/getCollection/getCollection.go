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

	collection, err := fsclient.GetCollection("third", ctx)
	if err != nil {
		log.Fatal(err)
	}
	for i, v := range collection.Works {
		fmt.Printf("citation of %vth element: %v\n", i, v.Citation)
		fmt.Printf("%vth element: %v\n", i, v)
		fmt.Printf("collections: %v\n", v.Collections)
		fmt.Printf("coll size %v\n", len(v.Collections))

	}

}
