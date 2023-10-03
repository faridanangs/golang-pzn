package golanggeneric

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

// kita menggunakan constraints untuk mencari tipe dari genericnya di dalam package exp
func Experigeneric[T constraints.Ordered](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestExperiGeneric(t *testing.T) {
	assert.Equal(t, 10, Experigeneric(10, 20))
}

func TestExperiSlice(t *testing.T) {
	first := []string{"farid"}
	last := []string{"farid"}
	assert.True(t, slices.Equal(first, last))
}
func TestExperiMap(t *testing.T) {
	first := map[string]string{"Name": "farid"}
	last := map[string]string{"Name": "farid"}
	assert.True(t, maps.Equal(first, last))
}
