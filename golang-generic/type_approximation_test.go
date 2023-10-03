package golanggeneric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Age int

// 2. setelah di beri tanda ~ maka dia akan mendeteksi jika type nya sama maka dia akan di bolehkan
type NumberAppro interface {
	~int | int32 | int64 | float32 | float64 |
		int8 | int16
}

func Appro[T NumberAppro](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestAppro(t *testing.T) {
	assert.Equal(t, 2, Appro[int](10, 2))
	assert.Equal(t, float32(4.4), Appro[float32](float32(4.4), float32(10.032)))
	assert.Equal(t, float64(4.4), Appro[float64](float64(4.4), float64(10.032)))
	assert.Equal(t, Age(4), Appro[Age](Age(4), Age(19)))
}

func TestInference(t *testing.T) {
	// di sini kita tidak perlu memanggil tanda [any] karna golang sudah secara otomatis membacanya
	// dan jika terjadi error itu terjadi mungkin karna go tidak bisa membacanya dan terpaksa kita harus memberinya type
	assert.Equal(t, 2, Appro(10, 2))
	assert.Equal(t, float32(4.4), Appro(float32(4.4), float32(10.032)))
	assert.Equal(t, float64(4.4), Appro(float64(4.4), float64(10.032)))
	assert.Equal(t, Age(4), Appro(Age(4), Age(19)))
}
