package pflib

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type Client struct {
	FsClient *firestore.Client
	RootDoc  string
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

func (client Client) GetWorks(name string, ctx context.Context) []Work {

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
	return works
}
