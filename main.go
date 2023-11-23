package main

import (
	"customerimporter/customerimporter"
	"log"
	"os"
)

func main() {

	filePath := "customers.csv"

	file, err := customerimporter.OpenFile(filePath)

	if err != nil {

		log.Fatal(err)
	}

	defer func(file *os.File) {

		err := file.Close()

		if err != nil {

		}
	}(file)

	_, err = customerimporter.CountEmailDomains(file)

	if err != nil {

		log.Fatal(err)
	}
}
