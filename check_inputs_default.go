package main

import "os"

type CheckInputsDefault struct {
	filePaths []string
}

func (c *CheckInputsDefault) GetFilePaths() []string {
	return c.filePaths
}

func (c *CheckInputsDefault) ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}
