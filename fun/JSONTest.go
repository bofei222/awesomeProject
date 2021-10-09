package main

import (
	"github.com/thedevsaddam/gojsonq"
)

func main() {
	const json = `{"name":{"first":"Tom","last":"Hanks"},"age":61}`
	name := gojsonq.New().FromString(json).Find("name.first")
	println(name.(string)) // Tom

}
