package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/rohandas-max/GuessTheStar/pkg/controller"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	lang := flag.String("l", "", "enter the programming language")
	since := flag.String("s", "weekly", "")
	flag.Parse()

	var url string = "https://gh-trending-api.herokuapp.com/repositories/" + *lang + "?since=" + *since
	score, err := controller.Controller(ctx, url)
	if err != nil {
		log.Fatal(err)
	} else if score > 3 {
		fmt.Println("LETS GO!!! you won the game!")
	} else {
		fmt.Println("Opps!! Better Luck next Time")
	}
}
