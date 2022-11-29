package Console

import (
	"fmt"
	"log"
	"os"
	"reflect"

	"com.moonlight/app/src/JavaScript"
	"com.moonlight/app/src/JavaScript/Number"
	"com.moonlight/app/src/JavaScript/Object"
)

// ? ----- Dont delete this flags because they are needed -----
var flags = log.LstdFlags | log.Lshortfile

var Global_Object *Object.New = &Object.New{Values: nil}
var Global_Number *Number.New = &Number.New{}

func HandleNumber(Number Number.New) string {
	return "\x1b[33m" + Global_Number.ToString(Number) + "\x1b[0m"
}

func HandleFunction() string {
	return ("\x1b[36m[Function]\x1b[0m")
}

func HandleObject(Obj Object.New) string {

	var object string = "Object: { "

	initial_lentght := len(object)

	entries := Global_Object.Entries(Obj)

	if len(entries) > 1 {
		object += "\n"
	}

	space := ""
	for i := 0; i < initial_lentght; i++ {
		space += " "
	}

	for key, entrie := range entries {

		var (
			key   string = key
			value string = ""
		)

		if _, v := entrie[1].(Object.New); v == true {

			value = "\x1b[36m[Object]\x1b[0m"
			object += space + (key + " : " + value) + ",\n"

		} else if num, v := entrie[1].(Number.New); v == true {

			value = HandleNumber(num)
			object += space + (key + " : " + value) + ",\n"

		} else if val, v := entrie[1].(string); v == true {
			value = "\x1b[32m" + `"` + val + `"` + "\x1b[0m"
			object += space + (key + " : " + value) + ",\n"
		} else if val, v := entrie[1].(bool); v == true {

			if val == true {
				value = "\x1b[33mtrue\x1b[0m"
			} else {
				value = "\x1b[33mfalse\x1b[0m"
			}

			object += space + (key + " : " + value) + ",\n"

		} else if reflect.ValueOf(entrie[1]).Kind().String() == "ptr" {
			value = HandlePointer()
			object += space + (key + " : " + value) + ",\n"
		} else if JavaScript.TypeOf(entrie[1]) == "map" {
			value = "\x1b[36m[Map]\x1b[0m"
			object += space + (key + " : " + value) + ",\n"
		} else if JavaScript.TypeOf(entrie[1]) == "array" {
			value = "\x1b[36m[Array]\x1b[0m"
			object += space + (key + " : " + value) + ",\n"
		} else if JavaScript.TypeOf(entrie[1]) == "function" {
			value = HandleFunction()
			object += space + (key + " : " + value) + ",\n"
		}

	}

	object += "},\n"

	return object

}

func HandlePointer() string {
	return "\x1b[32m[Pointer]\x1b[0m"
}

func Log(args ...any) {
	for i, arg := range args {

		//Pointers
		if reflect.ValueOf(arg).Kind().String() == "ptr" {
			args[i] = HandlePointer()
		}

		//Objects
		if obj, v := arg.(Object.New); v == true {
			Arg := HandleObject(obj)
			args[i] = Arg
		}

		//Numbers
		if num, v := arg.(Number.New); v == true {

			Num := HandleNumber(num)
			args[i] = Num

		}

		//Functions
		if JavaScript.TypeOf(arg) == "function" {
			args[i] = HandleFunction()
		}

	}

	fmt.Println(args...)
}

func Error(args ...any) {
	Error := log.New(os.Stdout, "\x1b[31m Error: \x1b[0m", flags)
	for i, arg := range args {

		//Pointers
		if reflect.ValueOf(arg).Kind().String() == "ptr" {
			args[i] = HandlePointer()
		}

		//Objects
		if obj, v := arg.(Object.New); v == true {
			Arg := HandleObject(obj)
			args[i] = Arg
		}

		//Numbers
		if num, v := arg.(Number.New); v == true {
			Num := HandleNumber(num)
			args[i] = Num
		}

		//Functions
		if JavaScript.TypeOf(arg) == "function" {
			args[i] = HandleFunction()
		}

	}

	Error.Println(args...)
}

func Warn(args ...any) {
	Warn := log.New(os.Stdout, "\x1b[33m Warn: \x1b[0m", flags)
	for i, arg := range args {

		//Pointers
		if reflect.ValueOf(arg).Kind().String() == "ptr" {
			args[i] = HandlePointer()
		}

		//Objects
		if obj, v := arg.(Object.New); v == true {
			Arg := HandleObject(obj)
			args[i] = Arg
		}

		//Numbers
		if num, v := arg.(Number.New); v == true {
			Num := HandleNumber(num)
			args[i] = Num
		}

		//Functions
		if JavaScript.TypeOf(arg) == "function" {
			args[i] = HandleFunction()
		}

	}
	Warn.Println(args...)
}

func Info(args ...any) {
	Info := log.New(os.Stdout, "\x1b[32m Info: \x1b[0m", flags)
	for i, arg := range args {

		//Pointers
		if reflect.ValueOf(arg).Kind().String() == "ptr" {
			args[i] = HandlePointer()
		}

		//Objects
		if obj, v := arg.(Object.New); v == true {
			Arg := HandleObject(obj)
			args[i] = Arg
		}

		//Numbers
		if num, v := arg.(Number.New); v == true {
			Num := HandleNumber(num)
			args[i] = Num
		}

		//Functions
		if JavaScript.TypeOf(arg) == "function" {
			args[i] = HandleFunction()
		}

	}
	Info.Println(args...)
}
