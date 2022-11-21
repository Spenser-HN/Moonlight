package syntax

import (
	"strings"
)

type Converted map[int64]any

func ConvertTotoken(Line string, line_Number int32) Converted {

	var column int = 1
	var Tokens Converted = Converted{}

	for iterator := 0; iterator < len(strings.Fields(Line)); iterator++ {

		token := string(strings.Fields(Line)[iterator])

		startColumn := column
		endColumn := (column + len(token))

		if strings.Contains(token, "`") == true || strings.Contains(token, `"`) == true {

			Tokens[int64(len(Tokens))] = DefineString(line_Number, int32(startColumn), int32(endColumn), token)

		} else {
			Tokens[int64(len(Tokens))] = DefineToken(token, line_Number, int32(startColumn), int32(endColumn))
		}

		column += (len(token) + 1)

	}

	return Tokens

}
