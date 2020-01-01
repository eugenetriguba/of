package fs

import (
	"fmt"
	"os"
)

// Close the file and handle the error if it occurs
// by printing it out and exiting with a code of 1.
//
// Example:
//     file, err := os.Open(filePath)
//     ... handle error
//     defer fs.CloseFile(file)
func CloseFile(file *os.File) {
	err := file.Close()

	if err != nil {
		fmt.Println("Error occurred while trying to close the file: ", err)
		os.Exit(1)
	}
}

// Checks if a file exists and is not a directory.
//
// The given filePath should be an absolute path.
// Returns true if the file exists; false otherwise.
func FileExists(filePath string) bool {
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// Checks if a directory exists.
//
// The given dirPath should be an absolute path.
// Returns true if the directory exists; false otherwise.
func DirExists(dirPath string) bool {
	info, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}
