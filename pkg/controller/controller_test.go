package controller

import (
	"context"
	"testing"
	"time"

	"github.com/rohandas-max/GuessTheStar/pkg/handler"
)

func TestController(t *testing.T) {
	ctxWithTimeout, cancle := context.WithTimeout(context.Background(), 1*time.Microsecond)
	defer cancle()
	ctxWithCancel, cancle := context.WithCancel(context.Background())
	defer cancle()

	type args struct {
		ctx   context.Context
		flags map[string]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "fail with context deadline",
			args: args{
				ctx:   ctxWithTimeout,
				flags: map[string]string{"lang": "go", "since": "weekly"},
			},
			wantErr: true,
		},
		{
			name: "passed with nil error",
			args: args{
				ctx:   ctxWithCancel,
				flags: map[string]string{"lang": "go", "since": "weekly"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Controller(tt.args.ctx, tt.args.flags); (err != nil) != tt.wantErr {
				t.Errorf("Controller() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_checkUserInputAndScore(t *testing.T) {
	tests := []struct {
		name       string
		totalStars int
		input      int
		want       int
	}{
		{
			name:       "should return 0 because of wrong input",
			totalStars: 123,
			input:      1,
			want:       0,
		},
		{
			name:       "should return 1 because of right input",
			totalStars: 123,
			input:      123,
			want:       1,
		},
		{
			name:       "should return 0 because of empty value",
			totalStars: 0,
			input:      0,
			want:       1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkScoreFromUserInput(tt.totalStars); got != tt.want {
				t.Errorf("checkUserInputAndScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_GettingTotalScore(t *testing.T) {
	type args struct {
		repos []handler.TrendingRepo
		want  int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "passed with empty values",
			args: args{
				repos: []handler.TrendingRepo{{
					Username:   "",
					Name:       "",
					TotalStars: 0,
					Language:   "",
				}, {}, {}, {}, {}},
				want: 5,
			},
		},
		{
			name: "passed with empty values",
			args: args{
				repos: []handler.TrendingRepo{{
					Username:   "",
					Name:       "",
					TotalStars: 0,
					Language:   "",
				}, {}, {}, {}, {}},
				want: 5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GettingTotalScore(tt.args.repos)
			if got != tt.args.want {
				t.Error("got != want", got, tt.args.want)
			}
		})
	}
}
