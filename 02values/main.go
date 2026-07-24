package values

import (
	"fmt"
	"math"
)

func swap(x, y string) (string, string) {
	return y, x
}
func main() {
	fmt.Println("Hello, World!")
	fmt.Println(math.Pi)

	a, b := swap("Hello", "World")
	fmt.Println(a, b)
}
