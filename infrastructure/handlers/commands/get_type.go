package commands

import "reflect"

func GetType(t interface{}) reflect.Type {
	v := reflect.ValueOf(t)
	if reflect.Pointer == v.Kind() {
		v = v.Elem()
	}
	return v.Type()
}
