package golanggeneric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// untuk membuat type Set kita bisa menggunakan tanda | jika type datanya
// lebih dari satu
type Number interface {
	int | int32 | int64 | float32 | float64 |
		int8 | int16 // | SiswaInter jika di masukan ini akan error
}

func Min[T Number](first, second T) T {
	// kita hanya bisa menggunakan semua oprasi yang ada di dalm type data setnya
	if first < second {
		return first
	} else {
		return second
	}
}

func TestMin(t *testing.T) {
	assert.Equal(t, 2, Min[int](10, 2))
	assert.Equal(t, float32(4.4), Min[float32](float32(4.4), float32(10.032)))
	assert.Equal(t, float64(4.4), Min[float64](float64(4.4), float64(10.032)))

	// jika kita menggunakan string namun di dalam type set number tidak ada string maka
	// ini akan mengalami error
	// assert.Equal(t, 10, Min[string]("anangs", "farid"))

}
