package Boolean

func ValueOf(Data bool) bool {
	return Data
}

func ToString(Data bool) string {
	if Data == true {
		return "true"
	} else {
		return "false"
	}
}
