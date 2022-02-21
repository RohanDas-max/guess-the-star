package handler

import (
	"context"
	"testing"
)

func TestGetTrendingRepos(t *testing.T) {
	ctx, cncl := context.WithCancel(context.Background())
	defer cncl()
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
			name: "with data",
			args: args{
				ctx:   ctx,
				flags: map[string]string{"lang": "go", "since": "weekly"},
			},
			wantErr: false,
		},
		{
			name: "pass2",
			args: args{
				ctx:   ctx,
				flags: map[string]string{"lang": "rust", "since": "weekly"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, got1 := GetTrendingRepos(tt.args.ctx, tt.args.flags); (got1 != nil) != tt.wantErr {
				t.Errorf("GetTrendingRepos() = %v, want %v", got1, tt.wantErr)
			}
		})
	}
}
