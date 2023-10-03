package golanggeneric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// inline generict kita langsung tulis di dalam genericnya
func InlineGeneric[T interface{ int | float32 | int32 }](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestInlineGener(t *testing.T) {
	assert.Equal(t, 10, InlineGeneric(10, 20))
	assert.Equal(t, int32(10), InlineGeneric(int32(10), int32(223)))
	assert.Equal(t, float32(10.102), InlineGeneric(float32(10.102), float32(101.4)))
}

// inline generic dan slice

// jadi di sini kita membuat sebuah func InlineGenericAndSlice[]()
// func InlineGenericAndSlice[T []E, E any](value T) E {}
// E := Bertipe Any
// kemudain kita memberi tipe E ke dalam slice T, kemudain kita menerima data T yang bertipe []
// dan kita mengembaikan data E/any/interface
func InlineGenericAndSlice[T []E, E any](value T) E {
	first := value[0]
	return first
}

func TestInlineGeneriAndSlice(t *testing.T) {
	data := []string{"farid", "anang", "samudra"}
	assert.Equal(t, "farid", InlineGenericAndSlice(data))
}
