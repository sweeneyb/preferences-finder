package main

import (
	"context"
	"fmt"
	"log"

	pflib "go-admin/pflib"
)

func main() {
	// Use a service account
	ctx := context.Background()
	fsclient, err := pflib.NewClient(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print("Bugger off")

	// collection := fsclient.GetCollection("first", ctx)
	// fmt.Printf("citation of 0th element: %v\n", collection.Works[0].Citation)

	// w := pflib.Work{Name: "Fishing Boats on the Beach at Les Saintes-Maries-de-la-Mer",
	// 	Artist:   "Vincent van Gogh",
	// 	Citation: "https://www.britannica.com/biography/Vincent-van-Gogh/images-videos#/media/1/237118/229363",
	// 	ImageURL: "/images/Fishing-Boats-on-the-Beach-oil-canvas-1888.jpg"}
	// collection.Works = append(collection.Works, w)
	// fmt.Printf("citation of 1th element: %v\n", collection.Works[1].Citation)

	// err = fsclient.AddWork("first", &w, ctx)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// log.Printf("doc after add %v", w)
	// err = fsclient.DeleteWork("first", &w, ctx)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// w := pflib.Work{Name: "Fishing Boats on the Beach at Les Saintes-Maries-de-la-Mer",
	// 	Artist:   "Vincent van Gogh",
	// 	Citation: "https://www.britannica.com/biography/Vincent-van-Gogh/images-videos#/media/1/237118/229363",
	// 	ImageURL: "/images/Fishing-Boats-on-the-Beach-oil-canvas-1888.jpg"}

	// secondCol := pflib.Collection{
	// 	Id:    "",
	// 	Name:  "second",
	// 	Works: []pflib.Work{w},
	// }
	// fsclient.AddCollection(&secondCol, ctx)

	fsclient.DeleteCollection("second", ctx)

}
