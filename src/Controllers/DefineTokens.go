package syntax

import (
	"strconv"
	"strings"
)

type Location struct {
	Line   int32
	Column int32
}

type Loc struct {
	Start Location
	End   Location
}

type Token struct {
	Type  string
	Value interface{}
	Loc   Loc
}

func IsNumber(str string) (bool, int) {
	val, err := strconv.Atoi(str)

	if err != nil {
		return false, 0
	}

	return true, val

}

func DefineToken(token string, lineNumber, startColumn, endColumn int32) []Token {

	if _IncludeSymbols(token) {

		Tokens := []Token{}

		for _, Value := range SplitSymbols(token) {

			if len(strings.ReplaceAll(Value, " ", "")) > 0 {
				Tokens = append(Tokens, Define(Value, lineNumber, startColumn, endColumn))
			}

		}

		return Tokens
	}

	return ([]Token{Define(token, lineNumber, startColumn, endColumn)})

}

func Define(token string, lineNumber, startColumn, endColumn int32) Token {

	if len(strings.ReplaceAll(token, " ", "")) == 0 {
		return Token{
			Type:  "Empty",
			Value: token,
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
	}

	if _IsSymbol(token) {
		return Token{
			Type:  "Punctuator",
			Value: token,
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
	}

	if _IsKeyword(token) {
		return Token{
			Type:  "Keyword",
			Value: token,
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
	}

	if _IsNumber(token) {
		return Token{
			Type:  "Numeric",
			Value: token,
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

	}

	return Token{
		Type:  "Indetifier",
		Value: token,
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

}

func SplitSymbols(token string) []string {
	r := strings.NewReplacer("{", " { ", "=", " = ", "}", " } ", "(", " ( ", ")", " ) ", ";", " ; ", ".", " . ", ",", " , ", `"`, ` " `, "\n", "", "`", " ` ", "$", " $ ", "&", " & ", "?", " ? ", ":", " : ", "<", " < ", ">", " > ", "*", " * ", "/", " / ", "+", " + ", "%", " % ", "[", " [ ", "]", " ] ")
	return strings.Split(r.Replace(token), " ")
}
