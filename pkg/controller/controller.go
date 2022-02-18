package controller

import (
	"context"

	"github.com/rohandas-max/GuessTheStar/pkg/handler"
)

func Controller(ctx context.Context, url string) (int, error) {
	var n int
	select {
	case <-ctx.Done():
		return n, ctx.Err()
	default:
		score, err := handler.Handler(ctx, url)
		if err != nil {
			return n, err
		} else {
			return score, nil
		}
	}
}
