package main

import "strings"

func GetLastFileExtension(path string) string {
	//tokens := strings.Split(path, ".")
	// If there is no extension, return empty string
	// In case of a file like ".gitignore", return empty string
	//if len(tokens) == 0 {
	//		return ""
	//	}
	//	return tokens[len(tokens)-1]

	tokens := strings.Split(path, ".")
	if len(tokens) <= 1 {
		return ""
	}
	// if starts with ".", return empty string
	if tokens[0] == "" {
		return ""
	}
	return tokens[len(tokens)-1]
}
