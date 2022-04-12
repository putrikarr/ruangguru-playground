package main

import (
	"fmt"
	"sort"
)

type Score struct {
	Name    string
	Correct int
	Wrong   int
	Empty   int
}

type Scores []Score

func (s Scores) Len() int {
	return len(s)
}

func (s Scores) Less(i, j int) bool {
	for _, score := range s {
		if score.Name == s[i].Name {
			if (score.Correct*4)-score.Wrong == (s[i].Correct*4)-s[i].Wrong {
				return (s[i].Correct*4)-s[i].Wrong > (s[j].Correct*4)-s[j].Wrong
			}
			return (score.Correct*4)-s[i].Wrong > (s[j].Correct*4)-s[j].Wrong
			//return true
		}
	}
	return false
}

// func (s Scores) Less(i, j int) bool {
// 	if (s[i].Correct*4)-s[i].Wrong > (s[j].Correct*4)-s[j].Wrong {
// 		return true
// 	}
// 	if (s[i].Correct*4)-s[i].Wrong == (s[j].Correct*4)-s[j].Wrong {
// 		if s[i].Name < s[j].Name {
// 			return true
// 		}
// 	}
// 	return false
// }

func (s Scores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Scores) TopStudents() []string {
	sort.Sort(s)
	names := []string{}
	for _, score := range s {
		names = append(names, score.Name)
	}
	return names
}

func main() {
	scores := Scores([]Score{
		{"Levi", 3, 2, 2},
		{"Agus", 3, 4, 0},
		{"Doni", 3, 0, 7},
		{"Ega", 3, 0, 7},
		{"Anton", 2, 0, 5},
	})

	for _, score := range scores {
		fmt.Printf("%s: %d\n", score.Name, (score.Correct*4)-score.Wrong)
	}

	sort.Sort(scores)
	fmt.Println(scores.TopStudents())
}
