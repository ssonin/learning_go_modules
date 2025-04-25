// Package learning_go_modules contains learning exercises from the chapter 10 of the [Learning Go] book
//
// [Learning Go]: https://www.amazon.com/Learning-Go-Jon-Bodner-ebook/dp/B0CS5DY1VN/ref=tmm_kin_swatch_0
package learning_go_modules

import (
	"golang.org/x/exp/constraints"
)

// Number is a type that permits any integer or floating-point type
type Number interface {
	constraints.Integer | constraints.Float
}

// Add [adds] two numbers
//
// [adds]: https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}
