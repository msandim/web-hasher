package main

import (
	"flag"
	"fmt"
	"os"
)

func cenas() {

}

func parseArgs() (parallel int, urls []string, err error) {

	flag.IntVar(&parallel, "parallel", 3, "Number of concurrent operations")
	flag.Parse()

	if parallel <= 0 {
		return -1, nil, fmt.Errorf("Parallel needs to be > 0")
	}

	urls = flag.Args()

	if len(urls) == 0 {
		return -1, nil, fmt.Errorf("At least one url needs to be specified")
	}

	return
}

func main() {
	_, _, err := parseArgs()

	if err != nil {
		fmt.Printf("Parsing error: %s\n", err)
		os.Exit(-1)
	}

}
