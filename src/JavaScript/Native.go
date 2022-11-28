package JavaScript

import (
	"fmt"
	"reflect"

	"moonlight.js/moonlight/src/utils/JavaScript/Array"
	"moonlight.js/moonlight/src/utils/JavaScript/Number"
	"moonlight.js/moonlight/src/utils/JavaScript/Object"
)

//All types in this file come from "./Types.go" &
//are part of Typescript package
/**
func SetTimeout(callback TimeoutCallback, timeout int){
	callback();
}

func SetInterval(callback TimeIntervalCallback, interval int){
	callback();
}

**/

func While(Condition func() bool, Callback func()) {

	for Condition() {
		Callback()
	}

}

func TypeOf(Item any) string {

	// TypeScript Types

	if _, v := Item.(Object.New); v == true {
		return "object"
	}

	if _, v := Item.(Number.New); v == true {
		return "number"
	}

	if _, v := Item.(Array.New); v == true {
		return "array"
	}

	Type := reflect.ValueOf(Item).Kind().String()

	//Go Native Types

	switch Type {
	case "ptr":
		return TypeOf(reflect.ValueOf(Item).Elem())
	case "string":
		return "string"
	case "struct":
		return "object"
	case "slice":
		return "array"
	case "array":
		return "array"
	case "map":
		return "map"
	case "bool":
		return "boolean"
	case "func":
		return "function"
	case "int":
		return "number"
	case "int16":
		return "number"
	case "int32":
		return "number"
	case "int64":
		return "number"
	case "float32":
		return "number"
	case "float64":
		return "number"
	case "complex64":
		return "number"
	case "complex128":
		return "number"
	default:
		return Type
	}

}

//fetch
func Fetch[T any](url string, Request Request[T]) {
	fmt.Println("url:" + url)
	fmt.Println("Request:\n", Request)
}
