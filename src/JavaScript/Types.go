package JavaScript

//null
type Null interface{ error }

//undefined
type Undefined interface{ error }

//number
type TypeNumber interface {
	~int | float64 | complex128
}

//unknown
type Unknown interface{}

//setTimeout
type TimeoutCallback func()

//setInterval
type TimeIntervalCallback func()

//Headers
type Headers struct {
	Content_Type   string
	Content_Length int
}

//Fetch
type Request[T any] struct {
	Method      string
	Body        T
	Credentials string
	Headers     Headers
}
