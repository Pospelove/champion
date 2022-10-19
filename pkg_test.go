package main

import "testing"

func TestUtilsPath(t *testing.T) {
	ext := GetLastFileExtension("test")
	if ext != "" {
		t.Errorf("Expected empty string, got %s", ext)
	}
	ext = GetLastFileExtension(".gitignore")
	if ext != "" {
		t.Errorf("Expected empty string, got %s", ext)
	}
	ext = GetLastFileExtension("test.txt")
	if ext != "txt" {
		t.Errorf("Expected txt, got %s", ext)
	}
	ext = GetLastFileExtension("test.txt.md")
	if ext != "md" {
		t.Errorf("Expected md, got %s", ext)
	}
}
