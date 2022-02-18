package handler

import (
	"github.com/rohandas-max/GuessTheStar/pkg/utils"
)

func score(round, input int, r response) int {
	var score int
	if utils.CheckTolerance(r.Star, input) {
		score++
	}
	round++
	return score
}
