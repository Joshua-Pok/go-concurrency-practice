package search

import (
	"os"
	"path/filepath"
)

func Search(rootPath string, searchTerm string) error {

	files, err := os.ReadDir(rootPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			full_path := filepath.Join(rootPath, file.Name())

			err := Search(full_path, searchTerm)
			if err != nil {
				return err
			}
		} else {
			full_path := filepath.Join(rootPath, file.Name())

			err := ProcessFile(full_path, searchTerm)

			if err != nil {
				return err
			}
		}
	}

	return nil

}
