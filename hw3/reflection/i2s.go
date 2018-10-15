package main

import (
	"reflect"
	"errors"
)


func i2s(data interface{}, out interface{}) error {
	if reflect.TypeOf(out).Kind() != reflect.Ptr {
		return errors.New("Is not ptr")
	}
	dataValue := reflect.ValueOf(data)
	dataType := reflect.TypeOf(data)

	switch dataType.Kind() {
	case reflect.Map:
		if reflect.TypeOf(out).Elem().Kind() != reflect.Struct {
			return errors.New("Type is not struct")
		}
		keys := dataValue.MapKeys()
		val := data.(map[string]interface{})
		outValue := reflect.ValueOf(out).Elem()

		for _, key := range keys {
			err := i2s(val[key.String()], outValue.FieldByName(key.String()).Addr().Interface())
			if err != nil {
				return err
			}
		}

	case reflect.Slice:
		if reflect.TypeOf(out).Elem().Kind() != reflect.Slice {
			return errors.New("Type is not slice")
		}
		val := data.([]interface{})
		outValue := reflect.ValueOf(out).Elem()
		outValue.Set(reflect.MakeSlice(outValue.Type(), len(val), cap(val)))
		for i := range val {
			err := i2s(val[i], outValue.Index(i).Addr().Interface())
			if err != nil {
				return err
			}
		}

	case reflect.Float64:
		if reflect.TypeOf(out).Elem().Kind() != reflect.Int {
			return errors.New("Type is not int")
		}
		value := int(dataValue.Float())

		outValue := reflect.ValueOf(out).Elem()
		outValue.Set(reflect.ValueOf(value))

	case reflect.Bool:
		if reflect.TypeOf(out).Elem().Kind() != reflect.Bool {
			return errors.New("Type is not bool")
		}
		value := dataValue.Bool()

		outValue := reflect.ValueOf(out).Elem()
		outValue.Set(reflect.ValueOf(value))

	case reflect.String:
		if reflect.TypeOf(out).Elem().Kind() != reflect.String {
			return errors.New("Type is not string")
		}
		value := dataValue.String()

		outValue := reflect.ValueOf(out).Elem()
		outValue.Set(reflect.ValueOf(value))

	default:
		return errors.New("Unknown type")
	}
	return nil
}