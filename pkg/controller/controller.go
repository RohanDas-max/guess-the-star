package controller

import (
	"context"

	"github.com/rohandas-max/GuessTheStar/pkg/handler"
)

func Controller(ctx context.Context, flags map[string]string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		if err := handler.StartTheGame(flags); err != nil {
			return err
		}
	}
	return nil
}
