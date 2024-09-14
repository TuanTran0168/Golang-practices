package main

import (
	"fmt"
	"go-exercises/session_2/lesson_1/structs"
	"math"
)

func main() {
	fmt.Println("Hello World")
	p1 := structs.Point{}
	p1.SetX(1)
	p1.SetY(2)
	p2 := structs.NewPoint(3, 4)

	fmt.Println(p1.ToString())
	fmt.Println(p2.ToString())

	distance := p1.Distance(p2)

	fmt.Println(roundFloat(distance, 2))

}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

// func DistanceBetweenPoints(point_A, point_B structs.Point) float64 {
// 	x_value := math.Pow(point_A.GetX()-point_B.GetX(), 2)
// 	y_value := math.Pow(point_A.GetY()-point_B.GetY(), 2)
// 	distance := math.Sqrt(x_value + y_value)

// 	return distance
// }
