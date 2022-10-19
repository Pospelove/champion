package main

import (
	"os"

	"github.com/go-yaml/yaml"
)

type Settings struct {
	rawContents  string
	loadedChecks []Check
}

func (s *Settings) Load(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	s.rawContents = string(data)

	var parsed interface{}
	err = yaml.Unmarshal(data, &parsed)
	if err != nil {
		return err
	}

	// get "checks" array from parsed
	checks := parsed.(map[interface{}]interface{})["checks"].([]interface{})

	// for each check in checks
	for _, check := range checks {
		uses := check.(map[interface{}]interface{})["uses"].(string)
		with := check.(map[interface{}]interface{})["with"].(map[interface{}]interface{})

		verify, err := New(uses)
		if err != nil {
			return err
		}

		verify.InitSettings(with)
		s.loadedChecks = append(s.loadedChecks, verify)
	}

	return nil
}
