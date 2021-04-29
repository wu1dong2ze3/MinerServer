package cgminer

func StatusType(str string) int {
	switch str {
	case "Alive":
		return 0
	default:
		return 1
	}
}
