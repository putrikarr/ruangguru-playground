package main

import "fmt"

func main() {
	/*
		Score Rank:
		90-100: A
		80-89: B
		70-79: C
		60-69: D
		0-59: E
	*/
	res := ScoreRank(50)
	fmt.Println(res)

	// Try correct answer:
	// resCorrect := ScoreRankCorrect(arr)
	// fmt.Println(resCorrect)
}

func ScoreRank(score int) string {
	if score > 100 || score < 0 {
		return "INVALID"
	}
	var rank string
	switch {
	case score > 60:
		rank = "D"
	case score > 70:
		rank = "C"
	case score > 80:
		rank = "B"
	case score > 90:
		rank = "A"
	default:
		rank = "E"
	}

	return rank
}

func ScoreRankCorrect(score int) string {
	if ScoreRank(score) == "INVALID" {
		return "INVALID"
	} else if ScoreRank(score) == "E" {
		return "E"
	} else if ScoreRank(score) == "A" {
		return "A"
	} else if ScoreRank(score) == "B" {
		return "B"
	} else if ScoreRank(score) == "C" {
		return "C"
	} else if ScoreRank(score) == "D" {
		return "D"
	}
	return "INVALID"
	//return "" // TODO: replace this
}
