package Object

//object
type New struct {
	Values     map[string]Object_Property[any]
	Properties map[string]Object_Property[func()]
	Frozen     bool
}

type Object_Property[T any] struct {
	Enumerable bool
	Value      T
}

type object interface {
	Freeze() bool
	GetOwnPropertyNames(Object New) []string
	Keys(Object New) []string
	DefineProperties_Function(Object New, Properties ...Function_Properties) //void
	DefineProperties_Default(Object New, Properties ...Function_Properties)  //void
	Entries(Object New) Entries_Return
}

/**

@method Freeze

@Description
The Object.freeze() method freezes an object. Freezing an object prevents extensions
and makes existing properties non-writable and non-configurable. A frozen object can
no longer be changed: new properties cannot be added, existing properties cannot be removed,
their enumerability, configurability, writability, or value cannot be changed, and the
object's prototype cannot be re-assigned. freeze() returns the same object that was passed in.
Freezing an object is the highest integrity level that JavaScript provides.

**/

func (this *New) Freeze() bool {

	if !this.Frozen {
		this.Frozen = true
		return this.Frozen
	} else {
		this.Frozen = false
		return this.Frozen
	}

}

/**
Returns the names of the own properties of an object.
The own properties of an object are those that are defined directly on that object,
and are not inherited from the object's prototype.
The properties of an object include both fields (objects) and functions.
**/

func (this *New) GetOwnPropertyNames(Object New) []string {

	//Getting Keys
	keys := make([]string, 0, len(Object.Values))

	for k := range Object.Values {
		keys = append(keys, k)
	}

	for property := range Object.Properties {
		keys = append(keys, property)
	}

	//Returning Keys
	return keys

}

/*
* Returns the names of the enumerable string properties and methods of an object.
 */

func (this *New) Keys(Object New) []string {

	//Getting Keys
	keys := make([]string, 0, len(Object.Values))

	for k := range Object.Values {

		if Object.Values[k].Enumerable == true {
			keys = append(keys, k)
		}

	}

	for property := range Object.Properties {

		if Object.Properties[property].Enumerable == true {
			keys = append(keys, property)
		}

	}

	//Returning Keys
	return keys

}

type Function_Properties = struct {
	Name       string
	Enumerable bool
	Value      func()
}

func (this *New) DefineProperties_Function(Object New, Properties ...Function_Properties) {

	for _, Property := range Properties {
		Object.Properties[Property.Name] = Object_Property[func()]{
			Enumerable: Property.Enumerable,
			Value:      Property.Value,
		}
	}
}

type Default_Properties = struct {
	Name       string
	Enumerable bool
	Value      any
}

func (this *New) DefineProperties_Default(Object New, Properties ...Default_Properties) {

	for _, Property := range Properties {
		Object.Values[Property.Name] = Object_Property[any]{
			Enumerable: Property.Enumerable,
			Value:      Property.Value,
		}
	}
}

type Entries_Return = map[string][]any

func (this *New) Entries(Object New) Entries_Return {

	var Entries Entries_Return = make(map[string][]any, 0)

	for k, v := range Object.Values {

		Entries[k] = []any{k, v.Value}
	}

	for property, v := range Object.Properties {

		if Object.Properties[property].Enumerable == true {
			Entries[property] = []any{property, v.Value}
		}
	}

	return Entries

}
