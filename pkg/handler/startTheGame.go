package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/rohandas-max/GuessTheStar/pkg/utils"
)

//startthegame will return error
//1.gettherepos
//2.print repo
//3.get the user input userguessstar
//4.iscorrectguess
//5.incremment the score
//6.isWinner

type TrendingRepo struct {
	Username   string `json:"username"`
	Name       string `json:"repositoryName"`
	TotalStars int    `json:"totalStars"`
	Language   string `json:"language"`
}

func StartTheGame(flags map[string]string) error {
	flagLanguage := flags["lang"]
	flagTrendSince := flags["since"]
	repos, _ := getTheRepos(flagLanguage, flagTrendSince)
	if len(repos) > 4 {
		var totalScore int
		for i := 0; i < 5; i++ {
			repo := repos[i]
			printsRoundAndRepos(i+1, repo)
			input := userGuessStar()
			totalScore += score(repo.TotalStars, input)
		}
		printResult(totalScore)
	} else {
		return errors.New("not enough repos to play the game")
	}
	return nil
}

func getTheRepos(lang, since string) ([]TrendingRepo, error) {
	var url string = "https://gh-trending-api.herokuapp.com/repositories/" + lang + "?since=" + since
	httpResponse, _ := getHttpResponse(url)
	defer httpResponse.Body.Close()
	var r []TrendingRepo
	if httpResponse.StatusCode == http.StatusOK {
		json.NewDecoder(httpResponse.Body).Decode(&r)
	} else {
		return nil, errors.New(strconv.FormatInt(http.StatusNotFound, 10))
	}
	return r, nil
}
func getHttpResponse(url string) (*http.Response, error) {
	httpResponse, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	return httpResponse, nil
}
func printsRoundAndRepos(round int, repo TrendingRepo) {
	fmt.Printf("Round:%d\tRepo:%s/%s\nlanguage:%s\n", round, repo.Username, repo.Name, repo.Language)
}

func userGuessStar() int {
	var input int
	fmt.Println("GUESS THE STAR")
	fmt.Scan(&input)
	return input
}

func score(stars, input int) int {
	if utils.IsCorrectGuess(stars, input) {
		return 1
	}
	return 0
}

func printResult(score int) {
	if score > 3 {
		fmt.Println("🔥 you won the game")
	} else {
		fmt.Println("😥 try Again!")
	}
}
