package main

import (
	"context"
	"flag"
	"time"

	"github.com/rohandas-max/GuessTheStar/pkg/controller"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	lang := flag.String("l", "", "enter the programming language")
	flag.Parse()
	var since = "weekly"
	var url string = "https://gh-trending-api.herokuapp.com/repositories/" + *lang + "?since=" + since
	controller.Controller(ctx, url)
}
