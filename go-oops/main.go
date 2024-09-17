package main

import (
	"fmt"
	"go-oops/composition"
)

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

	/*
			// Syntax 3
		circle := new(polymorphism.Circle)
		square := new(polymorphism.Square)

		polymorphism.ShapeRender(circle)
		polymorphism.ShapeRender(square)
	*/

	// ================ Composition ================

	// Syntax 1
	car := composition.NewCarWithProperties("Car 1", "Engine 1", 100, "Wheel 1", 10)

	fmt.Println(car)
	fmt.Println(car.GetEngine(), car.GetEngine().GetName(), car.GetEngine().GetHP())
	fmt.Println(car.GetWheel(), car.GetWheel().GetName(), car.GetWheel().GetWheelDimensions())
}
