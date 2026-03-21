package search

import (
	"os"
	"testing"
)

func TestProcessFile_Success(t *testing.T) {

	tempFile, err := os.CreateTemp("./", "*")
	if err != nil {
		t.Errorf("Error creating temp file: %v", err)
	}

	tempFile.WriteString("Line 1")
	tempFile.WriteString("Line 2")
	tempFile.WriteString("find-me")

	if err := ProcessFile(tempFile.Name(), "find-me"); err != nil {
		t.Errorf("Process file failed: %v", err)
	}

	//tempfile.Name() returns the filepath of the temporary file

}
