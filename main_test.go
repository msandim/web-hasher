package main

import (
	"reflect"
	"sort"
	"testing"

	"github.com/msandim/web-hasher/hasher"
)

// Code needed to sort a slice of hash results:
type ByURL []hashResult

func (a ByURL) Len() int           { return len(a) }
func (a ByURL) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByURL) Less(i, j int) bool { return a[i].url < a[j].url }

func TestHashUrls(t *testing.T) {

	urls := []string{"aa.pt", "bb.pt", "cc.pt"}

	hashMap := map[string]string{
		"aa.pt": "31a43f7cf3cdb223ea15f9f7d9da9e89",
		"bb.pt": "92ccb58fccd74b971492f0c41d6cee0e",
		"cc.pt": "51794ef89c0ce333b16a787765862dd0",
	}

	wantedResults := []hashResult{
		hashResult{
			url:  "aa.pt",
			hash: "31a43f7cf3cdb223ea15f9f7d9da9e89",
			err:  nil,
		},
		hashResult{
			url:  "bb.pt",
			hash: "92ccb58fccd74b971492f0c41d6cee0e",
			err:  nil,
		},
		hashResult{
			url:  "cc.pt",
			hash: "51794ef89c0ce333b16a787765862dd0",
			err:  nil,
		},
	}

	hasher := hasher.NewMockHasher(hashMap)

	results := hashUrls(urls, 2, hasher)
	sort.Sort(ByURL(results))

	if !reflect.DeepEqual(results, wantedResults) {
		t.Fatalf("The results obtained were not as expected. Got this:\n%+v\n", results)
	}
}
