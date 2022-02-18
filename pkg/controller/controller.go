package controller

import (
	"context"

	"github.com/rohandas-max/GuessTheStar/pkg/handler"
)

func Controller(ctx context.Context, url string) (int, error) {
	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	default:
		score, err := handler.Handler(ctx, url)
		if err != nil {
			return 0, err
		}
		return score, nil
	}
}
