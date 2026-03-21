package main

import (
	"fmt"
	"github.com/Joshua-Pok/multi-source-file-searcher/cli"
	"github.com/Joshua-Pok/multi-source-file-searcher/search"
	"sync"
)

func main() {

	searchTerm, rootPath := cli.ParseConfig()

	wg := &sync.WaitGroup{}

	err := search.Search(rootPath, searchTerm, wg)
	if err != nil {
		fmt.Errorf("Error searching path: %v", err)
	}

	wg.Wait()

	return
}
