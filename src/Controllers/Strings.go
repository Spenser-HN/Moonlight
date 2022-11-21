package syntax

// ----------- Falta definir bien las strings ----------
// ----------- Solo las strings bien definidas [ex:"Home"] son aceptadas por ahora -----

import (
	"log"
	"strings"
)

func _ValidString(Str string) bool {
	return (strings.Contains(Str, `"`) || strings.Contains(Str, "`") || strings.Contains(Str, "'")) &&
		(string(Str[len(Str)-1]) == `"` ||
			string(Str[len(Str)-1]) == "`" ||
			string(Str[len(Str)-1]) == "'") &&
		(string(Str[0]) == `"` ||
			string(Str[0]) == "`" ||
			string(Str[0]) == "'")
}

func CreateTokenTroughString(Str string, lineNumber, startColumn, endColumn int32, Tokens []Token) []Token {
	var _string string
	var InitSymbol string
	_init := false
	Done := false

	for i, Value := range SplitSymbols(Str) {

		// Founded "`" or `"`
		if Value == `"` || Value == "`" || Value == "'" {

			if _init == false {
				_init = true
				Done = false
				InitSymbol = Value
			} else if _init && (InitSymbol == Value) {
				_init = false
				Done = true

			}

		}

		// String value
		if Value != `"` && Value != "`" && Value != "'" && _init == true && !Done {
			_string += Value
		}

		//done

		//Missing "`" or `"` at end of statement
		if _init == true && i == len(Str) {
			log.Fatal(` Simbol (") waited at the end of the statement `, Str, _init, _string)
		}

		//All is Ok
		if _init == false && Value != `"` && Value != "`" && Value != "'" && Done &&
			len(strings.ReplaceAll(Value, " ", "")) > 0 && len(_string) > 0 {

			_Token := Token{
				Type:  "String",
				Value: _string,
				Loc: Loc{
					Start: Location{
						Line:   lineNumber,
						Column: startColumn,
					},
					End: Location{
						Line:   lineNumber,
						Column: int32(i),
					},
				},
			}

			Tokens = append(Tokens, _Token)

			_string = ""
		}

		if Done && _init == false && Value != `"` && Value != "`" && Value != "'" {

			result := Define(Value, lineNumber, startColumn, int32(i))

			if result.Type != "Empty" {
				Tokens = append(Tokens, result)
			}

		}

	}

	return Tokens

}

func DefineString(lineNumber, startColumn, endColumn int32, str ...string) []Token {
	var Tokens []Token

	for _, Str := range str {

		if string(Str[0]) != "'" && string(Str[0]) != "`" && string(Str[0]) != `"` &&
			_IncludeSymbols(Str) {

			return CreateTokenTroughString(Str, lineNumber, startColumn, endColumn, Tokens)

		}

		if _ValidString(Str) {

			//replace first: [`]
			replaced := strings.ReplaceAll(Str, "`", "")

			//Replace first: ["]
			replaced = strings.ReplaceAll(replaced, `"`, "")

			//Replace first: ["]
			replaced = strings.ReplaceAll(replaced, `'`, "")

			_Token := Token{
				Type:  "String",
				Value: replaced,
				Loc: Loc{
					Start: Location{
						Line:   lineNumber,
						Column: startColumn,
					},
					End: Location{
						Line:   lineNumber,
						Column: endColumn,
					},
				},
			}

			Tokens = append(Tokens, _Token)

		}

		if (string(Str[0]) == `"` || string(Str[0]) == "`" || string(Str[0]) == "'") && _IncludeSymbols(Str) {
			return CreateTokenTroughString(Str, lineNumber, startColumn, endColumn, Tokens)
		}

	}

	return Tokens
}
