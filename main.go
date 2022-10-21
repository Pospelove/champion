package main

import (
	"flag"
	"log"
)

func main() {
	settingsPath := flag.String("settings-path", "./champion.yml", "Path to settings file")

	flag.Parse()

	settings := Settings{}

	log.Println("Loading settings from", *settingsPath)
	err := settings.Load(*settingsPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Settings loaded successfully")
	
	checkInputs := CheckInputsDefault{}
	filePaths, err := IterateDirectoryRecurse(".")
	if err != nil {
		log.Fatal(err)
	}
	checkInputs.filePaths = filePaths

	log.Println("Running checks")
	for _, check := range settings.loadedChecks {
		log.Println("Running check", check.GetName())
		err := check.Run(&checkInputs)
		if err != nil {
			log.Fatal(err)
		}
	}
}
