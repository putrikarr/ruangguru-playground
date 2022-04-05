package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

type Menu struct {
	Name  string
	Price int
}

func WriteToCSV(fileName string, records []Menu) error {
	// TODO: answer here
	csvFile, err := os.Create("menu.csv")
	if err != nil {
		log.Fatalf("cannot create file %q: %s\n", "menu.csv", err)
	}

	defer csvFile.Close()
	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush()

	records = []Menu{
		{"Pizza", 100},
		{"Burger", 200},
		{"Coffee", 300},
		{"Tea", 400},
		{"Sandwich", 500},
	}

	for _, record := range records {
		row := []string{record.Name, strconv.Itoa(record.Price)}
		if err := csvWriter.Write(row); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
	return nil
}
