package main

import "fmt"

type Check interface {
	InitSettings(map[interface{}]interface{})
	Run(inputs CheckInputs) error
}

func New(uses string) (Check, error) {
	switch uses {
	case "VerifyFileNames":
		return &VerifyFileNames{}, nil
	default:
		return nil, fmt.Errorf("unknown check type: %s", uses)
	}
}
