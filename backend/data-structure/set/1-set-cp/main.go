// Tulis program sebagai fungsi yang memeriksa apakah dua set dikatakan intersection atau tidak.
// Jika kedua set intersection, fungsi tersebut harus mengembalikan nilai intersection nya.
//
// Contoh 1
// Input: a = {"Java", "Python", "Javascript", "C ++", "C#"}, b = {"Swift", "Java", "Kotlin", "Python"}
// Output: {"Java", "Python"}
// Explanation: intersection dari a dan b adalah "Java" dan "Python"
//
// Contoh 2
// Input: a = {"Java", "Python"}, b = {"Swift"}
// Output: {}
// Explanation: tidak ada intersection dari a dan b

package main

import "fmt"

func main() {
	var str1 = []string{"Java", "Python", "Javascript", "C ++", "C#"}
	var str2 = []string{"Swift", "Java", "Kotlin", "Python"}
	fmt.Println(Intersection(str1, str2))
}

func Intersection(str1, str2 []string) (inter []string) {
	//return []string{} // TODO: replace this
	hash := make(map[string]bool)
	for _, v := range str1 {
		hash[v] = true
	}
	for _, v := range str2 {
		if _, exist := hash[v]; exist {
			inter = append(inter, v)
		}
	}
	return inter
}

func RemoveDuplicates(elements []string) (nodups []string) {
	//return []string{} // TODO: replace this
	hash := make(map[string]bool)

	for _, v := range elements {
		if _, exist := hash[v]; !exist {
			hash[v] = true
			nodups = append(nodups, v)
		}
	}
	return nodups
}
