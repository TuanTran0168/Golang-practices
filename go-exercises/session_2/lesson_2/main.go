package main

import (
	"fmt"
	"go-exercises/session_2/lesson_2/structs"
	"math"
)

func main() {

	// Structs
	point_A := structs.NewPoint(3, 5)
	point_B := structs.NewPoint(1, 2)

	line_A := structs.NewLine(point_A, point_B)
	fmt.Println(point_A.Distance(point_B))

	fmt.Println(line_A.ToString())

	// like static function in java
	lenght_line_A := structs.LenghtLine(line_A)
	fmt.Println("Lenght Line A (like static function):", roundFloat(lenght_line_A, 2))

	// Get Lenght in Go
	fmt.Println("Get Lenght in Go:", roundFloat(line_A.GetLenght(), 2))

	// like static function in java
	middle_point_A := structs.MiddlePoint(line_A)
	fmt.Println("Middle Point A (like static function):", middle_point_A.ToString())

	/*
		Get Middle Point in Go
		USE THIS IF GetMiddlePoint() not return pointer
		middle_point_A_1 := line_A.GetMiddlePoint()
		fmt.Println("Get Middle Point in Go:", middle_point_A_1.ToString())
	*/

	// Get Middle Point in Go
	fmt.Println("Get Middle Point in Go:", line_A.GetMiddlePoint().ToString())

	lineA := structs.NewLine(structs.NewPoint(1, 2), structs.NewPoint(1, 3))
	lineB := structs.NewLine(structs.NewPoint(2, 2), structs.NewPoint(2, 1))

	isParallel := structs.IsParallel(lineA, lineB)
	fmt.Println("Is Parallel:", isParallel)

}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
