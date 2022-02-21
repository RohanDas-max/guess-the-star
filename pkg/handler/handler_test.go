package handler

import (
	"testing"
)

func TestGetTrendingRepos(t *testing.T) {
	type args struct {
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
				flags: map[string]string{"lang": "go", "since": "weekly"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, got1 := GetTrendingRepos(tt.args.flags); (got1 != nil) != tt.wantErr {
				t.Errorf("GetTrendingRepos() = %v, want %v", got1, tt.wantErr)
			}
		})
	}
}
