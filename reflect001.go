package main

import (
	"fmt"
	"reflect"
)

type Bird struct {
	BiName string
	Weight float32
}

func main() {
	bi := &Bird{"buGuNiao", 0.32}
	s := reflect.ValueOf(bi).Elem()

	for i := 0; i < s.NumField(); i++ {
		p := s.Field(i)
		fmt.Printf("no:%d, type:%s,\t interface:%v,\r\n", i, p.Type(), p.Interface())
	}

	typeT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		fmt.Printf("no:%d, typeT:%s\r\n", i, typeT.Field(i).Name)
	}
	fmt.Println("Hello World!")
}
