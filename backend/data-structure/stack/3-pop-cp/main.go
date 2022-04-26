package main

import "errors"

// Dari inisiasi stack dengan maksimal elemen sebanyak 10, implementasikan operasi pop.

var ErrStackUnderflow = errors.New("stack underflow")

type Stack struct {
	// TODO: answer here
	Top  int
	Data []int
	Size int
}

func NewStack(size int) Stack {
	// TODO: answer here
	return Stack{
		Top:  -1,
		Data: []int{},
		Size: size,
	}
}

func (s *Stack) Push(Elemen int) error {
	// TODO: answer here
	if s.Top == s.Size-1 {
		return ErrStackUnderflow
	}

	s.Top++
	s.Data = append(s.Data, Elemen)

	return nil
}

func (s *Stack) Pop() (int, error) {
	// TODO: answer here
	if s.Top == -1 {
		return 0, ErrStackUnderflow
	} else {
		poppedValue := s.Data[s.Top]
		s.Top--
		s.Data = s.Data[:len(s.Data)-1]
		return poppedValue, nil
	}

}
