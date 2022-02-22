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
		var count int
		if len(repos) > 4 {
			for i := 0; i < 5; i++ {
				repo := repos[i]
				fmt.Printf("Round:%d\tRepo:%s/%s\nlanguage:%s\n", i+1, repo.Username, repo.Name, repo.Language)
				var input int
				fmt.Println("GUESS THE STAR")
				fmt.Scanf("%d", &input)
				if score := checkScore(repo.TotalStars, input); score > 0 {
					count++
				}
			}
			if count > 3 {
				fmt.Println("LETS GO!!! you won the game!")
			} else {
				fmt.Println("Opps!! Better Luck next Time")
			}
		} else {
			return errors.New("not enough trending repos")
		}
	}
	return nil
}

func checkScore(TotalStars, input int) int {
	if utils.IsCorrectGuess(TotalStars, input) {
		return 1
	}
	return 0
}
