package main

import "strings"

func GetLastFileExtension(path string) string {
	tokens := strings.Split(path, ".")
	if len(tokens) <= 1 {
		return ""
	}
	startsWithDot := tokens[0] == ""
	if startsWithDot {
		return ""
	}
	return tokens[len(tokens)-1]
}
