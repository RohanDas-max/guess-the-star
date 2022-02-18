package handler

import "testing"

func TestScore(t *testing.T) {
	tests := []struct {
		name  string
		round int
		input int
		r     response
		want  int
	}{
		{
			name:  "pass1",
			round: 1,
			input: 12,
			r: response{
				Username: "repo",
				Repo:     "repo",
				Star:     12,
				Language: "any",
			},
			want: 1,
		},
		{
			name:  "pass1",
			round: 1,
			input: 12,
			r: response{
				Username: "repo",
				Repo:     "repo",
				Star:     1,
				Language: "any",
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := score(tt.round, tt.input, tt.r); got != tt.want {
				t.Errorf("got%v, instead of %v", got, tt.want)
			}
		})
	}
}
