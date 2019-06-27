package structEngin

import (
	"github.com/gohouse/gocar/helper"
	"github.com/gohouse/t"
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
	Fields        []interface{}
	Result        []map[string]interface{}
	TagName       string
	TagIgnoreName string
	ExtraCols     []string
}

func New() *StructEngin {
	s := new(StructEngin)
	s.TagName = "gorose"
	s.TagIgnoreName = "ignore"
	return s
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
			s.AppendFields(valueField.Addr().Interface())
			//} else {
			//	s.setFields(valueField.Interface())
			//}
		}
	}
	return s.GetFields()
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
		default:
			var fieldName = typeField.Tag.Get(s.GetTagName())
			// 如果该字段没有被忽略, 则获取值
			if fieldName != s.GetTagIgnoreName() {
				// 如果tag为空, 则获取字段名字
				if fieldName == "" {
					fieldName = typeField.Name
				}
				// 如果是struct字段类型的默认值, 则不获取
				if t.New(valueField.Interface()).Bool() {
					mapTmp[fieldName] = valueField.Interface()
				} else {
					// 如果指定了强制获取, 则也获取
					if helper.InArray(fieldName, s.ExtraCols) {
						mapTmp[fieldName] = valueField.Interface()
					}
				}
			}
		}
	}
	s.AppendResult(mapTmp)
}

func (s *StructEngin) AppendFields(arg interface{}) {
	s.Fields = append(s.Fields, arg)
}

func (s *StructEngin) SetFields(arg []interface{}) {
	s.Fields = arg
}
func (s *StructEngin) SetExtraCols(args []string) *StructEngin {
	s.ExtraCols = args
	return s
}

func (s *StructEngin) GetFields() []interface{} {
	return s.Fields
}

func (s *StructEngin) AppendResult(arg map[string]interface{}) {
	s.Result = append(s.Result, arg)
}

func (s *StructEngin) SetResult(arg []map[string]interface{}) {
	s.Result = arg
}

func (s *StructEngin) GetResult() []map[string]interface{} {
	return s.Result
}

func (s *StructEngin) SetTagName(arg string) *StructEngin {
	s.TagName = arg
	return s
}

func (s *StructEngin) GetTagName() string {
	return s.TagName
}

func (s *StructEngin) SetTagIgnoreName(arg string) *StructEngin {
	s.TagIgnoreName = arg
	return s
}

func (s *StructEngin) GetTagIgnoreName() string {
	return s.TagIgnoreName
}
