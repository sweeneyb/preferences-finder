package pflib

import (
	"context"
	"fmt"
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

type LocalWork struct {
	Work
	Path string
}

type Collection struct {
	Id    string `firestore:"id,omitempty"`
	Name  string `firestore:"name,omitempty"`
	Works []Work `firestore:"works,omitempty"`
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

func (client Client) AddWork(collection string, w *LocalWork, ctx context.Context) error {
	doc, _, err := client.FsClient.Collection("collections").Doc(client.RootDoc).Collection(collection).Add(ctx, w)
	if err != nil {
		log.Printf("An error has occurred: %s", err)
		return err
	}
	w.ID = doc.ID
	return nil
}

func (client Client) DeleteWork(collection string, w *Work, ctx context.Context) error {
	_, err := client.FsClient.Collection("collections").Doc(client.RootDoc).Collection(collection).Doc(w.ID).Delete(ctx)
	return err
}

func newCollection(ref *firestore.DocumentSnapshot) *Collection {
	var mapping map[string]interface{}
	ref.DataTo(&mapping)
	c := Collection{}
	c.Id = fmt.Sprintf("%v", mapping["id"])

	return &c
}

// func (client Client) DeleteCollection(collection *Collection, ctx context.Context) error {
// 	_, err := client.FsClient.Collection("collections").Doc(client.RootDoc).Collection(collection.Id).
// }
