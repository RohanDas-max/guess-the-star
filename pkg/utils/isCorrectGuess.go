package utils

func IsCorrectGuess(Totalstar, input int) bool {
	t := ((Totalstar * 10) / 100)
	if input <= Totalstar+t && input >= Totalstar-t {
		return true
	}
	return false
}
