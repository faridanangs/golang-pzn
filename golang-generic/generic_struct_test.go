package golanggeneric

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type DataName[T any] struct {
	Name, Gender T
}

// kita bisa membuat method di dalam strcut generict ini dengan catatan method tidak bisa menerima generic

func (D *DataName[T]) ChangeName(name T) T {
	D.Name = name
	return D.Name
}

// jika kita tidak menggunakan generic nya kita bisa menggunakan _
func (D *DataName[_]) SayHello(name string) string {
	return "hello " + name
}

func TestDataStruct(t *testing.T) {
	data := DataName[string]{
		Name:   "farid",
		Gender: "pria",
	}
	fmt.Println(data)
}

func TestDataStructMethod(t *testing.T) {
	data := DataName[string]{
		Name:   "farid",
		Gender: "pria",
	}

	assert.Equal(t, "hello farid", data.SayHello("farid"))
	assert.Equal(t, "anang", data.ChangeName("anang"))
}
