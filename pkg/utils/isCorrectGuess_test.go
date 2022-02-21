package utils

import "testing"

func TestCheckTolerance(t *testing.T) {
	tests := []struct {
		name  string
		star  int
		input int
		want  bool
	}{
		{
			name:  "pass",
			star:  100,
			input: 100,
			want:  true,
		},
		{
			name:  "pass2",
			star:  100,
			input: 90,
			want:  true,
		},
		{
			name:  "pass4",
			star:  100,
			input: 110,
			want:  true,
		},
		{
			name:  "fail",
			star:  100,
			input: 111,
			want:  false,
		},
		{
			name:  "fail2",
			star:  100,
			input: 89,
			want:  false,
		},
		{
			name:  "fail3",
			star:  100,
			input: 79,
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsCorrectGuess(tt.star, tt.input)
			if got != tt.want {
				t.Errorf("got=%v, instead of %v", got, tt.want)
			}
		})
	}
}
