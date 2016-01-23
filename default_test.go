package defaults

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Foo struct {
	Int         int      `default:"1"`
	Int8        int8     `default:"2"`
	Int16       int16    `default:"3"`
	Int32       int32    `default:"4"`
	Int64       int64    `default:"5"`
	Uint        uint     `default:"6"`
	Uint8       uint8    `default:"7"`
	Uint16      uint16   `default:"8"`
	Uint32      uint32   `default:"9"`
	Uint64      uint64   `default:"10"`
	Uintptr     uintptr  `default:"11"`
	Float32     float32  `default:"1.2"`
	Float64     float64  `default:"1.3"`
	BoolTrue    bool     `default:"true"`
	BoolFalse   bool     `default:"false"`
	String      string   `default:"cheese"`
	StructField struct{} `default:"{}"`
	NoDefault   string
}

func TestDefaults(t *testing.T) {
	foo := NewWithDefaults(Foo{}).(Foo)
	assert.Equal(t, foo.Int, int(1))
	assert.Equal(t, foo.Int8, int8(2))
	assert.Equal(t, foo.Int16, int16(3))
	assert.Equal(t, foo.Int32, int32(4))
	assert.Equal(t, foo.Int64, int64(5))
	assert.Equal(t, foo.Uint, uint(6))
	assert.Equal(t, foo.Uint8, uint8(7))
	assert.Equal(t, foo.Uint16, uint16(8))
	assert.Equal(t, foo.Uint32, uint32(9))
	assert.Equal(t, foo.Uint64, uint64(10))
	assert.Equal(t, foo.Uintptr, uintptr(11))
	assert.Equal(t, foo.Float32, float32(1.2))
	assert.Equal(t, foo.Float64, float64(1.3))
	assert.Equal(t, foo.BoolTrue, true)
	assert.Equal(t, foo.BoolFalse, false)
	assert.Equal(t, foo.String, "cheese")
	assert.Equal(t, foo.StructField, struct{}{})
	assert.Equal(t, foo.NoDefault, "")
}
