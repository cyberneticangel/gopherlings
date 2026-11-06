// Embedding puts one struct inside another without a field name,
// and promotes its fields: a Car reads Horsepower directly, no
// c.Engine.Horsepower needed. Methods on the embedded type are
// promoted the same way.
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
	return fmt.Sprintf("%s with %d HP", c.Make, c.Horsepower)
}

func main() {
	car := Car{Make: "Zephyr", Engine: Engine{Horsepower: 210}}
	fmt.Println(car.Describe())
}
