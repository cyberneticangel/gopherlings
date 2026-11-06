// Embedding nests one struct inside another by leaving the field
// name off:
//
//	type Car struct {
//		Make string
//		Engine // embedded: no field name
//	}
//
// The embedded type's fields are PROMOTED: a Car can read and
// write Horsepower directly, as if Engine's fields lived on Car.
// The long way, c.Engine.Horsepower, works too.
//
// Task: make Describe return "<make> with <hp> HP" -- for example
// "Zephyr with 210 HP" -- using the promoted field.
package main

import "fmt"

// Engine is a car engine.
type Engine struct {
	Horsepower int
}

// Car is a vehicle with an embedded Engine.
type Car struct {
	Make string
	Engine
}

// Describe returns a one-line summary of the car.
func (c Car) Describe() string {
	return "" // TODO: use c.Make and the promoted c.Horsepower
}

func main() {
	car := Car{Make: "Zephyr", Engine: Engine{Horsepower: 210}}
	fmt.Println(car.Describe())
}

// I AM NOT DONE
