package pflib

import (
	"fmt"
	"testing"
)

func TestMinimalWork_RemoveCollection(t *testing.T) {
	type args struct {
		collection string
	}
	tests := []struct {
		name string
		w    MinimalWork
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.RemoveCollection(tt.args.collection)
		})
	}
}

func TestRemove(t *testing.T) {
	minimalWork := MinimalWork{
		Name:        "Fishing Boats on the Beach at Les Saintes-Maries-de-la-Mer",
		Artist:      "Vincent van Gogh",
		Citation:    "https://www.britannica.com/biography/Vincent-van-Gogh/images-videos#/media/1/237118/229363",
		ImageURL:    "",
		Collections: []string{"first", "second"},
	}
	minimalWork.RemoveCollection("second")
	fmt.Printf("len %v", len(minimalWork.Collections))
	if len(minimalWork.Collections) != 1 || minimalWork.Collections[0] != "first" {
		t.Errorf("unexpected result %v", minimalWork.Collections)
	}

}
