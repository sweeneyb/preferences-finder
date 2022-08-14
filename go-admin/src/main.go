package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
)

type Work struct {
	props    map[string]string
	Name     string
	Citation string
	ImageURL string
	Artist   string
}

type Collection struct {
	works []Work
}

func newWork(mapping map[string]string) *Work {
	w := Work{props: mapping}
	w.Name = mapping["name"]
	w.Citation = mapping["citation"]
	w.ImageURL = mapping["imageURL"]
	w.Artist = mapping["artist"]
	return &w
}

func newWorkFromDocRef(ref *firestore.DocumentSnapshot) *Work {
	var doc map[string]string
	ref.DataTo(&doc)
	return newWork(doc)
}

type Client struct {
	*firestore.Client
}

func (client Client) GetCollection(name string, ctx context.Context) *Collection {

	var works []Work
	docIter := client.Collection("collections").Doc("TksLlbd0JskZZ0Bj0jvH").Collection(name).Documents(ctx)
	for {
		docRef, err := docIter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		works = append(works, *newWorkFromDocRef(docRef))
	}
	collection := Collection{works: works}
	return &collection
}

func main() {
	// Use a service account
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: os.Getenv("project_id")}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	fsclient := Client{client}

	collection := fsclient.GetCollection("first", ctx)
	fmt.Printf("citation of 0th element: %v\n", collection.works[0].Citation)
}
