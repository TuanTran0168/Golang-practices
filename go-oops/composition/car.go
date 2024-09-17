package composition

type Engine struct {
	name string
	hp   int
}

func (engine Engine) GetName() string {
	return engine.name
}

func (engine Engine) GetHP() int {
	return engine.hp
}

type Wheel struct {
	name            string
	wheelDimensions int
}

func (wheel Wheel) GetName() string {
	return wheel.name
}

func (wheel Wheel) GetWheelDimensions() int {
	return wheel.wheelDimensions
}

/*
composition (inheritance): embedding struct into another struct (like "extends" in Java)
*/
type Car struct {
	name   string
	engine Engine
	wheel  Wheel
}

func (car Car) GetName() string {
	return car.name
}

func (car Car) GetEngine() Engine {
	return car.engine
}

func (car Car) GetWheel() Wheel {
	return car.wheel
}

// constructor
func NewCarWithProperties(carName string, engineName string, hp int, wheelName string, wheelDimensions int) Car {
	return Car{
		name: carName,
		engine: Engine{
			name: engineName,
			hp:   hp,
		},
		wheel: Wheel{
			name:            wheelName,
			wheelDimensions: wheelDimensions,
		},
	}
}

// constructor
func NewCar(engine Engine, wheel Wheel) Car {
	return Car{
		engine: engine,
		wheel:  wheel,
	}
}
