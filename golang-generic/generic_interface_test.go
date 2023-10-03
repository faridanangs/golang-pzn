package golanggeneric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type GetterSetter[T any] interface {
	GetValue() T
	SetValue(value T)
}

// kita menjalankan interface method jika ChangeValue di jalankan dan dia akan secara
// otomatis akan mengirim nilai yang di terimanya ke dalam method struct jika dia memenuhi syarat yaitu
// memiliki kesamaan dengan method interfce jika param.SetValue(value) ini menerima nilai maka dia akan melemparnya
// ke dalam method stuct yang sama dengannya
func ChangeValue[T any](param GetterSetter[T], value T) T {
	param.SetValue(value)
	return param.GetValue()
}

type MyDataKu[T any] struct {
	Value T
}

func (data *MyDataKu[T]) SetValue(value T) {
	data.Value = value
}
func (data *MyDataKu[T]) GetValue() T {
	return data.Value
}

func TestInterfaceGenerict(t *testing.T) {
	assert.Equal(t, "farid", ChangeValue(&MyDataKu[string]{}, "farid"))
}
