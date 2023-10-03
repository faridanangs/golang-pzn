package golanggeneric

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// untuk membuat type generict kita gunakan tanda [] di funcnya dan kita beri tipe constraint
// contoh [T any] jika kita tidak memberi constraint dia akam error
// dan hanya bisa di gunakan di local func nya saja
func lengthAny[T any](params T) T {
	fmt.Println(params)
	return params
}

// // // type Any
// jika kita mnggunakan type any maka kita tidak bisa melakukan perbandingan logika di dalamnya
// func LengthPerbandingan[T any](params T) T {
// 	var angka int
// 	if reflect.TypeOf(params) == reflect.TypeOf(angka) {
// 		fmt.Println(strconv.Itoa(params.(int)))
// 		return params
// 	}
// 	return params
// }

func TestSampleAny(t *testing.T) {
	// kita mengitim nilai ke dalam func Length dan memberi tipe dari genericnya dan memasukan nilai ke dalam func nya
	// contoh penggunaannya var name = LengthAny[type](value)
	var name = lengthAny[string]("Farid anang s")
	assert.Equal(t, "Farid anang s", name)
	var kelas = lengthAny[int](12)
	assert.Equal(t, 12, kelas)
}

//////// untuk melakuka perbandingan kita bisa menggunakan type comparable

// // type comparable
// jika kita mnggunakan type comparable maka kita bisa melakukan perbandingan logika di dalamnya
func lengthComparable[T comparable](params1, params2 T) bool {
	if params1 == params2 {
		return true
	} else {
		return false
	}
}

func TestSampleComparable(t *testing.T) {
	// kita mengirim nilai ke dalam func lengthComparable dan memberi tipe dari genericnya dan memasukan nilai ke dalam func nya
	// contoh penggunaannya var name = lengthComparable[type](value1, value2)
	var name = lengthComparable[string]("Farid anang s", "wagas")
	assert.Equal(t, false, name)
	var kelas = lengthComparable[int](12, 12)
	assert.Equal(t, true, kelas)
}
