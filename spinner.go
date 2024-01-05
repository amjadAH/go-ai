package main

type Spinner struct {
	chars   []rune
	current int
}

func newSpinner() *Spinner {
	return &Spinner{
		chars: []rune{'|', '/', '-', '\\'},
	}
}

func (s *Spinner) spin() string {
	if s.current == len(s.chars)-1 {
		s.current = 0
	} else {
		s.current++
	}

	return string(s.chars[s.current])
}
