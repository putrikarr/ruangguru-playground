package main

import (
	// import
	"errors"
	"fmt"
)

// Dari contoh yang diberikan, buatlah sentinel error tambahan untuk error handling pada function FindData
// Kalian dapat menambahkan sentinel error untuk mengecek apakah umurnya valid (misal kurang dari 0) dengan nama ErrInvalidAge

var ErrDataNotFound = errors.New("error data not found")

// TODO: answer here
var ErrInvalidAge = errors.New("error age is invalid")

func GetAge(data map[string]int, name string) (int, error) {
	if _, ok := data[name]; !ok {
		return 0, ErrDataNotFound
	}

	if data[name] < 0 {
		// Isilah baris ini dengan return 0 dan sentinel error ErrInvalidAge
		// TODO: answer here
		return 0, fmt.Errorf("%w, less than 0", ErrInvalidAge)
	}

	return data[name], nil
}

func main() {
	peopleAge := map[string]int{
		"John": 20,
		"Raz":  8,
		"Tony": -1,
	}

	_, err := GetAge(peopleAge, "John")
	if err != nil {
		fmt.Println(err)
	}
}
