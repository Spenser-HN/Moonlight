package Array

import (
	"errors"
	"log"
	"reflect"

	//JavaScript

	"com.moonlight/app/src/JavaScript/Number"
	"com.moonlight/app/src/JavaScript/Object"
)

//Globals
var Global_Number *Number.New = &Number.New{}

//array class
type New struct {
	Values []any
}

type Entrie struct {
	I   int
	Val []*Entrie_Value
}

type Entrie_Value struct {
	Index int
	Value any
}

//Length
func (this *New) Length() int {
	return len(this.Values)
}

//Reverse Array
func (this *New) Reverse() []any {

	length := this.Length()

	if length > 1 {

		var new_arr []any

		for r := (length - 1); r > -1; r-- {
			new_arr = append(new_arr, this.Values[r])
		}

		return new_arr

		//Empty
	} else {
		return (make([]any, 0))
	}
}

//Array.forEach
func (this *New) ForEach(callback func(Param any, i int)) {

	length := this.Length()

	for i := 0; i < length; i++ {
		callback(this.Values[i], i)
	}
}

//Array.Map
func (this *New) Map(callback func(Param any, i int) any) []any {
	length := this.Length()
	values := make([]any, 0)

	for i := 0; i < length; i++ {
		values = append(values, callback(this.Values[i], i))
	}

	return values
}

//Array.Fill
func (this *New) Fill(Element any, start int, end int) error {

	length := this.Length()

	if end-start <= 0 || start-end >= 0 {
		return errors.New("Error: Start is ( higger || lower || equal) than End or viceversa")
	}

	if length == 1 {
		this.Values[0] = Element
	}

	if start < 0 {
		start = (length + start)

		if start < 0 {
			return errors.New("Error: Start is higger or lower than Array length")
		}

	}

	if end < 0 {
		end = (length + end)

		if end < 0 {
			return errors.New("Error: Start is higger or lower than Array length")
		}

	}

	for i := start; i < end; i++ {
		this.Values[i] = Element
	}

	return nil

}

//Array.Push
func (this *New) Push(Element any) {
	this.Values = append(this.Values, Element)
}

//Array.Pop
func (this *New) Pop() (any, error) {

	length := this.Length()

	if length > 1 {

		new_arr := make([]any, 0)

		for i := 0; i < (length - 1); i++ {
			new_arr = append(new_arr, this.Values[i])
		}

		var last_item any

		last_item = this.Values[(length - 1)]

		this.Values = new_arr

		return last_item, nil

	} else if length == 1 {

		var last_item any

		last_item = this.Values[0]

		this.Values = make([]any, 0)

		return last_item, nil
	} else {
		var last_item any
		return last_item, errors.New("Cannot use method (pop) in an empty array")

	}

}

//Array.At
func (this *New) At(index int) (any, error) {

	length := this.Length()

	if length == 0 || index > length {
		return this.Values[0], errors.New("Cannot use method (at) in an empty array")
	}

	if length >= 1 && index == 0 {
		return this.Values[0], nil
	}

	if index < 0 && (length+index) >= 0 {
		return this.Values[(length + index)], nil
	}

	return this.Values[index], nil

}

//Array.Concat
func (this *New) Concat(elements ...[]any) []any {

	var new_arr []any

	new_arr = make([]any, 0)

	new_arr = append(new_arr, this.Values...)

	for i := 0; i < len(elements); i++ {
		new_arr = append(new_arr, elements[i]...)
	}

	return new_arr
}

//Array.CopyWithin
func (this *New) CopyWithin(target int, start int, end int) ([]any, error) {

	arr_length := this.Length()

	if arr_length == 0 {
		return this.Values, errors.New("Target is out of range")
	}

	if target > arr_length {
		return this.Values, errors.New("Target is out of range")
	}

	if start > arr_length || end > arr_length {
		return this.Values, errors.New("Start or End is out of range")
	}

	if start < 0 {

		start = arr_length + start

		if start < 0 {
			return this.Values, errors.New("Start is out of range")
		}

	}

	if end < 0 {

		end = arr_length + end

		if end < 0 {
			return this.Values, errors.New("End is out of range")
		}

	}

	if target < 0 {
		target = arr_length + target
	}

	if target == start {
		return this.Values, errors.New("Cannot copy any element because target value is same than start value")
	}

	if start == end {
		this.Values[start] = this.Values[target]
		return this.Values, nil
	}

	for i := start; i < end; i++ {
		this.Values[i] = this.Values[target]
	}

	return this.Values, nil

}

