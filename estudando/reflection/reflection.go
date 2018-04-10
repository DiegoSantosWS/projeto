package main

import (
	"fmt"
	"reflect"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name, omitempty"`
	Email    string `json:"email, omitempty"`
	Password string `json:"password, omitempty"`
}

func main() {
	//pegando os metadados da struct por reflection
	var u User

	f := reflect.Indirect(reflect.ValueOf(u))

	for i := 0; i < f.NumField(); i++ {
		fmt.Println("Nome:", f.Type().Field(i).Name)
		fmt.Println("Tag:", f.Type().Field(i).Tag)
		fmt.Println("Index:", f.Type().Field(i).Index)
		fmt.Println("Offset:", f.Type().Field(i).Offset)
		fmt.Println("Type:", f.Type().Field(i).Type)
		fmt.Println("-=-=-=-=-=-=-=-=-STRUCT=-=-=-=-=-=-=-=-=-=-=-")
	}

	// Pegando metadados da interface pro reflection

	var iface interface{}

	iface = u

	f = reflect.Indirect(reflect.ValueOf(iface))
	for i := 0; i < f.NumField(); i++ {
		fmt.Println("Nome:", f.Type().Field(i).Name)
		fmt.Println("Tag:", f.Type().Field(i).Tag)
		fmt.Println("Index:", f.Type().Field(i).Index)
		fmt.Println("Offset:", f.Type().Field(i).Offset)
		fmt.Println("Type:", f.Type().Field(i).Type)
		fmt.Println("-=-=-=-=-=-=-=-INTERFACE=-=-=-=-=-=-=-=-=-=-=-=-")
	}
}
