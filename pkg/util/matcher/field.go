package matcher

import (
	"fmt"
	"reflect"

	"github.com/golang/mock/gomock"
)

// fieldMatcher is helper let compare field with value
type fieldMatcher struct {
	Key   string
	Value interface{}
}

// NewFieldMatcher create field matcher
func NewFieldMatcher(key string, val interface{}) gomock.Matcher {
	return fieldMatcher{
		Key:   key,
		Value: val,
	}
}

// Matches implementation for Matcher
func (m fieldMatcher) Matches(x interface{}) bool {
	val := reflect.ValueOf(x)
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		if field.Name == m.Key {
			if reflect.DeepEqual(getValue(val.Field(i)), m.Value) {
				return true
			}
		}
	}
	return false
}

func (m fieldMatcher) String() string {
	return fmt.Sprintf("obj.%v is equal to %v", m.Key, m.Value)
}

func getValue(v reflect.Value) interface{} {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return v.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint()
	case reflect.Bool:
		return v.Bool()
	case reflect.String:
		return v.String()
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return uint64(v.Pointer())
	default:
		return v.Type().String() + " value"
	}
}
