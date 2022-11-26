package pflib

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
)

type MinimalWork struct {
	// props    map[string]string
	ID          string   `firestore:"id,omitempty"`
	Name        string   `firestore:"name,omitempty"`
	Citation    string   `firestore:"citation,omitempty"`
	ImageURL    string   `firestore:"imageURL,omitempty"`
	Artist      string   `firestore:"artist,omitempty"`
	Collections []string `firestore:Collections,omitempty`
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
	// var mapping map[string]interface{}
	// ref.DataTo(&mapping)
	// w := MinimalWork{}
	// w.Name = mapping["name"].(string)
	// w.Citation = mapping["citation"].(string)
	// w.ImageURL = mapping["imageURL"].(string)
	// w.Artist = mapping["artist"].(string)
	// w.Collections = mapping["Collections"].([]string)
	// fmt.Printf("collections: %v\n", mapping["Collections"])
	// return &w
	var w MinimalWork
	ref.DataTo(&w)
	return &w
}

func (client Client) AddWork(collection string, lwg WorkWithLocalPath, ctx context.Context) error {
	doc, _, err := client.FsClient.Collection("works").Add(ctx, lwg.MinimalWork)
	if err != nil {
		log.Printf("An error has occurred: %s", err)
		return err
	}
	lwg.ID = doc.ID
	fmt.Printf("%v,%v", lwg.ID, lwg.Path)
	return nil
}

func (w *MinimalWork) RemoveCollection(collection string) {
	for i, value := range w.Collections {
		if value == collection {
			w.Collections = append(w.Collections[:i], w.Collections[i+1:]...)
			break
		}
	}
}

func (client Client) DeleteWork(collection string, wg MinimalWork, ctx context.Context) error {
	_, err := client.FsClient.Collection("works").Doc(wg.ID).Delete(ctx)
	return err
}

func (w MinimalWork) IsInCollection(name string) bool {
	for _, a := range w.Collections {
		if a == name {
			return true
		}
	}
	return false
}

func (w MinimalWork) String() string {
	return fmt.Sprintf("%v: %v, %v", w.ID, w.Name, w.ImageURL)
}
