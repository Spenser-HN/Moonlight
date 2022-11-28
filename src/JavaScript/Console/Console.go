package Console

import (
	"fmt"
	"log"
	"os"
	"reflect"

	"moonlight.js/moonlight/src/utils/JavaScript"
	"moonlight.js/moonlight/src/utils/JavaScript/Number"
	"moonlight.js/moonlight/src/utils/JavaScript/Object"
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

func HandleObject(Obj Object.New, ConsoleMessage string) (string, string) {

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
			ConsoleMessage = "This log contains pointers"
			value = "\x1b[32m[Pointer]\x1b[0m"
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

	return ConsoleMessage, object

}

func Log(args ...any) {
	var PreMessage string
	for i, arg := range args {

		//Pointers
		if reflect.ValueOf(arg).Kind().String() == "ptr" {
			PreMessage = "This log contains pointers"
		}

		//Objects
		if obj, v := arg.(Object.New); v == true {
			ConsoleMessage, Arg := HandleObject(obj, PreMessage)
			args[i] = Arg
			if len(ConsoleMessage) > 1 && len(PreMessage) == 0 {
				PreMessage = ConsoleMessage
			}
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

	if len(PreMessage) > 1 {
		fmt.Println("\x1b[31m [Alert] \x1b[0m ", PreMessage)
	}

	fmt.Println(args...)
}

func Error(args ...any) {
	Error := log.New(os.Stdout, "\x1b[31m Error: \x1b[0m", flags)
	var PreMessage string
	for i, arg := range args {

		//Pointers
		if reflect.ValueOf(arg).Kind().String() == "ptr" {
			PreMessage = "This log contains pointers"
		}

		//Objects
		if obj, v := arg.(Object.New); v == true {
			ConsoleMessage, Arg := HandleObject(obj, PreMessage)
			args[i] = Arg
			if len(ConsoleMessage) > 1 && len(PreMessage) == 0 {
				PreMessage = ConsoleMessage
			}
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

	if len(PreMessage) > 1 {
		fmt.Println("\x1b[31m [Alert] \x1b[0m ", PreMessage)
	}

	fmt.Println(args...)

	Error.Println(args...)
}

func Warn(args ...any) {
	Warn := log.New(os.Stdout, "\x1b[33m Warn: \x1b[0m", flags)
	Warn.Println(args...)
}

func Info(args ...any) {
	Info := log.New(os.Stdout, "\x1b[32m Info: \x1b[0m", flags)
	Info.Println(args...)
}
