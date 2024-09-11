package infrastructure

import "reflect"

func GetValueType(t interface{}) reflect.Type {
	v := reflect.ValueOf(t)
	if reflect.Pointer == v.Kind() {
		v = v.Elem()
	}
	return v.Type()
}
