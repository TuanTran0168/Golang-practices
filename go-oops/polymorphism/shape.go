package polymorphism

import "fmt"

type Shape interface {
	Render()
}

func ShapeRender(shape Shape) {
	/*
		When you call shape.Render(),
		Go dynamically determines and calls the Render method of the actual object passed,
		which is based on its concrete type. This is the essence of polymorphism.
		You can pass in a Circle, Square, or any type that implements the Render method, and Go will handle it appropriately.

		Explanation:
			- 	ShapeRender(circle): You pass a *Circle object into the ShapeRender function.
				The function will call the Render() method of Circle and print "Rendering Circle".

			- 	ShapeRender(square): Similarly, you pass a *Square object into the function,
				and it will print "Rendering Square".
	*/
	shape.Render()
}

type Circle struct {
}

func (circle *Circle) Render() {
	fmt.Println("Rendering Circle")
}

type Square struct {
}

func (square *Square) Render() {
	fmt.Println("Rendering Square")
}
