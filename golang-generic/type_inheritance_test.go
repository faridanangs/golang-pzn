package golanggeneric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Data interface {
	GetName() string
}

func GolangGeneericInheren[T Data](params T) string {
	return params.GetName()
}

//* jika sebuah struct memiliki method yang sama dengan Data interface itu maka
// secara otomatis kita sedang melakukan kontrak atau inheritance

// ini juga masuk ke dalam data karna dia memiliki method GetName() yang sama dengan interface Data
type SiswaInter interface {
	GetName() string
	DataSiswa() string
}

type Siswa struct {
	Name string
}

func (s *Siswa) GetName() string {
	return s.Name
}

func (s *Siswa) DataSiswa() string {
	return s.Name
}

// ini juga masuk ke dalam data karna dia memiliki method GetName() yang sama dengan interface Data
type SiswiInter interface {
	GetName() string
	DataSiswi() string
}

type Siswi struct {
	Name string
}

func (s *Siswi) GetName() string {
	return s.Name
}
func (s *Siswi) DataSiswi() string {
	return s.Name
}

// jika interface NangInter ini tidak memiliki metoh DataName() maka dia tidak bisa di kirim ke func generic
type NangInter interface {
	DataNang() string
}

type Nang struct {
	Name string
}

func (s *Nang) DataNang() string {
	return s.Name
}

func TestInherenGenerict(t *testing.T) {
	assert.Equal(t, "farid", GolangGeneericInheren[SiswaInter](&Siswa{Name: "farid"}))
	assert.Equal(t, "anang", GolangGeneericInheren[SiswiInter](&Siswi{Name: "anang"}))
	// ini tidak bis di kirim karna tidak memiliki methid GetName()
	// assert.Equal(t, "samudra", GolangGeneericInheren[NangInter](&Nang{Name: "samudra"}))
}
