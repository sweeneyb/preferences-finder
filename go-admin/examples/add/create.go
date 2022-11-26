package main

import (
	"context"
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
	w := pflib.WorkWithLocalPath{
		MinimalWork: pflib.MinimalWork{
			Name:        "Fishing Boats on the Beach at Les Saintes-Maries-de-la-Mer",
			Artist:      "Vincent van Gogh",
			Citation:    "https://www.britannica.com/biography/Vincent-van-Gogh/images-videos#/media/1/237118/229363",
			ImageURL:    "/images/Fishing-Boats-on-the-Beach-oil-canvas-1888.jpg",
			Collections: []string{"testColl"},
		},
		Path: "..\\frontend\\public\\images\\small.png"}

	err = fsclient.AddCollection("second", []pflib.WorkWithLocalPath{w}, ctx)
	if err != nil {
		log.Fatal(err)

	}

}
