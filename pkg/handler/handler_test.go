package handler

import (
	"context"
	"testing"
)

func TestHandler(t *testing.T) {
	ctx, cncl := context.WithCancel(context.Background())
	defer cncl()
	tests := []struct {
		name    string
		ctx     context.Context
		url     string
		want    int
		wantErr bool
	}{
		{
			name:    "with url",
			ctx:     ctx,
			url:     "http://google.com",
			want:    0,
			wantErr: false,
		},
		{
			name:    "without url",
			ctx:     ctx,
			url:     "",
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			got, err := Handler(tt.ctx, tt.url)
			if got != tt.want && (err == nil) != tt.wantErr {
				t.Errorf("got=%v and error=%v, instead of %v,%v", got, err, tt.want, tt.wantErr)
			}
		})
	}
}

func TestCheckStar(t *testing.T) {
	tests := []struct {
		name string
		res  []response
		want int
	}{
		{
			name: "nil",
			res:  []response{},
			want: 0,
		},
		{
			name: "not nil",
			res: []response{
				{
					Username: "random",
					Repo:     "random",
					Star:     123,
					Language: "",
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkStar(tt.res); got != tt.want {
				t.Errorf("got=%v, want %v", got, tt.want)
			}
		})
	}
}
