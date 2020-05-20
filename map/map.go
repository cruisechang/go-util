package maper

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

//GetMapKeys return keys slice with reflect.Value
func GetMapKeys(v interface{}) ([]reflect.Value, error) {
	typeStr := reflect.TypeOf(v).String()
	if !strings.Contains(typeStr, "map") {
		return nil, errors.New("argument is not map")
	}
	keys := reflect.ValueOf(v).MapKeys()

	return keys, nil
}

//GetMapIntKeys return keys slice with int type
func GetMapIntKeys(v interface{}) ([]int, error) {
	typeStr := reflect.TypeOf(v).String()
	if !strings.Contains(typeStr, "map") {
		return nil, errors.New("argument is not map")
	}
	keys := reflect.ValueOf(v).MapKeys()

	resKeys := []int{}
	for _, v := range keys {
		if i, ok := v.Interface().(int); ok {
			resKeys = append(resKeys, i)
		} else {
			return nil, errors.New("key is not int")
		}
	}
	return resKeys, nil
}

//GetMapStringKeys return keys slice with string type
func GetMapStringKeys(v interface{}) ([]string, error) {
	typeStr := reflect.TypeOf(v).String()
	if !strings.Contains(typeStr, "map") {
		return nil, errors.New("argument is not map")
	}
	keys := reflect.ValueOf(v).MapKeys()

	resKeys := []string{}

	for _, v := range keys {
		if i, ok := v.Interface().(string); ok {
			resKeys = append(resKeys, i)
		} else {
			return nil, errors.New("key is not string")
		}

	}
	return resKeys, nil
}

//MapToStruct mapping map to target struct
//Struct must be respoding correct to that map or mapping will failed
func MapToStruct(m map[string]interface{}, stct interface{}) error {
	err := fillStruct(stct, m)
	if err != nil {
		return err
	}
	return nil
}
func fillStruct(stct interface{}, m map[string]interface{}) error {
	for k, v := range m {
		err := setField(stct, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
func setField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		return errors.New("Provided value type didn't match obj field type")
	}

	structFieldValue.Set(val)
	return nil
}
