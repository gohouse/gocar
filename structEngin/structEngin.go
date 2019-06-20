package structEngin

import (
	"reflect"
)

//import "github.com/gohouse/t"

//type Type int
//
//const (
//	Map Type = iota	// map[t.T]t.T
//	MapString
//	MapInt64
//	MapInterface
//	Slice	// []t.T
//	SliceString
//	SliceInterface
//	SliceInt64
//)
type StructEngin struct {
	IsSlice bool
	Fields  []interface{}
	Result  []map[string]interface{}
}

func New() *StructEngin {
	return new(StructEngin)
}

func (s *StructEngin) GetStructFields(data interface{}) []interface{} {
	val := reflect.Indirect(reflect.ValueOf(data))
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		switch valueField.Kind() {
		case reflect.Struct:
			//if valueField.CanAddr() {
			s.GetStructFields(valueField.Addr().Interface())
			//} else {
			//	s.GetStructFields(valueField.Interface())
			//}
		default:
			//if valueField.CanAddr() {
			s.setFields(valueField.Addr().Interface())
			//} else {
			//	s.setFields(valueField.Interface())
			//}
		}
	}
	return s.GetFields()
}
func (s *StructEngin) setFields(arg interface{}) {
	s.Fields = append(s.Fields, arg)
}
func (s *StructEngin) GetFields() []interface{} {
	return s.Fields
}

func (s *StructEngin) StructContent2Map(data interface{}) []map[string]interface{} {
	val := reflect.Indirect(reflect.ValueOf(data))
	switch val.Kind() {
	case reflect.Struct: // struct
		s.getStructContent(val)
	case reflect.Slice: // []struct
		//eltType := val.Type().Elem()
		//switch eltType.Kind() {
		//case reflect.Struct:
		for i := 0; i < val.Len(); i++ {
			s.getStructContent(reflect.Indirect(val.Index(i)))
		}
		//}
	}
	return s.GetResult()
}
func (s *StructEngin) getStructContent(val reflect.Value) {
	valType := val.Type()
	var mapTmp = make(map[string]interface{}, 0)
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := valType.Field(i)
		switch valueField.Kind() {
		case reflect.Struct:
			continue
			//s.StructContent2Map(valueField.Interface())
		default:
			//s.setFields(valueField.Addr().Interface())
			var fieldName = typeField.Tag.Get("gorose")
			if fieldName == "" {
				fieldName = typeField.Name
			}
			mapTmp[fieldName] = valueField.Interface()
		}
	}
	s.setResult(mapTmp)
}
func (s *StructEngin) setResult(arg map[string]interface{}) {
	s.Result = append(s.Result, arg)
}
func (s *StructEngin) GetResult() []map[string]interface{} {
	return s.Result
}
