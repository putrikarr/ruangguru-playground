package main

import (
	"encoding/csv"
	//"log"
	"os"
	"strconv"
)

type Menu struct {
	Name  string
	Price int
}

func WriteToCSV(fileName string, records []Menu) error {
	// TODO: answer here
	csvFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush()

	for _, record := range records {
		row := []string{record.Name, strconv.Itoa(record.Price)}
		if err := csvWriter.Write(row); err != nil {
			return err
		}
	}
	return nil
}
