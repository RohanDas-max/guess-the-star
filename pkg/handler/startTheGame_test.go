package handler

import (
	"testing"
)

func TestStartTheGame(t *testing.T) {
	tests := []struct {
		name    string
		flags   map[string]string
		wantErr bool
	}{
		{
			name:    "passed with language go and since weekly",
			flags:   map[string]string{"lang": "go", "since": "weekly"},
			wantErr: false,
		},
		{
			name:    "return error with wrong language",
			flags:   map[string]string{"lang": "goasas", "since": "weekly"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := StartTheGame(tt.flags); (err != nil) != tt.wantErr {
				t.Errorf("StartTheGame() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getTheRepos(t *testing.T) {
	tests := []struct {
		name    string
		lang    string
		since   string
		wantErr bool
	}{
		{
			name:    "nil error with lang go and since weekly",
			lang:    "go",
			since:   "weekly",
			wantErr: false,
		},
		{
			name:    "returned error with wrong lang and since",
			lang:    "golang",
			since:   "yearly",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := getTheRepos(tt.lang, tt.since)
			if (err != nil) != tt.wantErr {
				t.Errorf("getTheRepos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func Test_printRoundAndRepo(t *testing.T) {
	tests := []struct {
		name string
		i    int
		repo TrendingRepo
	}{
		{
			name: "nil error return with right input",
			i:    0,
			repo: TrendingRepo{
				Username:   "random",
				Name:       "random",
				TotalStars: 12,
				Language:   "go",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printsRoundAndRepo(tt.i, tt.repo)
		})
	}
}

func Test_userGuessStar(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "passed",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := userGuessStar(); got != tt.want {
				t.Errorf("userGuessStar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isWinner(t *testing.T) {
	tests := []struct {
		name  string
		score int
	}{
		{
			name:  "passed with a input",
			score: 1212,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result(tt.score)
		})
	}
}

func Test_scoreIncrement(t *testing.T) {
	tests := []struct {
		name  string
		star  int
		input int
		want  int
	}{
		{
			name:  "returned score 1 with equal input & stars",
			star:  123,
			input: 123,
			want:  1,
		}, {
			name:  "return score 0 with wrong guess",
			star:  123,
			input: 2,
			want:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := score(tt.star, tt.input); got != tt.want {
				t.Errorf("scoreIncrement() = %v, want %v", got, tt.want)
			}
		})
	}
}
