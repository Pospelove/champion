package main

import "strings"

func GetLastFileExtension(path string) string {
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
