package controller

import (
	"context"
	"errors"
	"fmt"

	"github.com/rohandas-max/GuessTheStar/pkg/handler"
	"github.com/rohandas-max/GuessTheStar/pkg/utils"
)

func Controller(ctx context.Context, flags map[string]string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		repos, err := handler.GetTrendingRepos(flags)
		if err != nil {
			return err
		}
		if len(repos) > 4 {
			count := GettingTotalScore(repos)
			if count > 3 {
				fmt.Println("LETS GO!!! you won the game!")
			} else {
				fmt.Println("Opps!! Better Luck next Time")
			}
		} else {
			return errors.New("not enough trending repos to play the game")
		}
	}
	return nil
}

func GettingTotalScore(repos []handler.TrendingRepo) int {
	var count int
	for i := 0; i < 5; i++ {
		repo := repos[i]
		fmt.Printf("Round:%d\tRepo:%s/%s\nlanguage:%s\n", i+1, repo.Username, repo.Name, repo.Language)
		if checkScoreFromUserInput(repo.TotalStars) > 0 {
			count++
		}
	}
	return count
}

func checkScoreFromUserInput(TotalStars int) int {
	var input int
	fmt.Println("GUESS THE STAR")
	fmt.Scanf("%d", &input)
	if utils.IsCorrectGuess(TotalStars, input) {
		return 1
	}
	return 0
}
