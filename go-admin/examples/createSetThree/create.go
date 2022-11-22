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

	w := pflib.WorkWithLocalPath{
		MinimalWork: pflib.MinimalWork{
			Name:     "Fishing Boats on the Beach at Les Saintes-Maries-de-la-Mer",
			Artist:   "Vincent van Gogh",
			Citation: "https://www.britannica.com/biography/Vincent-van-Gogh/images-videos#/media/1/237118/229363",
			ImageURL: "",
		},
		Path: "..\\frontend\\public\\images\\Fishing-Boats-on-the-Beach-oil-canvas-1888.jpg"}

	err = fsclient.AddCollection("third", []pflib.WorkWithLocalPath{w,
		{
			MinimalWork: pflib.MinimalWork{
				Name:     "Dublin",
				Artist:   "Robert Ballagh",
				Citation: "https://www.catawiki.com/en/l/53843773-robert-ballagh-dublin",
				ImageURL: "",
			},
			Path: "..\\frontend\\public\\images\\Dublin.jpg",
		},
		{
			MinimalWork: pflib.MinimalWork{
				Name:     "Village by the water",
				Artist:   "Jean Luc Lecoindre",
				Citation: "https://www.catawiki.com/en/l/64175407-jean-luc-lecoindre-1932-village-au-bord-de-l-eau",
				ImageURL: "",
			},
			Path: "..\\frontend\\public\\images\\Village by the water.jpg",
		},
	}, ctx)
	if err != nil {
		log.Fatal(err)

	}

}
