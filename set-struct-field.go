package main

import (
  "reflect"
  "fmt"
)

func Set(object *struct, fieldName string, value interface{}) (err error) {
  // pointer to struct - addressable
  pointerStruct := reflect.ValueOf(&object)
  // struct
  structElem := pointerStruct.Elem()
  if structElem.Kind() != reflect.Struct {
    return err
  }
  // exported field
  field := structElem.FieldByName("N")
  if !field.IsValid() { 
    return err
  }
  // A Value can be changed only if it is 
  // addressable and was not obtained by 
  // the use of unexported struct fields.
  if !field.CanSet() { 
    return err
  }
  // change value of N
  switch field.Kind() {
    case reflect.Int:
      // error if the type is different
      if reflect.TypeOf(value) != "int" {
        return  fmt.Errorf("Parmeter type is not the same as the field type")
      }
      field.SetInt(value)
    case reflect.String:
      // error if the type is different
      if reflect.TypeOf(value) != "string" {
        return  fmt.Errorf("Parmeter type is not the same as the field type")
      }
      field.SetString(value)
  }
}