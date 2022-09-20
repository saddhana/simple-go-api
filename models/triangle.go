package models

type Triangle struct {
	Figure
}

func (triangle Triangle) FigureArea() int {
	area := triangle.Base * triangle.Height / 2
	return area
}
