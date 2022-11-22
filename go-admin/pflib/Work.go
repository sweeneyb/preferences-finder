package pflib

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
)

type MinimalWork struct {
	// props    map[string]string
	ID       string `firestore:"id,omitempty"`
	Name     string `firestore:"name,omitempty"`
	Citation string `firestore:"citation,omitempty"`
	ImageURL string `firestore:"imageURL,omitempty"`
	Artist   string `firestore:"artist,omitempty"`
}

type WorkWithLocalPath struct {
	MinimalWork
	Path string `firestore:"-"`
}

func (w MinimalWork) GetID() string {
	return w.ID
}

type Collection struct {
	Id    string        `firestore:"id,omitempty"`
	Name  string        `firestore:"name,omitempty"`
	Works []MinimalWork `firestore:"works,omitempty"`
}

func newWork(ref *firestore.DocumentSnapshot) *MinimalWork {
	var mapping map[string]string
	ref.DataTo(&mapping)
	w := MinimalWork{}
	w.Name = mapping["name"]
	w.Citation = mapping["citation"]
	w.ImageURL = mapping["imageURL"]
	w.Artist = mapping["artist"]
	return &w
}

func (client Client) AddWork(collection string, lwg WorkWithLocalPath, ctx context.Context) error {
	doc, _, err := client.FsClient.Collection("collections").Doc(client.RootDoc).Collection(collection).Add(ctx, lwg.MinimalWork)
	if err != nil {
		log.Printf("An error has occurred: %s", err)
		return err
	}
	lwg.ID = doc.ID
	fmt.Printf("%v,%v", lwg.ID, lwg.Path)
	return nil
}

func (client Client) DeleteWork(collection string, wg MinimalWork, ctx context.Context) error {
	_, err := client.FsClient.Collection("collections").Doc(client.RootDoc).Collection(collection).Doc(wg.ID).Delete(ctx)
	return err
}
