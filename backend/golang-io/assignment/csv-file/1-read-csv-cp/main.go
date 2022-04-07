package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	fmt.Print("dummy commit")
}

func CSVToMap(data map[string]string, fileName string) (map[string]string, error) {
	//TODO : answer here
	f, err := os.Open(fileName)
	if err != nil {
		return data, err
	}

	defer f.Close()

	r := csv.NewReader(f)

	// skip baris pertama (nama kolom)

	if _, err := r.Read(); err != nil {
		return data, err
	}

	dataToMap, err := r.ReadAll()
	if err != nil {
		return data, err
	}

	for _, record := range dataToMap {
		data[record[0]] = record[1]
	}

	return data, nil
}
