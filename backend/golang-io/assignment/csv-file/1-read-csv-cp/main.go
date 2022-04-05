package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("dummy commit")
}

func CSVToMap(data map[string]string, fileName string) (map[string]string, error) {
	csvFile, err := os.Open("questions.csv")
	if err != nil {
		log.Fatalf("cannot open file %q: %s\n", "questions.csv", err)
	}

	defer csvFile.Close()
	csvReader := csv.NewReader(csvFile)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalln("error reading csv file", err)
	}

	for _, record := range records {
		data[record[0]] = record[1]
	}

	return data, nil
	// TODO: answer here

}
