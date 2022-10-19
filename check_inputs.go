package main

type CheckInputs interface {
	// Get list of file paths to check
	GetFilePaths() []string
	// Read file
	ReadFile(path string) ([]byte, error)
}
