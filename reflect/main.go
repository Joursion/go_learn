package main

import (
	"fmt"
	"reflect"
	"log"
)

type user struct {
	id int "id"
	name string "name"
}

type Student struct {
	Id int
	Name string
	Sex bool
	Grade float32
}

func (s Student) SayHi() {
	fmt.Printf("Hi~")
}

func (s Student) MyName() {
	fmt.Println("My name is %s", s.Name)
}

func StructInfo (o interface{}) {
	t := reflect.TypeOf(o)

	if k := t.Kind(); k != reflect.Struct {
		fmt.Printf("Tish value is not a struct, it's %v", k)
		return
	}
	fmt.Println("Struct name is :", t.Name())
	fmt.Println("Fields of the struct is:")
	v := reflect.ValueOf(o)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()
		// Anonymous
		fmt.Printf("%6s: %v = %v \n", field.Name, field.Type, value)
	}
	fmt.Println("Method of the struct is:")
	for i := 0; i < t.NumMethod(); i ++{
		method := t.Method(i)
		fmt.Printf("%6s: %v\n", method.Name, method.Type)
	}
}


func main () {
	u := user{100, "Eric"}
	t := reflect.TypeOf(u)
	v := reflect.ValueOf(u)

	for i := 0; i < t.NumField(); i ++ {
		f := t.Field(i)
		log.Fatal(t, v, f)
	}


	stu := Student{10001, "Eric", false, 90.5}
	StructInfo(stu)

	SetValue(stu)
	fmt.Println(stu)
}