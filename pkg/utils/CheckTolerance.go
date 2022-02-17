package utils

func CheckTolerance(star, input int) bool {
	t := ((star * 10) / 100)
	if input <= star+t && input >= star-t {
		return true
	}
	return false
}
