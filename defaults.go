package defaults

import (
	"errors"
	"reflect"
	"strconv"
	"time"
)

var errInvalidFieldType = errors.New("invalid field type")

const defaultFieldName = "default"

func NewWithDefaults(src interface{}) interface{} {
	t := reflect.TypeOf(src)
	clone := reflect.New(reflect.TypeOf(src))
	for i := 0; i < t.NumField(); i++ {
		setField(clone.Elem().Field(i), t.Field(i).Tag.Get(defaultFieldName))
	}
	return clone.Elem().Interface()
}

func setField(field reflect.Value, defaultVal string) {
	if defaultVal == "" {
		return
	}

	var iface interface{}
	var err error

	switch field.Kind() {
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
		t, err := time.ParseDuration(defaultVal)
		if err == nil {
			iface, err = t, nil
		} else {
			iface, err = strconv.ParseInt(defaultVal, 10, 64)
		}
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
		if field.CanSet() {
			field.Set(reflect.ValueOf(iface))
		}
	}
}
