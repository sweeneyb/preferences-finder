package pflib

import (
	"context"
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

	var works []Work
	docIter := client.FsClient.Collection("collections").Doc(client.RootDoc).Collection(name).Documents(ctx)
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
	collection := Collection{Works: works}
	return &collection
}

func (client Client) GetWorks(collectionName string, ctx context.Context) []Work {

	var works []Work
	docIter := client.FsClient.Collection("collections").Doc(client.RootDoc).Collection(collectionName).Documents(ctx)
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

func (client Client) AddCollection(collection *Collection, ctx context.Context) error {
	for _, value := range collection.Works {
		err := client.AddWork(collection.Name, &value, ctx)
		if err != nil {
			log.Printf("An error has occurred: %s", err)
			return err
		}

	}
	return nil
}
