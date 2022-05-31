package utils

import (
	"reflect"
)

func Strcut2Map(s interface{}, tag string) map[string]interface{} {
	keys := reflect.TypeOf(s)
	values := reflect.ValueOf(s)
	m := map[string]interface{}{}
	for i := 0; i < keys.NumField(); i++ {
		if values.Field(i).CanInterface() {
			m[keys.Field(i).Tag.Get(tag)] = values.Field(i).Interface()
		}
	}
	return m
}

func Strcut2MapExceptZero(s interface{}, tag string) map[string]interface{} {
	keys := reflect.TypeOf(s)
	values := reflect.ValueOf(s)
	m := map[string]interface{}{}
	for i := 0; i < keys.Elem().NumField(); i++ {
		key := keys.Elem().Field(i).Tag.Get(tag)
		if key != "" && values.Elem().Field(i).CanInterface() && !values.Elem().Field(i).IsZero() {
			m[key] = values.Elem().Field(i).Interface()
		}
	}
	return m
}
