package search

import (
	"os"
	"path/filepath"
	"sync"
)

func Search(rootPath string, searchTerm string, wg *sync.WaitGroup) error {

	files, err := os.ReadDir(rootPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			full_path := filepath.Join(rootPath, file.Name())

			err := Search(full_path, searchTerm, wg)
			if err != nil {
				return err
			}
		} else {

			full_path := filepath.Join(rootPath, file.Name())
			wg.Add(1)

			go func(path string) error {
				defer wg.Done()
				err := ProcessFile(path, searchTerm)

				if err != nil {
					return err
				}

				return nil
			}(full_path)

		}
	}

	return nil

}
