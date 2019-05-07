package main

import (
	"flag"
	"fmt"
	"os"

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

type hashJob struct {
	url string
	hasher
}

func (job *hashJob) Process() workerpool.JobResult {
	hasher
}

func hashUrls(urls []string, nWorkers int) {

	var results []string

	pool := workerpool.New(nWorkers)

	for _, url := range urls {
		pool.AddJob()
	}
}

func main() {
	nWorkers, urls, err := parseArgs()

	if err != nil {
		fmt.Printf("parsing error: %s\n", err)
		os.Exit(-1)
	}

	hashUrls(urls, nWorkers)
}
