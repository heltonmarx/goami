package ami

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, "", reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	buf.WriteString("\r\n")
	return buf.Bytes(), nil
}

func writeString(buf *bytes.Buffer, tag, value string) {
	if tag != "" {
		buf.WriteString(tag)
		buf.WriteString(": ")
	}
	buf.WriteString(value)
	buf.WriteString("\r\n")
}

func encode(buf *bytes.Buffer, tag string, v reflect.Value) error {
	switch v.Kind() {
	case reflect.String:
		writeString(buf, tag, v.String())
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		writeString(buf, tag, strconv.FormatInt(v.Int(), 10))
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		writeString(buf, tag, strconv.FormatUint(v.Uint(), 10))
	case reflect.Bool:
		writeString(buf, tag, strconv.FormatBool(v.Bool()))
	case reflect.Float32:
		writeString(buf, tag, strconv.FormatFloat(v.Float(), 'E', -1, 32))
	case reflect.Float64:
		writeString(buf, tag, strconv.FormatFloat(v.Float(), 'E', -1, 64))
	case reflect.Ptr, reflect.Interface:
		if !v.IsNil() {
			return encode(buf, tag, v.Elem())
		}
	case reflect.Struct:
		return encodeStruct(buf, v)
	case reflect.Map:
		return encodeMap(buf, v)
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			elem := v.Index(i)
			if !isZero(elem) {
				if err := encode(buf, tag, elem); err != nil {
					return err
				}
			}
		}
	default:
		return fmt.Errorf("unsupported kind %v", v.Kind())
	}
	return nil
}

func isOmitempty(tag string) (string, bool, error) {
	fields := strings.Split(tag, ",")
	if len(fields) > 1 {
		for _, flag := range fields[1:] {
			if strings.TrimSpace(flag) == "omitempty" {
				return fields[0], true, nil
			}
			return tag, false, fmt.Errorf("unsupported flag %q in tag %q", flag, tag)
		}
	}
	return tag, false, nil
}

func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return len(v.String()) == 0
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		return v.Int() == 0
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		return v.Uint() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Struct:
		for i := v.NumField() - 1; i >= 0; i-- {
			if !isZero(v.Field(i)) {
				return false
			}
		}
		return true
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	}
	return false
}

func encodeStruct(buf *bytes.Buffer, v reflect.Value) error {
	var omitempty bool
	var err error
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		tag, ok := field.Tag.Lookup("ami")
		switch {
		case !ok:
			tag = string(field.Tag)
		case tag == "-":
			continue
		}
		tag, omitempty, err = isOmitempty(tag)
		if err != nil {
			return err
		}
		value := v.Field(i)
		if omitempty && isZero(value) {
			continue
		}

		if err := encode(buf, tag, value); err != nil {
			return err
		}
	}
	return nil
}

func encodeMap(buf *bytes.Buffer, v reflect.Value) error {
	for _, key := range v.MapKeys() {
		value := v.MapIndex(key)
		if key.Kind() == reflect.String {
			tag := key.String()
			if err := encode(buf, tag, value); err != nil {
				return err
			}
		}
	}
	return nil
}
