package main

import (
	"customerimporter/customerimporter"
	"log"
)

func main() {

	filePath := "customers.csv"
	err := customerimporter.CountEmailDomains(filePath)

	if err != nil {

		log.Fatal(err)
	}
}
