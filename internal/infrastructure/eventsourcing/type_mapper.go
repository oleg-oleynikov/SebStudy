package eventsourcing

import (
	"fmt"
	"reflect"
)

type DataToType func(map[string]interface{}) interface{}
type TypeToDataWithName func(interface{}) (string, map[string]interface{})

type TypeMapper struct {
	dataToType map[string]DataToType
	typeToData map[reflect.Type]TypeToDataWithName
}

func NewTypeMapper() *TypeMapper {
	return &TypeMapper{
		dataToType: make(map[string]DataToType),
		typeToData: make(map[reflect.Type]TypeToDataWithName),
	}
}

func (tm *TypeMapper) MapEvent(eventType reflect.Type, name string, dt DataToType, td TypeToDataWithName) error {
	if name == "" {
		return fmt.Errorf("need name for type mapping")
	}

	if _, exists := tm.typeToData[eventType]; exists {
		return nil
	}

	tm.dataToType[name] = dt
	tm.typeToData[eventType] = td

	return nil
}

func (tm *TypeMapper) GetDataToType(eventType string) (DataToType, error) {
	if dt, exists := tm.dataToType[eventType]; exists {
		return dt, nil
	}
	return nil, fmt.Errorf("failed to find type mapped with '%s'", eventType)
}

func (tm *TypeMapper) GetTypeToData(t reflect.Type) (TypeToDataWithName, error) {
	if td, exists := tm.typeToData[t]; exists {
		return td, nil
	}
	return nil, fmt.Errorf("failed to find name mapped with '%s'", t)
}
