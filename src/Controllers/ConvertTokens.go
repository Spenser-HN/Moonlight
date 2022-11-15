package syntax

import (
	"fmt"
	"strings"
)

func ConvertTotoken(Line string, line_Number int32) {

	for iterator, token := range strings.Fields(Line) {
		fmt.Println(DefineToken(int32(iterator), token, line_Number))
	}

}
