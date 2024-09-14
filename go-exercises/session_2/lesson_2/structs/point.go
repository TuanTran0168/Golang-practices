package structs

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

// func NewPoint(x float64, y float64) *Point {
// 	return &Point{x, y}
// }

// Constructor in Golang
func NewPoint(x float64, y float64) Point {
	return Point{x, y}
}

func (point *Point) GetX() float64 {
	return point.x
}

func (point *Point) GetY() float64 {
	return point.y
}

func (point *Point) SetX(x float64) {
	point.x = x
}

func (point *Point) SetY(y float64) {
	point.y = y
}

func (point *Point) ToString() string {
	return fmt.Sprintf("(%v, %v)", point.x, point.y)
}

func (point_A *Point) Distance(point_B Point) float64 {
	x_value := math.Pow(point_A.x-point_B.x, 2)
	y_value := math.Pow(point_A.y-point_B.y, 2)
	distance := math.Sqrt(x_value + y_value)
	return distance
}
