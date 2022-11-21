package pflib

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
)

type WorkGetter interface {
	GetWork() *MinimalWork

	// Name() string
	// Citation() string
	// ImageURL() string
	// Artist() string
}

type LocalWorkGetter interface {
	WorkGetter
	GetPath() string
}

type SimpleWorkGetter struct {
	Work *MinimalWork
}

func (sg SimpleWorkGetter) GetWork() *MinimalWork {
	return sg.Work
}

type LocaleWorkGetter struct {
	Work *MinimalWork
	Path string
}

func (lwg LocaleWorkGetter) GetWork() *MinimalWork {
	return lwg.Work
}

func (lwg LocaleWorkGetter) GetPath() string {
	return lwg.Path
}

type MinimalWork struct {
	// props    map[string]string
	ID       string `firestore:"id,omitempty"`
	Name     string `firestore:"name,omitempty"`
	Citation string `firestore:"citation,omitempty"`
	ImageURL string `firestore:"imageURL,omitempty"`
	Artist   string `firestore:"artist,omitempty"`
}

func (w MinimalWork) GetID() string {
	return w.ID
}

type Collection struct {
	Id    string       `firestore:"id,omitempty"`
	Name  string       `firestore:"name,omitempty"`
	Works []WorkGetter `firestore:"works,omitempty"`
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

func (client Client) AddWork(collection string, lwg LocalWorkGetter, ctx context.Context) error {
	doc, _, err := client.FsClient.Collection("collections").Doc(client.RootDoc).Collection(collection).Add(ctx, lwg.GetWork())
	if err != nil {
		log.Printf("An error has occurred: %s", err)
		return err
	}
	tmp := lwg.GetWork()
	tmp.ID = doc.ID
	fmt.Printf("%v,%v", lwg.GetWork().ID, lwg.GetPath())
	return nil
}

func (client Client) DeleteWork(collection string, wg WorkGetter, ctx context.Context) error {
	_, err := client.FsClient.Collection("collections").Doc(client.RootDoc).Collection(collection).Doc(wg.GetWork().ID).Delete(ctx)
	return err
}

func newCollection(ref *firestore.DocumentSnapshot) *Collection {
	var mapping map[string]interface{}
	ref.DataTo(&mapping)
	c := Collection{}
	c.Id = fmt.Sprintf("%v", mapping["id"])

	return &c
}
