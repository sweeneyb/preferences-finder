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
		ref, err := docRef.Ref.Get(ctx)
		if err != nil {
			return err
		}
		MinimalWork := newWork(ref)
		ctx, cancel := context.WithTimeout(ctx, time.Second*50)
		defer cancel()
		bucket, err := client.Storage.Bucket(os.Getenv("project_id") + ".appspot.com")
		if err != nil {
			log.Printf("An error has occurred...: %s", err)
			return err
		}

		err = bucket.Object(strings.TrimPrefix(MinimalWork.ImageURL, "/")).Delete(ctx)
		if err != nil {
			log.Printf("An error has occurred...: %s", err)
			return err
		}

		docRef.Ref.Delete(ctx)
	}
	return nil
}

func (client Client) AddCollection(name string, lwgs []LocalWorkGetter, ctx context.Context) error {
	coll := client.FsClient.Collection("collections").Doc(client.RootDoc).Collection(name)
	for _, value := range lwgs {
		f, err := os.Open(value.GetPath())
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
		fmt.Printf("Blob %v uploaded.\n", "foo")

		value.GetWork().ImageURL = fmt.Sprintf("/%v", uuid)
		// below works
		ref, _, err := coll.Add(ctx, value.GetWork())
		if err != nil {
			log.Printf("An error has occurred: %s", err)
			return err
		}
		tmp := value.GetWork()
		tmp.ID = ref.ID
		// tmp.ImageURL = fmt.Sprintf("/%v", uuid)
		fmt.Printf("%v,%v", value.GetWork().ID, value.GetPath())
	}

	return nil

}
