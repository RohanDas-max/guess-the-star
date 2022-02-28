package utils

func IsCorrectGuess(Totalstar, input int) bool {
	tolerance := 10
	t := ((Totalstar * tolerance) / 100)
	if input <= Totalstar+t && input >= Totalstar-t {
		return true
	}
	return false
}
