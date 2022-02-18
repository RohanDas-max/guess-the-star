package controller

import (
	"context"
	"testing"
)

func TestController(t *testing.T) {
	ctx, cncle := context.WithCancel(context.Background())
	defer cncle()

	type args struct {
		ctx context.Context
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "without url",
			args: args{
				ctx: ctx,
				url: "",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "with url",
			args: args{
				ctx: ctx,
				url: "http://www.google.com",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "with url",
			args: args{
				ctx: ctx,
				url: "https://gh-trending-api.herokuapp.com/repositories",
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Controller(tt.args.ctx, tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("Controller() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Controller() = %v, want %v", got, tt.want)
			}
		})
	}
}
