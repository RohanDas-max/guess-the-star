package handler

import (
	"context"
	"encoding/json"

	"github.com/rohandas-max/GuessTheStar/pkg/utils"
)

type TrendingRepo struct {
	Username   string `json:"username"`
	Name       string `json:"repositoryName"`
	TotalStars int    `json:"totalStars"`
	Language   string `json:"language"`
}

func GetTrendingRepos(ctx context.Context, flags map[string]string) ([]TrendingRepo, error) {
	lang := flags["lang"]
	since := flags["since"]
	var url string = "https://gh-trending-api.herokuapp.com/repositories/" + lang + "?since=" + since
	bytes, err := utils.Get(ctx, url)
	if err != nil {
		return []TrendingRepo{}, err
	}
	var repos []TrendingRepo
	if err := json.Unmarshal(bytes, &repos); err != nil {
		return []TrendingRepo{}, err
	}
	return repos, nil
}
