package syntax

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
