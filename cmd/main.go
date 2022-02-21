package main

import (
	"context"
	"flag"
	"log"

	"github.com/rohandas-max/GuessTheStar/pkg/controller"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	lang := flag.String("l", "", "enter the programming language")
	since := flag.String("s", "weekly", "eg: daily/weekly/monthly")
	flag.Parse()
	flags := make(map[string]string)
	flags["lang"] = *lang
	flags["since"] = *since
	err := controller.Controller(ctx, flags)
	if err != nil {
		log.Fatal(err)
	}
}
