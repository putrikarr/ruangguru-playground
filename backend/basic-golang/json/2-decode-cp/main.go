package main

import (
	"encoding/json"
	"fmt"
)

// Dari contoh yang telah diberikan, cobalah untuk melakukan decode dari json menjadi objek dari struct.
// Mengambil kasus pada encode, lengkapi function DecodeToLeaderboard agar json dapat di decode menjadi objek Leaderboard

// type UserRank struct {
// 	Name  string
// 	Email string
// 	Rank  int
// }

type UserRank struct {
	// TODO: answer here
	Name  string `json:"name"`
	Rank  int    `json:"rank"`
	Email string
}

type Leaderboard struct {
	Users []*UserRank
}

func DecodeToLeaderboard(jsonData []byte) (Leaderboard, error) {
	// TODO: answer here
	byteJSONdata := []byte(`[
		{
			Name: "Roger",
			Rank: 1,
		},
		{
			Name: "Tony",
			Rank: 2,
		},
		{
			Name: "Bruce",
			Rank: 3,
		},
		{
			Name: "Natasha",
			Rank: 4,
		},
		{
			Name: "Clint",
			Rank: 5,
		},
}]`)

	u := UserRank{}
	err := json.Unmarshal(byteJSONdata, &u)
	if err != nil {
		println(err)
	}
	fmt.Println(u)
	return Leaderboard{}, nil

}
