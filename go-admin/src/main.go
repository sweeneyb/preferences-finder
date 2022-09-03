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
	// props    map[string]string
	ID       string `firestore:"id,omitempty"`
	Name     string `firestore:"name,omitempty"`
	Citation string `firestore:"citation,omitempty"`
	ImageURL string `firestore:"imageURL,omitempty"`
	Artist   string `firestore:"artist,omitempty"`
}

type Collection struct {
	works []Work
}

func newWork(ref *firestore.DocumentSnapshot) *Work {
	var mapping map[string]string
	ref.DataTo(&mapping)
	w := Work{}
	w.Name = mapping["name"]
	w.Citation = mapping["citation"]
	w.ImageURL = mapping["imageURL"]
	w.Artist = mapping["artist"]
	return &w
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
		works = append(works, *newWork(docRef))
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

	w := Work{Name: "Fishing Boats on the Beach at Les Saintes-Maries-de-la-Mer",
		Artist:   "Vincent van Gogh",
		Citation: "https://www.britannica.com/biography/Vincent-van-Gogh/images-videos#/media/1/237118/229363",
		ImageURL: "/images/Fishing-Boats-on-the-Beach-oil-canvas-1888.jpg"}

	collection.works = append(collection.works, w)
	fmt.Printf("citation of 1th element: %v\n", collection.works[1].Citation)

	err = fsclient.addWork("first", &w, ctx)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("doc after add %v", w)
}

func (client Client) addWork(collection string, w *Work, ctx context.Context) error {
	doc, _, err := client.Collection("collections").Doc("TksLlbd0JskZZ0Bj0jvH").Collection(collection).Add(ctx, w)
	if err != nil {
		log.Printf("An error has occurred: %s", err)
		return err
	}
	w.ID = doc.ID
	return nil

}
