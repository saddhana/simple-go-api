package models

type Square struct {
	Figure
}

func (square Square) FigureArea() int {
	area := square.Side * square.Side
	return area
}
