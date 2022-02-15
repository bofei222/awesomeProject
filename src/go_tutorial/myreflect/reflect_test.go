package myreflect

import (
	"fmt"
	"reflect"
	"testing"
)

type User struct {
	Id   int
	Name string `json:"name" id:"100"`
}

func TestTag(t *testing.T) {
	//创建结构体实例
	ins := User{Id: 1, Name: "root"}
	//获取结构体实例的反射类型对象
	typ := reflect.TypeOf(ins)
	//遍历结构体成员
	for i := 0; i < typ.NumField(); i++ {
		//获取成员
		field := typ.Field(i)
		//获取成员属性
		name := field.Name
		tag := field.Tag
		fmt.Printf("name = %v, tag = %v\n", name, tag)

	}
}

func TestA(t *testing.T) {
	user := &User{Id: 1, Name: "root"}

	val := reflect.ValueOf(user)
	fmt.Printf("%v\n", val) //&{1 root}

	elem := val.Elem()
	fmt.Printf("%v\n", elem) //{1 root}

	elemType := elem.Type()
	fmt.Printf("%v\n", elemType) //main.User

	numField := elem.NumField()
	fmt.Printf("%v\n", numField) //2

	for i := 0; i < numField; i++ {
		fieldName := elemType.Field(i).Name

		field := elem.Field(i)
		fmt.Printf("%v\n", field)

		fieldType := field.Type()
		fieldValue := field.Interface()

		fmt.Printf("%d: %s %s = %v\n", i, fieldName, fieldType, fieldValue)
	}

}
