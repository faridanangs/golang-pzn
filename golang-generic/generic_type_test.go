package golanggeneric

import (
	"fmt"
	"testing"
)

type Bag[T comparable] []T

func PrintBag[E comparable](params Bag[E]) {
	for _, value := range params {
		fmt.Println(value)
	}
}

func TestBagString(t *testing.T) {
	name := Bag[string]{"farid", "anang", "samudra"}
	PrintBag(name)
}
func TestBagInt(t *testing.T) {
	number := Bag[int]{1, 2, 3}
	PrintBag(number)
}
