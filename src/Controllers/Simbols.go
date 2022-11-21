package syntax

import "strings"

func _IsSymbol(value string) bool {
	switch value {
	case "<":
		return true
	case ">":
		return true
	case ":":
		return true
	case ";":
		return true
	case ".":
		return true
	case "!":
		return true
	case "`":
		return true
	case `"`:
		return true
	case "+":
		return true
	case "-":
		return true
	case "=":
		return true
	case "/":
		return true
	case "?":
		return true
	case "^":
		return true
	case "%":
		return true
	case "*":
		return true
	case "#":
		return true
	case "[":
		return true
	case "]":
		return true
	case "&":
		return true
	case "~":
		return true
	case ",":
		return true
	case "|":
		return true
	case "(":
		return true
	case ")":
		return true
	case "{":
		return true
	case "}":
		return true
	default:
		return false
	}
}

func _IncludeSymbols(str string) bool {
	return (strings.Contains(str, "{") || strings.Contains(str, "}") ||
		strings.Contains(str, "(") || strings.Contains(str, ")") ||
		strings.Contains(str, "=") || strings.Contains(str, "<") ||
		strings.Contains(str, ">") || strings.Contains(str, ".") ||
		strings.Contains(str, "/") || strings.Contains(str, ",") ||
		strings.Contains(str, "[") || strings.Contains(str, "]") ||
		strings.Contains(str, "+") || strings.Contains(str, "-") ||
		strings.Contains(str, "*") || strings.Contains(str, "%") ||
		strings.Contains(str, "^") || strings.Contains(str, "!") ||
		strings.Contains(str, "#") || strings.Contains(str, "~") ||
		strings.Contains(str, "&") || strings.Contains(str, ":") ||
		strings.Contains(str, ";"))
}
