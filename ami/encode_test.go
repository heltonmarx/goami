package ami

import (
	"bytes"
	"fmt"
	"math"
	"reflect"
	"testing"

	"github.com/facebookgo/ensure"
)

func TestEncodeBool(t *testing.T) {
	var buf bytes.Buffer

	b := true
	ensure.Nil(t, encode(&buf, reflect.ValueOf(b)))
	ensure.DeepEqual(t, "true\r\n", buf.String())
	buf.Reset()

	b = false
	ensure.Nil(t, encode(&buf, reflect.ValueOf(b)))
	ensure.DeepEqual(t, "false\r\n", buf.String())
}

func TestEncodeInt(t *testing.T) {
	var buf bytes.Buffer

	// int
	i := 11
	ensure.Nil(t, encode(&buf, reflect.ValueOf(i)))
	ensure.DeepEqual(t, fmt.Sprintf("%d\r\n", i), buf.String())
	buf.Reset()

	// Int8
	i8 := math.MaxInt8
	ensure.Nil(t, encode(&buf, reflect.ValueOf(i8)))
	ensure.DeepEqual(t, fmt.Sprintf("%d\r\n", i8), buf.String())
	buf.Reset()

	i8 = math.MinInt8
	ensure.Nil(t, encode(&buf, reflect.ValueOf(i8)))
	ensure.DeepEqual(t, fmt.Sprintf("%d\r\n", i8), buf.String())
	buf.Reset()

	// Int16
	i16 := math.MaxInt16
	ensure.Nil(t, encode(&buf, reflect.ValueOf(i16)))
	ensure.DeepEqual(t, fmt.Sprintf("%d\r\n", i16), buf.String())
	buf.Reset()

	i16 = math.MinInt16
	ensure.Nil(t, encode(&buf, reflect.ValueOf(i16)))
	ensure.DeepEqual(t, fmt.Sprintf("%d\r\n", i16), buf.String())
	buf.Reset()

	//  Int32
	i32 := math.MaxInt32
	ensure.Nil(t, encode(&buf, reflect.ValueOf(i32)))
	ensure.DeepEqual(t, fmt.Sprintf("%d\r\n", i32), buf.String())
	buf.Reset()

	i32 = math.MinInt32
	ensure.Nil(t, encode(&buf, reflect.ValueOf(i32)))
	ensure.DeepEqual(t, fmt.Sprintf("%d\r\n", i32), buf.String())
	buf.Reset()

	// Int64
	i64 := math.MaxInt64
	ensure.Nil(t, encode(&buf, reflect.ValueOf(i64)))
	ensure.DeepEqual(t, fmt.Sprintf("%d\r\n", i64), buf.String())
	buf.Reset()

	i64 = math.MinInt64
	ensure.Nil(t, encode(&buf, reflect.ValueOf(i64)))
	ensure.DeepEqual(t, fmt.Sprintf("%d\r\n", i64), buf.String())
	buf.Reset()

	// Uint8
	ui8 := math.MaxUint8
	ensure.Nil(t, encode(&buf, reflect.ValueOf(ui8)))
	ensure.DeepEqual(t, fmt.Sprintf("%d\r\n", ui8), buf.String())
	buf.Reset()

	// Uint16
	ui16 := math.MaxUint16
	ensure.Nil(t, encode(&buf, reflect.ValueOf(ui16)))
	ensure.DeepEqual(t, fmt.Sprintf("%d\r\n", ui16), buf.String())
	buf.Reset()

	// Uint32
	ui32 := math.MaxUint32
	ensure.Nil(t, encode(&buf, reflect.ValueOf(ui32)))
	ensure.DeepEqual(t, fmt.Sprintf("%d\r\n", ui32), buf.String())
	buf.Reset()

	// Uint64
	var ui64 uint64
	ui64 = math.MaxUint64
	ensure.Nil(t, encode(&buf, reflect.ValueOf(ui64)))
	ensure.DeepEqual(t, fmt.Sprintf("%v\r\n", ui64), buf.String())
	buf.Reset()
}
