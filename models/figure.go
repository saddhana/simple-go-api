package models

import (
	"gorm.io/gorm"
)

type Figure struct {
	gorm.Model
	Area          int
	Circumference int
	Side          int
	Base          int
	Height        int
}

func (figure Figure) FigureArea() int {
	return figure.Area
}
