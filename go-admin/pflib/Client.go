package pflib

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"

	firebase "firebase.google.com/go"
	fsStorage "firebase.google.com/go/storage"
	"google.golang.org/api/iterator"

	"strings"

	"github.com/google/uuid"
)

type Client struct {
	FsClient *firestore.Client
	Storage  *fsStorage.Client
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
		Storage:  storageClient,
		RootDoc:  "TksLlbd0JskZZ0Bj0jvH",
	}

	return &theClient, nil
}

func (client Client) GetCollection(name string, ctx context.Context) (*Collection, error) {

	works, err := client.GetWorks(name, ctx)
	if err != nil {
		return nil, err
	}
	collection := Collection{Works: works}
	return &collection, nil
}

func (client Client) GetWorks(collectionName string, ctx context.Context) ([]MinimalWork, error) {
	var works []MinimalWork
	iter := client.FsClient.Collection("works").Where("collections", "array-contains-any", []string{collectionName}).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		works = append(works, *newWork(doc))
	}
	return works, nil
}

func (client Client) DeleteCollection(collectionName string, ctx context.Context) error {
	docIter := client.FsClient.Collection("works").Where("collections", "array-contains-any", []string{collectionName}).Documents(ctx)
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()
	bucket, err := client.Storage.Bucket(os.Getenv("project_id") + ".appspot.com")
	if err != nil {
		log.Printf("An error has occurred...: %s", err)
		return err
	}

	for {
		docRef, err := docIter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		ref, err := docRef.Ref.Get(ctx)
		if err != nil {
			return err
		}
		minimalWork := newWork(ref)

		minimalWork.RemoveCollection(collectionName)
		fmt.Printf("name: %v\n", minimalWork.Name)
		fmt.Printf("url: %v\n", minimalWork.ImageURL)

		fmt.Printf("collections: %v\n", strings.Join(minimalWork.Collections, ", "))
		fmt.Printf("coll size %v\n", len(minimalWork.Collections))
		if len(minimalWork.Collections) == 0 {
			err = bucket.Object(strings.TrimPrefix(minimalWork.ImageURL, "/")).Delete(ctx)
			if err != nil {
				log.Printf("An error has occurred....: %s", err)
				return err
			}

			docRef.Ref.Delete(ctx)
		} else {
			docRef.Ref.Set(ctx, minimalWork)
		}
		// else update to remove the collection
	}
	return nil
}

func (client Client) AddCollection(name string, lwgs []WorkWithLocalPath, ctx context.Context) error {
	coll := client.FsClient.Collection("works")
	for _, value := range lwgs {
		f, err := os.Open(value.Path)
		if err != nil {
			return fmt.Errorf("os.Open: %v", err)
		}
		defer f.Close()
		ctx, cancel := context.WithTimeout(ctx, time.Second*50)
		defer cancel()
		bucket, err := client.Storage.Bucket(os.Getenv("project_id") + ".appspot.com")

		if err != nil {
			log.Printf("An error has occurred...: %s", err)
			return err
		}

		uuid := uuid.New()
		o := bucket.Object(uuid.String())
		o = o.If(storage.Conditions{DoesNotExist: true})
		wc := o.NewWriter(ctx)
		if _, err = io.Copy(wc, f); err != nil {
			return fmt.Errorf("io.Copy: %v", err)
		}
		if err := wc.Close(); err != nil {
			return fmt.Errorf("Writer.Close: %v", err)
		}

		value.ImageURL = fmt.Sprintf("/%v", uuid)
		if !value.IsInCollection(name) {
			value.Collections = append(value.Collections, name)
		}
		// below works
		ref, _, err := coll.Add(ctx, value)
		if err != nil {
			log.Printf("An error has occurred: %s", err)
			return err
		}
		tmp := value
		tmp.ID = ref.ID
		// tmp.ImageURL = fmt.Sprintf("/%v", uuid)
		fmt.Printf("%v,%v", value.ID, value.Path)
	}

	return nil

}
