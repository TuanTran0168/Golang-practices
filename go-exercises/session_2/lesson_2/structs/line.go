package structs

import (
	"fmt"
)

type Line struct {
	startPoint Point
	endPoint   Point
}

func NewLine(point_A, point_B Point) Line {
	return Line{point_A, point_B}
}

func (line *Line) GetStartPoint() Point {
	return line.startPoint
}

func (line *Line) GetEndPoint() Point {
	return line.endPoint
}

func (line *Line) SetStartPoint(point Point) {
	line.startPoint = point
}

func (line *Line) SetEndPoint(point Point) {
	line.endPoint = point
}

func (line *Line) ToString() string {
	return fmt.Sprintf("(%v, %v), (%v, %v)", line.startPoint.x, line.startPoint.y, line.endPoint.x, line.endPoint.y)
}

// like static function in java
func LenghtLine(line Line) float64 {
	return line.startPoint.Distance(line.endPoint)
}

// Get Lenght in Go
func (line *Line) GetLenght() float64 {
	return line.startPoint.Distance(line.endPoint)
}

// like static function in java
func MiddlePoint(line Line) Point {
	x_value := (line.startPoint.x + line.endPoint.x) / 2
	y_value := (line.startPoint.y + line.endPoint.y) / 2

	return Point{x_value, y_value}
}

// Get Middle Point in Go
func (line *Line) GetMiddlePoint() *Point {
	x_value := (line.startPoint.x + line.endPoint.x) / 2
	y_value := (line.startPoint.y + line.endPoint.y) / 2

	return &Point{x_value, y_value}
}

func IsParallel(line_A, line_B Line) bool {
	checkLineA := (line_A.startPoint.x - line_A.startPoint.x) / (line_A.endPoint.x - line_A.endPoint.y)
	checkLineB := (line_B.startPoint.x - line_B.startPoint.x) / (line_B.endPoint.x - line_B.endPoint.y)

	return checkLineA == checkLineB
}
