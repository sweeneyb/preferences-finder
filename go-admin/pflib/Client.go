package pflib

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/firestore"

	firebase "firebase.google.com/go"
	fsStorage "firebase.google.com/go/storage"
	"google.golang.org/api/iterator"
)

type Client struct {
	FsClient *firestore.Client
	foo      *fsStorage.Client
	RootDoc  string
}

func NewClient(ctx context.Context) (client *Client, err error) {
	conf := &firebase.Config{ProjectID: os.Getenv("project_id")}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		return nil, err
	}
	fsClient, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}

	storageClient, err := app.Storage(ctx)
	if err != nil {
		return nil, err
	}

	theClient := Client{
		FsClient: fsClient,
		foo:      storageClient,
		RootDoc:  "TksLlbd0JskZZ0Bj0jvH",
	}

	return &theClient, nil
}

func (client Client) GetCollection(name string, ctx context.Context) *Collection {

	var works []WorkGetter
	docIter := client.FsClient.Collection("collections").Doc(client.RootDoc).Collection(name).Documents(ctx)
	for {
		docRef, err := docIter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		works = append(works, SimpleWorkGetter{Work: newWork(docRef)})
	}
	collection := Collection{Works: works}
	return &collection
}

func (client Client) GetWorks(collectionName string, ctx context.Context) []WorkGetter {

	var works []WorkGetter
	docIter := client.FsClient.Collection("collections").Doc(client.RootDoc).Collection(collectionName).Documents(ctx)
	for {
		docRef, err := docIter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		works = append(works, SimpleWorkGetter{Work: newWork(docRef)})
	}
	return works
}

func (client Client) DeleteCollection(collectionName string, ctx context.Context) error {
	docIter := client.FsClient.Collection("collections").Doc(client.RootDoc).Collection(collectionName).Documents(ctx)
	for {
		docRef, err := docIter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		docRef.Ref.Delete(ctx)
	}
	return nil
}

func (client Client) AddCollection(name string, lwgs []LocalWorkGetter, ctx context.Context) error {
	coll := client.FsClient.Collection("collections").Doc(client.RootDoc).Collection(name)
	for _, value := range lwgs {
		ref, _, err := coll.Add(ctx, value.GetWork())
		if err != nil {
			log.Printf("An error has occurred: %s", err)
			return err
		}
		tmp := value.GetWork()
		tmp.ID = ref.ID
		fmt.Printf("%v,%v", value.GetWork().ID, value.GetPath())
	}

	return nil

}
