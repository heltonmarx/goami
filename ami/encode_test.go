package ami

import (
	"bytes"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"testing"

	"github.com/facebookgo/ensure"
)

func TestEncodeBool(t *testing.T) {
	var buf bytes.Buffer

	b := true
	ensure.Nil(t, encode(&buf, "", reflect.ValueOf(b)))
	ensure.DeepEqual(t, "true\r\n", buf.String())
	buf.Reset()

	b = false
	ensure.Nil(t, encode(&buf, "", reflect.ValueOf(b)))
	ensure.DeepEqual(t, "false\r\n", buf.String())
}

func TestEncodeInt(t *testing.T) {
	var buf bytes.Buffer

	// int
	var i int
	i = 11
	ensure.Nil(t, encode(&buf, "", reflect.ValueOf(i)))
	ensure.DeepEqual(t, fmt.Sprintf("%d\r\n", i), buf.String())
	buf.Reset()

	// Int8
	var i8 int8
	i8 = math.MaxInt8
	ensure.Nil(t, encode(&buf, "", reflect.ValueOf(i8)))
	ensure.DeepEqual(t, fmt.Sprintf("%d\r\n", i8), buf.String())
	buf.Reset()

	i8 = math.MinInt8
	ensure.Nil(t, encode(&buf, "", reflect.ValueOf(i8)))
	ensure.DeepEqual(t, fmt.Sprintf("%d\r\n", i8), buf.String())
	buf.Reset()

	// Int16
	var i16 int16
	i16 = math.MaxInt16
	ensure.Nil(t, encode(&buf, "", reflect.ValueOf(i16)))
	ensure.DeepEqual(t, fmt.Sprintf("%d\r\n", i16), buf.String())
	buf.Reset()

	i16 = math.MinInt16
	ensure.Nil(t, encode(&buf, "", reflect.ValueOf(i16)))
	ensure.DeepEqual(t, fmt.Sprintf("%d\r\n", i16), buf.String())
	buf.Reset()

	//  Int32
	var i32 int32
	i32 = math.MaxInt32
	ensure.Nil(t, encode(&buf, "", reflect.ValueOf(i32)))
	ensure.DeepEqual(t, fmt.Sprintf("%d\r\n", i32), buf.String())
	buf.Reset()

	i32 = math.MinInt32
	ensure.Nil(t, encode(&buf, "", reflect.ValueOf(i32)))
	ensure.DeepEqual(t, fmt.Sprintf("%d\r\n", i32), buf.String())
	buf.Reset()

	// Int64
	var i64 int64
	i64 = math.MaxInt64
	ensure.Nil(t, encode(&buf, "", reflect.ValueOf(i64)))
	ensure.DeepEqual(t, fmt.Sprintf("%d\r\n", i64), buf.String())
	buf.Reset()

	i64 = math.MinInt64
	ensure.Nil(t, encode(&buf, "", reflect.ValueOf(i64)))
	ensure.DeepEqual(t, fmt.Sprintf("%d\r\n", i64), buf.String())
	buf.Reset()

}

func TestEncodeUint(t *testing.T) {
	var buf bytes.Buffer

	// Uint8
	var ui8 uint8
	ui8 = math.MaxUint8
	ensure.Nil(t, encode(&buf, "", reflect.ValueOf(ui8)))
	ensure.DeepEqual(t, fmt.Sprintf("%d\r\n", ui8), buf.String())
	buf.Reset()

	// Uint16
	var ui16 uint16
	ui16 = math.MaxUint16
	ensure.Nil(t, encode(&buf, "", reflect.ValueOf(ui16)))
	ensure.DeepEqual(t, fmt.Sprintf("%d\r\n", ui16), buf.String())
	buf.Reset()

	// Uint32
	var ui32 uint32
	ui32 = math.MaxUint32
	ensure.Nil(t, encode(&buf, "", reflect.ValueOf(ui32)))
	ensure.DeepEqual(t, fmt.Sprintf("%d\r\n", ui32), buf.String())
	buf.Reset()

	// Uint64
	var ui64 uint64
	ui64 = math.MaxUint64
	ensure.Nil(t, encode(&buf, "", reflect.ValueOf(ui64)))
	ensure.DeepEqual(t, fmt.Sprintf("%v\r\n", ui64), buf.String())
	buf.Reset()
}

func TestEncodeFloat(t *testing.T) {
	var buf bytes.Buffer

	var f32 float32
	f32 = math.MaxFloat32
	ensure.Nil(t, encode(&buf, "", reflect.ValueOf(f32)))
	s := strconv.FormatFloat(float64(f32), 'E', -1, 32) + "\r\n"
	ensure.DeepEqual(t, s, buf.String())
	buf.Reset()

	var f64 float64
	f64 = math.MaxFloat64
	ensure.Nil(t, encode(&buf, "", reflect.ValueOf(f64)))
	s = strconv.FormatFloat(f64, 'E', -1, 64) + "\r\n"
	ensure.DeepEqual(t, s, buf.String())
	buf.Reset()
}

func TestEncodeMap(t *testing.T) {
	var buf bytes.Buffer
	m := map[string]interface{}{
		"Name":  "foobar",
		"Age":   99,
		"Valid": true,
	}
	expect := "Name: foobar\r\nAge: 99\r\nValid: true\r\n"
	ensure.Nil(t, encode(&buf, "", reflect.ValueOf(m)))
	verifyResponse(t, buf.String(), expect)
}

func TestEncodeStruct(t *testing.T) {
	var buf bytes.Buffer
	st := struct {
		Action string `ami:"Action"`
		ID     string `ami:"ActionID"`
		Foo    int    `ami:"Foo,omitempty"`
		Bar    string `ami:"-"`
	}{
		Action: "A", ID: "B", Bar: "C",
	}
	expect := "Action: A\r\nActionID: B\r\n"
	ensure.Nil(t, encode(&buf, "", reflect.ValueOf(st)))
	ensure.DeepEqual(t, expect, buf.String())
}
