package defaults

import (
	"errors"
	"reflect"
	"strconv"
)

var errInvalidFieldType = errors.New("invalid field type")

const defaultFieldName = "default"

func NewWithDefaults(src interface{}) interface{} {
	v := reflect.ValueOf(src)
	t := reflect.TypeOf(src)
	clone := reflect.New(reflect.TypeOf(src))

	for i := 0; i < v.NumField(); i++ {
		tField := t.Field(i)
		cloneField := clone.Elem().Field(i)
		defaultVal := tField.Tag.Get("default")
		if defaultVal == "" {
			continue
		}

		var iface interface{}
		var err error

		switch cloneField.Kind() {
		case reflect.Bool:
			iface, err = strconv.ParseBool(defaultVal)
		case reflect.Int:
			iface, err = strconv.ParseInt(defaultVal, 10, 64)
			iface = int(iface.(int64))
		case reflect.Int8:
			iface, err = strconv.ParseInt(defaultVal, 10, 8)
			iface = int8(iface.(int64))
		case reflect.Int16:
			iface, err = strconv.ParseInt(defaultVal, 10, 16)
			iface = int16(iface.(int64))
		case reflect.Int32:
			iface, err = strconv.ParseInt(defaultVal, 10, 32)
			iface = int32(iface.(int64))
		case reflect.Int64:
			iface, err = strconv.ParseInt(defaultVal, 10, 64)
		case reflect.Uint:
			iface, err = strconv.ParseUint(defaultVal, 10, 64)
			iface = uint(iface.(uint64))
		case reflect.Uint8:
			iface, err = strconv.ParseUint(defaultVal, 10, 8)
			iface = uint8(iface.(uint64))
		case reflect.Uint16:
			iface, err = strconv.ParseUint(defaultVal, 10, 16)
			iface = uint16(iface.(uint64))
		case reflect.Uint32:
			iface, err = strconv.ParseUint(defaultVal, 10, 32)
			iface = uint32(iface.(uint64))
		case reflect.Uint64:
			iface, err = strconv.ParseUint(defaultVal, 10, 64)
		case reflect.Uintptr:
			iface, err = strconv.ParseUint(defaultVal, 10, 64)
			iface = uintptr(iface.(uint64))
		case reflect.Float32:
			iface, err = strconv.ParseFloat(defaultVal, 32)
			iface = float32(iface.(float64))
		case reflect.Float64:
			iface, err = strconv.ParseFloat(defaultVal, 64)
		case reflect.String:
			iface = defaultVal
		default:
			err = errInvalidFieldType
		}

		if err == nil {
			cloneField.Set(reflect.ValueOf(iface))
		}
	}
	return clone.Elem().Interface()
}