package search

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSearch_Recursive(t *testing.T) {
	root := t.TempDir()

	level1Path := filepath.Join(root, "level1")
	err := os.Mkdir(level1Path, 0755)

	if err != nil {
		t.Errorf("Error creating level 1 path: %v", err)
	}

	filePath := filepath.Join(level1Path, "found.txt")
	err = os.WriteFile(filePath, []byte("hello world"), 0644)

	if err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}

	err = Search(root, "hello")
	if err != nil {
		t.Fatalf("Search failed: %v", err)
	}

}
