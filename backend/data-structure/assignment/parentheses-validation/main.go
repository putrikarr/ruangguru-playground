package main

import (
	"errors"
	"fmt"
)

// Dari inisiasi stack dengan jumlah maksimal elemen 10, cobalah implementasikan operasi push.

var ErrStackOverflow = errors.New("stack overflow")
var ErrStackUnderflow = errors.New("stack underflow")
var ErrEmptyStack = errors.New("stack is empty")

type Stack struct {
	// TODO: answer here
	Top  int
	Size int
	Data []int
}

func NewStack(size int) Stack {
	// TODO: answer here
	return Stack{
		Top:  -1,
		Size: size,
		Data: []int{},
	}
}

func (s *Stack) Push(Elemen int) error {
	// TODO: answer here
	if len(s.Data) >= s.Size { // 10 , 10
		return ErrStackOverflow
	}

	s.Top += 1
	s.Data[s.Top] = Elemen
	return nil
}

// type Stack struct {
// 	Top  int
// 	Data []int
// }

func (s *Stack) IsEmpty() bool {
	return s.Top == -1
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, ErrStackUnderflow
	}

	poppedValue := s.Data[s.Top]    // simpan data top di variable proppedValue
	s.Top -= 1                      // {1,2,3,4} => top = 3 - 1 = 2 => {1,2,3} =
	s.Data = s.Data[:len(s.Data)-1] // bakalan slicing data dengan menghapus elemen terakhir
	return poppedValue, nil         // return data pop, error nil
}

func (s *Stack) Peek() (int, error) {
	// TODO: answer here
	if s.IsEmpty() { // cek kondisi bila data itu kosong/empty
		return 0, ErrEmptyStack // jika data kosong, return 0, dan error empty stack
	}

	// jika data tidak kosong, maka return data pada index ke Top, dan error nil
	return s.Data[s.Top], nil
}

// Salah satu problem populer yang dapat diselesaikan dengan menggunakan Stack adalah mengecek validitas tanda kurung.
// Diberikan sebuah string yang hanya terdapat karakter '(', ')', '{', '}', '[', dan ']'.
// Tentukan apakah sebuah string merupakan sekumpulan tanda kurung yang valid.
// String tanda kurung yang valid adalah
// 1. Tanda kurung buka harus ditutup oleh pasangannya.
// 2. Tanda kurung buka harus ditutup di urutan yang benar.

// Contoh 1
// Input: s = "()"
// Output: true
// Penjelasan: tanda kurung buka '(' ditutup dengan pasangannya yaitu ')'.

// Contoh 2
// Input: s = "{)"
// Output: false
// Penjelasan: tanda kurung buka '{' tidak ditutup oleh pasangannya.

// Contoh 3
// Input: s = "([])" // true
// Output: true
// Penjelasan: tanda kurung buka ditutup dengan pasangannya dan sesuai dengan urutan.

func IsValidParentheses(s string) bool {
	// if item == "(" {

	// } else if item == "{" {

	// } else if item == "[" {

	// }

	// mapKurung := map[int]int{}
	// val, exist := mapKurung[item]
	// if exist {
	// } else {
	// }

	// {()[()]} => map [")"] = "(", map["]"] = "["
	// { push, ( push, ) pop dan compare dengan mapping
	/*
	   data { => push to stack
	       stack = " {"
	   data ( => push to stack
	     stack = "{, ("
	   data ) => pop dan check apakaha match atau tidak dengan data Top
	     pop => return ( => compare dg map[data] => map[")"]  => (
	     stack = "{"
	   data [ => push
	     stack = "{, ["
	   data ( => push
	     stack = "{, [, ("
	   data ) => pop
	     return pop = ( => check pop.Data == map[data] == continue
	     stack = "{, ["
	   data ] => pop
	     return pop => [ => check pop.Data == map[data] == contu
	     stack "{"
	   data } => pop
	     return pop => { => check pop data = map[data]
	     stack = ""

	   return true
	*/
	//
	// ex: s = ([]) => check apakah semua kurung mempunyai pasangan dan pasangannya
	// sesuai dengan urutan
	// TODO: answer here
	// bracket map[string]string => bracket["]"] = "[", semua kurung di define
	// ketika buka kurung, simpan char ke dalam stack.
	// ketika char itu tutup kurung, kita panggil map, apakah yg data.Top = val map bracket
	return false
}

// type TextEditor struct {
// 	UndoStack *stack.Stack
// 	RedoStack *stack.Stack
// }

// func NewTextEditor() *TextEditor {
// 	return &TextEditor{
// 		UndoStack: stack.NewStack(),
// 		RedoStack: stack.NewStack(),
// 	}
// }

// func (te *TextEditor) Write(ch rune) {
// TODO: answer here
/*
	ide:
	  ch akan di push ke UndoStack
*/
// }

// func (te *TextEditor) Read() []rune {
// TODO: answer here
/*
  return semua data ke dalam []rune
  dengan urutan sesuai dengan waktu write
  write(h)
  write(e)
  write(l)
  UndoStack
  Read() => []Rune{h,e,l}
*/
// }

// func (te *TextEditor) Undo() {
// TODO: answer here
/*
  pop UndoStack
 dataPop = te pop()
 response dari undo stack di Push ke dalam RedoStack
 RedoStack.Push(dataPop)
*/
// }

// func (te *TextEditor) Redo() {
// TODO: answer here
/*
	pop RedoStack
	 dataPop = te pop()
	 response dari pop Redo stack di Push ke dalam UndoStack
	 UndoStack.Push(dataPop)
*/
// }

func main() {
	fmt.Printf("Go Hello, World!")
}
