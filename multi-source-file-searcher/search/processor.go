package search

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ProcessFile(filePath string, searchTerm string) error {

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, searchTerm) {
			fmt.Println(line)
		}
	}

	if scanner.Err() != nil {
		return err
	}

	return nil

}
