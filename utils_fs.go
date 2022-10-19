package main

import "os"

func IterateDirectoryRecurse(path string) ([]string, error) {
	var files []string
	readDirResult, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, e := range readDirResult {
		entryPath := path + "/" + e.Name()
		if e.IsDir() {
			subFiles, err := IterateDirectoryRecurse(entryPath)
			if err != nil {
				return nil, err
			}
			files = append(files, subFiles...)
		}
		if e.Type().IsRegular() {
			files = append(files, entryPath)
		}
	}
	return files, nil
}
