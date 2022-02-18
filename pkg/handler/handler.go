package handler

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/rohandas-max/GuessTheStar/pkg/utils"
)

type response struct {
	Username string `json:"username"`
	Repo     string `json:"repositoryName"`
	Star     int    `json:"totalStars"`
	Language string `json:"language"`
}

func Handler(ctx context.Context, url string) (int, error) {
	bytes, err := utils.Get(ctx, url)
	if err != nil {
		return 0, err
	}
	var res []response
	if err := json.Unmarshal(bytes, &res); err != nil {
		return 0, err
	}
	return checkStar(res), nil
}

func checkStar(res []response) int {
	var s int
	for i, r := range res {
		if i < 5 {
			var input int
			fmt.Printf("Round:=%d\tRepo:%s/%s\n", i+1, r.Username, r.Repo)
			fmt.Println("GUESS THE STAR")
			fmt.Scan(&input)
			s += score(i, input, r)
		} else {
			break
		}
	}
	return s
}