//Every
func (this *New) Every(callback func(value any, index int) bool) bool {

	var _Every bool

	_Every = false

	for i := 0; i < this.Length(); i++ {
		result := callback(this.Values[i], i)

		if result == true {
			_Every = true
		} else {
			_Every = false
			break
		}

	}

	return _Every
}

//Some
func (this *New) Some(callback func(value any, index int) bool) bool {

	var _Status bool

	_Status = false

	for i := 0; i < this.Length(); i++ {
		result := callback(this.Values[i], i)

		if result == true {
			_Status = true
			break
		}

	}

	return _Status
}

//Entries

func (this *New) Entries() *Entrie {
	iterador := &Entrie{I: 0}

	for i := 0; i < len(this.Values); i++ {
		iterador.Val = append(iterador.Val, &Entrie_Value{Index: i, Value: this.Values[i]})
	}

	return iterador
}

//Entrie.Next Generator function
func (this *Entrie) Next() *Entrie_Value {

	if len(this.Val) < this.I+1 {
		return this.Val[this.I]
	}

	val := this.Val[this.I]
	this.I++
	return val

}

//Filter
func (this *New) Filter(callback func(value any, index int) bool) []any {

	var _Filtered []any = make([]any, 0)

	length := this.Length()

	for i := 0; i < length; i++ {
		result := callback(this.Values[i], i)

		if result == true {
			_Filtered = append(_Filtered, this.Values[i])
		}

	}

	return _Filtered
}

//Find
func (this *New) Find(callback func(value any, index int) bool) any {

	length := this.Length()

	var Finded any

	for i := 0; i < length; i++ {
		result := callback(this.Values[i], i)

		if result == true {
			Finded = this.Values[i]
			break
		}

	}

	return Finded
}

//FindLast
func (this *New) FindLast(callback func(value any, index int) bool) any {

	length := this.Length()

	var Finded any

	for i := 0; i < length; i++ {
		result := callback(this.Values[i], i)

		if result == true {
			Finded = this.Values[i]
		}

	}

	return Finded
}

//FindIndex
func (this *New) FindIndex(callback func(value any, index int) bool) int {

	length := this.Length()

	var Finded int

	//Equivalent to not found
	Finded = -1

	for i := 0; i < length; i++ {
		result := callback(this.Values[i], i)

		if result == true {
			Finded = i
			break
		}

	}

	return Finded
}

//FindLastIndex
func (this *New) FindLastIndex(callback func(value any, index int) bool) int {

	length := this.Length()

	var Finded int

	//Equivalent to not found
	Finded = -1

	for i := 0; i < length; i++ {
		result := callback(this.Values[i], i)

		if result == true {
			Finded = i
		}

	}

	return Finded
}

//Pointer
func (this *New) Pointer() **New {
	return &this
}

//IsArray
//(i) Remember always pass nil as argument if you want a self verification for example array.IsArray(nil);
//so it will find if array variable is an array else pass anything you would like to check.
//Alse remember that slices are treated as arrays in JS so here the IsArray method will do the same;
func (this *New) IsArray(obj any) bool {

	if _, condition := obj.(New); condition == true {
		return true
	} else {
		return false
	}

}

//Flat
func (this *New) Flat(Depth int) any {

	if Depth <= 0 {
		Flattend := this.Values
		return Flattend
	}

	Flattend := flattenDeep(nil, reflect.ValueOf(this.Values), Depth)

	return Flattend

}

//To use flat map
func flattenDeep(args []interface{}, v reflect.Value, Depth int) []interface{} {
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	if (v.Kind() == reflect.Array || v.Kind() == reflect.Slice) && Depth > 0 {
		for i := 0; i < v.Len(); i++ {
			args = flattenDeep(args, v.Index(i), Depth-1)
		}
	} else if (v.Kind() == reflect.Array || v.Kind() == reflect.Slice) && Depth <= 0 {

		for i := 0; i < v.Len(); i++ {
			args = append(args, v.Index(i))
		}

	} else {

		args = append(args, v.Interface())

	}

	return args
}

//FlatMap
func (this *New) FlatMap(Callback func(Element any) []any) any {

	var Pre_Array *New

	Pre_Array = &New{Values: []any{}}

	this.ForEach(func(element any, i int) {
		Pre_Array.Values = append(Pre_Array.Values, Callback(element))
	})

	return Pre_Array.Flat(2)
}

//From
//func(this *Array[T])From()

//Group
//GroupToMap
//Includes
//IndexOf
//LastIndexOf

// ? ----- Crear el jueves -----
//Keys
//Of

//Reduce
func (this *New) Reduce(Callback func(acc any, cur any, idx int) any, InitVal int) any {

	length := this.Length()

	if length == 0 {
		return this.Values[0]
	}

	var ACC any

	for i := InitVal; i < length; i++ {
		ACC = Callback(ACC, this.Values[i], i)
	}

	return ACC
}

