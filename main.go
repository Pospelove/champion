package main

import (
	"flag"
	"log"
)

func main() {
	settingsPath := flag.String("settings-path", "./champion.yml", "Path to settings file")

	flag.Parse()

	settings := Settings{}

	err := settings.Load(*settingsPath)
	if err != nil {
		log.Fatal(err)
	}

	checkInputs := CheckInputsDefault{}
	filePaths, err := IterateDirectoryRecurse(".")
	if err != nil {
		log.Fatal(err)
	}
	checkInputs.filePaths = filePaths

	for _, check := range settings.loadedChecks {
		err := check.Run(&checkInputs)
		if err != nil {
			log.Fatal(err)
		}
	}
}
