package golanggeneric

import (
	"fmt"
	"testing"
)

func MultipleParameter[T1 any, T2 any](params1 T1, params2 T2) {
	fmt.Println(params1)
	fmt.Println(params2)
}

func TestMultipleParameter(t *testing.T) {
	MultipleParameter[string, int]("Farid anang s", 12)
}
