package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/msandim/web-hasher/hasher"
	"github.com/msandim/web-hasher/workerpool"
)

func parseArgs() (nWorkers int, urls []string, err error) {

	flag.IntVar(&nWorkers, "parallel", 10, "number of concurrent operations")
	flag.Parse()

	if nWorkers <= 0 {
		return -1, nil, fmt.Errorf("parallel needs to be > 0")
	}

	urls = flag.Args()

	if len(urls) == 0 {
		return -1, nil, fmt.Errorf("at least one url needs to be specified")
	}

	return
}

type hashResult struct {
	url  string
	hash string
	err  error
}

func hashUrls(urls []string, nWorkers int, hasher hasher.Hasher) (result []hashResult) {

	pool := workerpool.New(nWorkers)

	for _, url := range urls {
		pool.AddJob(hashJob{url: url, hasher: hasher})
	}

	pool.EndJobs()
	pool.Run()

	poolResults := pool.GetResultsChannel()

	for poolResult := range poolResults {
		url := poolResult.GetJob().(hashJob).url
		hash := poolResult.(hashJobResult).hash
		err := poolResult.(hashJobResult).err
		result = append(result, hashResult{url: url, hash: hash, err: err})
	}

	return
}

func main() {
	nWorkers, urls, err := parseArgs()

	if err != nil {
		fmt.Printf("parsing error: %s\n", err)
		os.Exit(-1)
	}

	results := hashUrls(urls, nWorkers, hasher.NewHTTPHasher(&hasher.HTTPClient{}))

	for _, result := range results {
		fmt.Println(result.url, result.hash)
	}
}
