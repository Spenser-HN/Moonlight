package syntax

import (
	"strconv"
)

func _IsNumber(str string) bool {

	_, err := strconv.ParseFloat(str, 64)

	if err == nil {
		return true
	}

	return false
}

/*




func CreateNumberToken (_Tokens []Token) []Token {

	var _IsString bool = false
	var String Token
	var StringValue string
	var _IsNumber bool = false
	var Number Token
	var IncovertedNumber string
	var line int32
	var VerifiedTokens []Token

	for _, v := range _Tokens {

		if v.Value == `"` && _IsString == false {
			_IsString = true
			line = v.Line

		} else if v.Value != `"` && _IsString == true {
			StringValue += " " + v.Value

		} else if v.Value == `"` && _IsString == true {
			_IsString = false
			String = VerifiedToken{Type: "String", Value: StringValue, Line: line}
			VerifiedTokens = append(VerifiedTokens, String)

		} else {
			Res, _ := IsNumber(v.Value)

			if Res == true && _IsNumber == false {
				_IsNumber = true
				line = v.Line
				IncovertedNumber += v.Value

			} else if Res == true && _IsNumber == true {
				IncovertedNumber += v.Value

			} else if Res == false && _IsNumber == true && v.Value == "." {
				IncovertedNumber += v.Value

			} else if Res == false && _IsNumber == true && v.Value != "." {
				_IsNumber = false
				number, _ := strconv.ParseFloat(IncovertedNumber, 64)
				IncovertedNumber = ""
				Number = VerifiedToken{Type: "Numeric", Value: number, Line: line}
				VerifiedTokens = append(VerifiedTokens, Number)

			} else {
				_default := VerifiedToken{
					Type:  v.Type,
					Value: v.Value,
					Line:  v.Line,
				}

				VerifiedTokens = append(VerifiedTokens, _default)
			}

		}
	}

	return VerifiedTokens

}
**/
