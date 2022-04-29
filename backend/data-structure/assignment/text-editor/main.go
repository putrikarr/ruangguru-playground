package main

import (
	"fmt"

	"github.com/ruang-guru/playground/backend/data-structure/assignment/text-editor/stack"
)

// Implementasikan teks editor sederhana
// Teks editor ini memiliki 4 method
// Write: digunakan untuk menulis per karakter
// Read: digunakan untuk mencetak karakter yang telah ditulis
// Undo: digunakan untuk melakukan operasi undo
// Redo: digunakan untuk melakukan operasi redo

type TextEditor struct {
	UndoStack *stack.Stack
	RedoStack *stack.Stack
}

func NewTextEditor() *TextEditor {
	return &TextEditor{
		UndoStack: stack.NewStack(),
		RedoStack: stack.NewStack(),
	}
}

func (te *TextEditor) Write(ch rune) {
	if te.UndoStack.IsEmpty() {
		te.UndoStack.Push(ch)
		fmt.Println("empty")
	} else {
		lastChar, _ := te.UndoStack.Peek()
		if lastChar != ch {
			te.UndoStack.Push(ch)
			fmt.Println("Write:", string(te.UndoStack.Data))
		}
	}
	// TODO: answer here
}

func (te *TextEditor) Read() []rune {
	if te.UndoStack.IsEmpty() {
		fmt.Println("empty")
		return nil
	} else {
		ch, _ := te.UndoStack.Pop()
		te.UndoStack.Push(ch)
		fmt.Println("Read:", string(te.UndoStack.Data))
		return te.UndoStack.Data
	}
	// TODO: answer here
}

func (te *TextEditor) Undo() {
	if te.UndoStack.IsEmpty() {
		return
	} else {
		ch, _ := te.UndoStack.Pop()
		te.RedoStack.Push(ch)
	}
	// TODO: answer here
}

func (te *TextEditor) Redo() {
	if te.RedoStack.IsEmpty() {
		return
	} else {
		ch, _ := te.RedoStack.Pop()
		te.UndoStack.Push(ch)
	}
	// TODO: answer here
}
