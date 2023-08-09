package main

import (
	"fmt"
	"reflect"
)

type user struct {
	Id       string `json:"id"`
	Pw       string `json:"pw"`
	Nickname string `json:"nickname"`
}

func main() {
	u := user{
		Id:       "testId",
		Pw:       "testPw",
		Nickname: "Aivyss",
	}

	ref := reflect.TypeOf(u) // you can get reflect.Type
	field, _ := ref.FieldByName("Id")
	fmt.Println("field.Type =", field.Type)
	fmt.Println("field.Name =", field.Name)
	fmt.Println("field.Tag =", field.Tag)
	fmt.Println("field.PkgPath =", field.PkgPath)

	pointerStruct := reflect.ValueOf(&u) // you can get reflect.Value
	s := pointerStruct.Elem()
	if s.Kind() == reflect.Struct {
		f := s.FieldByName("Nickname")
		if f.IsValid() && f.CanSet() && f.Kind() == reflect.String {
			f.SetString("aivyss")
		}
	}

	fmt.Println(u)
}
