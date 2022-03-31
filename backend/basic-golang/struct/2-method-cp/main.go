package main

import "fmt"

// buat struct Rectangle (persegi panjang) dengan dua atribut yaitu Width dan Length
// tambah dua method :
// GetArea() dan GetPerimeter()
// GetArea() digunakan untuk menampilkan (print) luas persegi panjang
// GetPerimeter() digunakan untuk menampilkan (print) keliling persegi panjang

type Rectangle struct {
	// TODO: answer here
	Width  int
	Length int
}

// TODO: answer here
func (r Rectangle) GetArea() {
	area := r.Width * r.Length
	fmt.Printf("Area : %d\n", area)
}

func (r Rectangle) GetPerimeter() {
	perimeter := 2 * (r.Width + r.Length)
	fmt.Printf("Perimeter : %d\n", perimeter)
}

func main() {
	var r Rectangle
	r.Width = 10
	r.Length = 20
	fmt.Println(r)

	r.GetArea()
	r.GetPerimeter()
}
