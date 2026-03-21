package cli

import (
	"flag"
	"fmt"
	"os"
)

func ParseConfig() (string, string) {

	flag.Usage = usage

	searchTerm := flag.String("term", "", "Text to search for")

	flag.Parse()

	args := flag.Args()

	if len(args) > 1 {

		fmt.Errorf("more than 1 root file specified")
		flag.Usage()
		os.Exit(1)
	}

	rootPath := args[0]

	return *searchTerm, rootPath
}

func usage() {

	fmt.Fprintf(os.Stderr, "Usage: searcher -term <string> <path>")

	flag.PrintDefaults()
}
