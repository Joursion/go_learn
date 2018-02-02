package main

import (
	"fmt"
	"reflect"
)

func SetValue(o interface{}) {
	v := reflect.ValueOf(o)
	if v.Kind() != reflect.Ptr || !v.Elem().CanSet() {
		fmt.Println("Cannot set!")
		return
	} else {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++{
		switch v.Field(i).Kind() {
		case reflect.Int: 
			v.Field(i).SetInt(10002)
		case reflect.String:
			v.Field(i).SetString("zhangheng")
		case reflect.Bool:
			v.Field(i).SetBool(true)
		case reflect.Float32:
			v.Field(i).SetFloat(95.9)
		}
	}
}