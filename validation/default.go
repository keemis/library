package validation

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/pkg/errors"
)

const (
	// DefaultTag struct tag
	DefaultTag = "default"
)

// SetDefault 设置默认值
func SetDefault(obj interface{}) error {
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)
	if objV.Kind() != reflect.Ptr || objV.IsNil() || objT.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("%v must be a struct pointer", obj)
	}
	for i := 0; i < objT.Elem().NumField(); i++ {
		tag := objT.Elem().Field(i).Tag.Get(DefaultTag)
		if len(tag) > 0 {
			name := objT.Elem().Field(i).Name
			field := objV.Elem().FieldByName(name)
			if err := setValue(field, tag); err != nil {
				return errors.Wrap(err, name)
			}
		}
	}
	return nil
}

// setValue 赋值
func setValue(field reflect.Value, newVal string) error {
	switch field.Type().String() {
	case "int8":
		if fmt.Sprintf("%v", field.Interface()) != "0" {
			return nil
		}
		nv, err := strconv.Atoi(newVal)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(int8(nv)))

	case "int16":
		if fmt.Sprintf("%v", field.Interface()) != "0" {
			return nil
		}
		nv, err := strconv.Atoi(newVal)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(int16(nv)))

	case "int32":
		if fmt.Sprintf("%v", field.Interface()) != "0" {
			return nil
		}
		nv, err := strconv.Atoi(newVal)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(int32(nv)))

	case "int":
		if fmt.Sprintf("%v", field.Interface()) != "0" {
			return nil
		}
		nv, err := strconv.Atoi(newVal)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(nv))

	case "int64":
		if fmt.Sprintf("%v", field.Interface()) != "0" {
			return nil
		}
		nv, err := strconv.ParseInt(newVal, 10, 64)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(nv))

	case "uint8":
		if fmt.Sprintf("%v", field.Interface()) != "0" {
			return nil
		}
		nv, err := strconv.Atoi(newVal)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(uint8(nv)))

	case "uint16":
		if fmt.Sprintf("%v", field.Interface()) != "0" {
			return nil
		}
		nv, err := strconv.Atoi(newVal)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(uint16(nv)))

	case "uint32":
		if fmt.Sprintf("%v", field.Interface()) != "0" {
			return nil
		}
		nv, err := strconv.Atoi(newVal)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(uint32(nv)))

	case "uint":
		if fmt.Sprintf("%v", field.Interface()) != "0" {
			return nil
		}
		nv, err := strconv.Atoi(newVal)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(uint(nv)))

	case "uint64":
		if fmt.Sprintf("%v", field.Interface()) != "0" {
			return nil
		}
		nv, err := strconv.Atoi(newVal)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(uint64(nv)))

	case "float32":
		if fmt.Sprintf("%v", field.Interface()) != "0" {
			return nil
		}
		nv, err := strconv.ParseFloat(newVal, 32)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(float32(nv)))

	case "float64":
		if fmt.Sprintf("%v", field.Interface()) != "0" {
			return nil
		}
		nv, err := strconv.ParseFloat(newVal, 64)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(nv))

	case "string":
		if fmt.Sprintf("%v", field.Interface()) != "" {
			return nil
		}
		field.Set(reflect.ValueOf(newVal))

	case "bool":
		if fmt.Sprintf("%v", field.Interface()) != "false" {
			return nil
		}
		if newVal == "true" || newVal == "1" {
			field.Set(reflect.ValueOf(true))
		} else if newVal == "false" || newVal == "0" {
			field.Set(reflect.ValueOf(false))
		} else {
			return errors.New("value is no bool")
		}

	default:

	}
	return nil
}
