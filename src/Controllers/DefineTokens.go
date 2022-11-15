package syntax

import (
	"strings"
)

type Token struct {
	Type  string
	Value string
	Line  int32
}

func DefineToken(iterator int32, token string, lineNumber int32) []Token {

	if strings.Contains(token, "{") || strings.Contains(token, "}") || strings.Contains(token, "(") || strings.Contains(token, ")") || strings.Contains(token, "=") || strings.Contains(token, ">") || strings.Contains(token, ".") || strings.Contains(token, "/") || strings.Contains(token, `"`) || strings.Contains(token, ",") || strings.Contains(token, "`") && len(token) > 1 {

		Tokens := []Token{}

		for _, Value := range SplitSymbols(token) {

			if len(strings.ReplaceAll(Value, " ", "")) > 0 {
				Tokens = append(Tokens, Define(Value, lineNumber))
			}

		}

		return Tokens
	}

	return []Token{Define(token, lineNumber)}

}

func Define(token string, lineNumber int32) Token {

	if _IsSymbol(token) {
		return Token{
			Type:  "Punctuator",
			Value: token,
			Line:  lineNumber,
		}
	}

	return Token{
		Type:  "Indetifier",
		Value: token,
		Line:  lineNumber,
	}

}

func SplitSymbols(token string) []string {
	r := strings.NewReplacer("{", " { ", "=", " = ", "}", " } ", "(", " ( ", ")", " ) ", ";", " ; ", ".", " . ", ",", " , ", `"`, ` " `, "\n", "", "`", " ` ", "$", " $ ", "&", " & ", "?", " ? ", ":", " : ", "<", " < ", ">", " > ", "*", " * ", "/", " / ", "+", " + ", "%", " % ", "[", " [ ", "]", " ] ")
	return strings.Split(r.Replace(token), " ")
}
