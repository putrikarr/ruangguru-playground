// Reverse

package main

import "fmt"

func Reverse(st []string, depth int) string {
	str := ""
	for i := depth; i >= 0; i-- {
		str += st[i]
	}
	// TODO: answer here
	return str
}

func main() {
	str := []string{"A", "I", "S", "E", "N", "O", "D", "N", "I"}
	s := Reverse(str, len(str)-1)
	fmt.Println(s)
}
