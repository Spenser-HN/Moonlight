package Number

import (
	"strconv"
)

type New struct {
	Value float64
}

type NumberInterface interface {
	ToString(Number New) string
}

func (This *New) ToString(Number New) string {
	return strconv.FormatFloat(float64(Number.Value), 'f', -1, 64)
}
