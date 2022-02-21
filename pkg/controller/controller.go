package controller

import (
	"context"
	"fmt"

	"github.com/rohandas-max/GuessTheStar/pkg/handler"
	"github.com/rohandas-max/GuessTheStar/pkg/utils"
)

func Controller(ctx context.Context, flags map[string]string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		repos, err := handler.GetTrendingRepos(ctx, flags)
		if err != nil {
			return err
		}
		var count int
		for i := 0; i < 5; i++ {
			var input int
			fmt.Printf("Round:%d\tRepo:%s/%s\nlanguage:%s\n", i+1, repos[i].Username, repos[i].Name, repos[i].Language)
			fmt.Println("GUESS THE STAR")
			fmt.Scan(&input)
			if utils.IsCorrectGuess(repos[i].TotalStars, input) {
				count++
			}
		}
		if count > 3 {
			fmt.Println("LETS GO!!! you won the game!")
		} else {
			fmt.Println("Opps!! Better Luck next Time")
		}
	}
	return nil
}
