// Package fs contains utilities to check if files or directories
// exist and closing files.
package fs

import (
	"fmt"
	"os"
)

// CloseFile closes the file and handles the error if it
// occurs by printing it out and exiting with a code of 1.
//
// Example:
//  file, err := os.Open(filePath)
//  ... handle error
//  defer fs.CloseFile(file)
func CloseFile(file *os.File) {
	err := file.Close()

	if err != nil {
		fmt.Println("Error occurred while trying to close the file: ", err)
		os.Exit(1)
	}
}

// FileExists checks if a file exists and is not a directory.
//
// The given filePath should be an absolute path.
// Returns true if the file exists; false otherwise.
//
// Example:
//  exists, err := fs.FileExists(filePath)
//	if err != nil {
//    ... handle error
//  }
//
//  if exists {
//    ... do work
//  }
func FileExists(filePath string) (bool, error) {
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return !info.IsDir(), nil
}

// DirExists checks if a directory exists.
//
// The given dirPath should be an absolute path.
// Returns true if the directory exists; false otherwise.
//
// Example:
//  exists, err := fs.DirExists(dirPath)
//	if err != nil {
//    ... handle error
//  }
//
//  if exists {
//    ... do work
//  }
func DirExists(dirPath string) (bool, error) {
	info, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return info.IsDir(), nil
}
