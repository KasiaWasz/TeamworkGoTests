package main

import (
	"customerimporter/customerimporter"
	"log"
	"os"
)

func main() {

	filePath := "data/customers.csv"

	file, err := customerimporter.OpenFile(filePath)

	if err != nil {

		log.Fatal(err)
	}

	defer func(file *os.File) {

		err := file.Close()

		if err != nil {

		}
	}(file)

	err = customerimporter.CountEmailDomains(file)

	if err != nil {

		log.Fatal(err)
	}
}
