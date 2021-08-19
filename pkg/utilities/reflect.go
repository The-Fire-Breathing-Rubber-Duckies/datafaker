package utilities

import (
	"fmt"
	"reflect"
)

func ReflectValue(Iface interface{}, FieldName string) (interface{}, error) {
	ValueIface := reflect.ValueOf(Iface)

	// Check if the passed interface is a pointer
	if ValueIface.Type().Kind() != reflect.Ptr {
		// Create a new type of Iface's Type, so we have a pointer to work with
		ValueIface = reflect.New(reflect.TypeOf(Iface))
	}

	// 'dereference' with Elem() and get the field by name
	field := reflect.Indirect(ValueIface).FieldByName(FieldName)
	if !field.IsValid() {
		return nil, fmt.Errorf("")
	}

	return field.Interface(), nil
}
