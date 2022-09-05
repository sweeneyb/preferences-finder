package pflib

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
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
	Works []Work
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

func (client Client) AddWork(collection string, w *Work, ctx context.Context) error {
	doc, _, err := client.Collection("collections").Doc("TksLlbd0JskZZ0Bj0jvH").Collection(collection).Add(ctx, w)
	if err != nil {
		log.Printf("An error has occurred: %s", err)
		return err
	}
	w.ID = doc.ID
	return nil
}

func (client Client) DeleteWork(collection string, w *Work, ctx context.Context) error {
	_, err := client.Collection("collections").Doc("TksLlbd0JskZZ0Bj0jvH").Collection(collection).Doc(w.ID).Delete(ctx)
	return err
}