//ReduceRight
//Shift

//Slice
func (this *New) Slice(start int, end int) []any {

	if start < 0 {
		start = (this.Length() + start)
	}

	if end < 0 {
		end = (this.Length() + end)
	}

	if start > end || start > (this.Length()-1) {
		log.Fatal("start is higger than end or is out of range")
	}

	if end > this.Length() {
		log.Fatal("end is out of range")
	}

	return this.Values[start:end]

}

//Splice
func (this *New) Splice(start int, end int, Item ...any) []any {

	if start < 0 {
		start = (this.Length() + start)
	}

	if end < 0 {
		end = (this.Length() + end)
	}

	if start > end || start > (this.Length()-1) {
		log.Fatal("start is higger than end or is out of range")
	}

	if end > (this.Length() - 1) {
		log.Fatal("end is out of range")
	}

	new_arr := this.Values[:start]

	if start == end {

		new_arr = append(new_arr, Item...)

		//Updating array
		this.Values = new_arr

		return new_arr

	} else {

		for i := start; i < end; i++ {
			new_arr = append(new_arr, Item...)
		}

	}

	if end == this.Length() {
		//Updating array
		this.Values = new_arr

		return new_arr
	}

	//else push the other items
	new_arr = append(new_arr, this.Values[end:]...)

	//Updating array
	this.Values = new_arr
	return new_arr

}

//Sort

// ? ----- Crear el viernes ------

//ToLocalString
//ToString

func (this *New) ToString(Arr any) string {

	_Array, condition := Arr.(New)

	if condition != true {
		log.Fatal("Elements are has not a valid Javascript type")
	}

	if Arr == nil {
		_Array = *this
	}

	_Array.Flat(4)

	var values string

	for _, value := range _Array.Values {

		//String
		if value, _isString := value.(string); _isString == true {
			values = AddToString(values, value, ",")
		}

		//Number
		if value, _isNumber := value.(Number.New); _isNumber == true {
			values = AddToString(values, Global_Number.ToString(value), ",")
		}

		//Pointer
		if reflect.ValueOf(value).Kind().String() == "ptr" {
			values = AddToString(values, "[Pointer]", ",")
		}

		//Objects

		// ? --- Object ----
		if _, _isObject := value.(Object.New); _isObject == true {
			values = AddToString(values, "[object Object]", ",")
		}

		// ? --- Map ----

	}

	if len(values) == 0 {
		log.Fatal("Elements are has not a valid Javascript type")
	}

	return values

}

func (this *New) Join(Separator any) string {

	if Separator == nil {
		Separator = ","
	}

	separator, _isString := Separator.(string)

	if _isString != true {
		log.Fatal("Elements are not a valid Javascript type")
	}

	var ARR New = *this
	ARR.Flat(4)

	var values string

	for _, value := range ARR.Values {

		//String
		if value, _isString := value.(string); _isString == true {
			values = AddToString(values, value, separator)
		}

		//Number
		if value, _isNumber := value.(Number.New); _isNumber == true {
			values = AddToString(values, Global_Number.ToString(value), separator)
		}

		//Pointer
		if reflect.ValueOf(value).Kind().String() == "ptr" {
			values = AddToString(values, "[Pointer]", separator)
		}

		//Objects

		// ? --- Object ----
		if _, _isObject := value.(Object.New); _isObject == true {
			values = AddToString(values, "[object Object]", separator)
		}

		// ? --- Map ----

	}

	if len(values) == 0 {
		log.Fatal("Elements are not a valid Javascript type")
	}

	return values

}

func AddToString(ActualValue, ValueToConcat, Divisor string) string {
	if len(ActualValue) > 0 {
		ActualValue += (Divisor + ValueToConcat)
	} else {
		ActualValue = ValueToConcat
	}

	return ActualValue

}

//Unshift
//Values

// ? ------ Sabado Empezar con los String

// ? ------ Proximo miercoles empezar con los Number

// ? ------ Proximo Sabado empezar con los Object

// ? ------ Siguiente Proximo miercoles empezar con el transpilador

// ? ------ Siguiente Proximo Sabado manejar Funciones Nativas como fetch, while, SetTimeOut, SetInterval

// ? ------ Crear el modulo http de node usando el modulo nativo net/http de go

// ? ------- Crear el module fs de node usando el modulo nativo os de go.

// ? ------- Presentarlo en la UNAH y hacerlo OpenSource en Github

// ? ------- Crear el modulo Crypto,etc.. de node

// ? ------- Manejar WebApis
