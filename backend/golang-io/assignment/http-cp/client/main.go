package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Pokemon struct {
	Name    string `json:"name"`
	Species struct {
		Name string `json:"name"`
	} `json:"species"`
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
		}
	}
}

func GetPokemonData() (*Pokemon, error) {
	apiPath := "https://pokeapi.co/api/v2/pokemon/1"
	fmt.Println(apiPath)

	// TODO: answer here
	client := &http.Client{}
	res, err := client.Get(apiPath)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	pokemon := Pokemon{}
	json.NewDecoder(res.Body).Decode(&pokemon)

	return &pokemon, nil
}

func main() {
	pokemon, err := GetPokemonData()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", pokemon)
}
