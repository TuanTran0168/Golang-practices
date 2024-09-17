package main

import "go-oops/polymorphism"

func main() {
	// ================ Polymorphism ================
	/*
		// Syntax 1

		var circle polymorphism.Shape = &polymorphism.Circle{}
		circle.Render()

		var square polymorphism.Shape = &polymorphism.Square{}
		square.Render()

		fmt.Println(circle, square)
	*/

	/*
		// Syntax 2

		circle := &polymorphism.Circle{}
		square := &polymorphism.Square{}

		polymorphism.ShapeRender(circle)
		polymorphism.ShapeRender(square)
	*/

	// Syntax 3
	circle := new(polymorphism.Circle)
	square := new(polymorphism.Square)

	polymorphism.ShapeRender(circle)
	polymorphism.ShapeRender(square)
}
