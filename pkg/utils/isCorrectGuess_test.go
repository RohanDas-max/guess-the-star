package utils

import "testing"

func TestIsCorrectGuess(t *testing.T) {
	tests := []struct {
		name  string
		star  int
		input int
		want  bool
	}{
		{
			name:  "checking tolerance with exact input",
			star:  100,
			input: 100,
			want:  true,
		},
		{
			name:  "checking tolerance with 10% lower input",
			star:  100,
			input: 90,
			want:  true,
		},
		{
			name:  "checking tolerance with 10% higher input",
			star:  100,
			input: 110,
			want:  true,
		},
		{
			name:  "failed with more then tolerance",
			star:  100,
			input: 111,
			want:  false,
		},
		{
			name:  "checking tolerance with less then tolerance",
			star:  100,
			input: 89,
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
