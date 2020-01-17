package fs

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestCloseFile(t *testing.T) {
	file := createTempFile(t)
	defer os.Remove(file.Name())
	CloseFile(file)

	_, err := file.Stat()
	if !strings.Contains(err.Error(), "use of closed file") {
		t.Errorf("CloseFile(%s) did not close the file", file.Name())
	}
}

func TestCloseFileTwice(t *testing.T) {
	if os.Getenv("CLOSE_FILE_EXIT") == "1" {
		file := createTempFile(t)
		defer os.Remove(file.Name())
		CloseFile(file)
		CloseFile(file)
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestCloseFileTwice")
	cmd.Env = append(os.Environ(), "CLOSE_FILE_EXIT=1")
	err := cmd.Run()

	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}

func TestDirExists(t *testing.T) {
	dir := createTempDir(t)
	defer os.RemoveAll(dir)
	exists, err := DirExists(dir)
	if !exists && err == nil {
		t.Errorf("DirExists(%s) = %t", dir, exists)
	}
}

func TestFileExists(t *testing.T) {
	file := createTempFile(t)
	defer os.Remove(file.Name())
	exists, err := FileExists(file.Name())
	if !exists && err == nil {
		t.Errorf("FileExists(%s) = %t", file.Name(), exists)
	}
}

func createTempFile(t *testing.T) *os.File {
	file, err := ioutil.TempFile(".", "tempFile")
	if err != nil {
		t.Log("Temporary file could not be created.")
	}
	return file
}

func createTempDir(t *testing.T) string {
	dir, err := ioutil.TempDir(".", "tempDir")
	if err != nil {
		t.Log("Temporary file could not be created.")
	}
	return dir
}
