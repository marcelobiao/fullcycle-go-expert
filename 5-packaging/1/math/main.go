package myMath

type Math struct {
	A, B int
}

func (m Math) Add() int {
	return m.A + m.B
}
