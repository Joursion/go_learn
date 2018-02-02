package main

import (
	"reflect"
)

func Call (o interface{}) {
	v := reflect.ValueOf(stu)
	mv := v.MethodByName("SayHi")
}