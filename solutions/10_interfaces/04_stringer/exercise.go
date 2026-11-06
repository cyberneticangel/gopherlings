// Implement one method, String() string, and your type satisfies
// fmt.Stringer: every %v-style print calls it. fmt discovers the
// method at runtime -- the same implicit satisfaction you met with
// Speaker, this time plugging your type into fmt's machinery.
package main

import "fmt"

// Temperature is a temperature in degrees Celsius.
type Temperature float64

// String formats the temperature with one decimal and a degree sign.
func (t Temperature) String() string {
	return fmt.Sprintf("%.1f°C", float64(t))
}

func main() {
	fmt.Printf("room temperature: %v\n", Temperature(21.5))
}
