package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `info:"name" doc:"姓名"`
	Sex  string `info:"sex" doc:"性别"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()
	for i := 0; i < t.NumField(); i++ {
		tagInfo := t.Field(i).Tag.Get("info")
		tagDoc := t.Field(i).Tag.Get("doc")
		fmt.Printf("Field: %s, Tag info: %s, Tag doc: %s\n", t.Field(i).Name, tagInfo, tagDoc)
	}
}

func main() {
	var re resume
	findTag(re)
}
