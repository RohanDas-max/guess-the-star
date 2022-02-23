package controller

import (
	"context"
	"testing"
	"time"
)

func TestController(t *testing.T) {
	ctxWithdeadline, cancel := context.WithTimeout(context.Background(), 1*time.Microsecond)
	defer cancel()
	ctxWIthcancel, cancel := context.WithCancel(context.Background())
	defer cancel()
	tests := []struct {
		name    string
		ctx     context.Context
		flags   map[string]string
		wantErr bool
	}{
		{
			name:    "return error ctx deadline",
			ctx:     ctxWithdeadline,
			flags:   map[string]string{"lang": "rust", "since": "weekly"},
			wantErr: true,
		},
		{
			name:    "return nil error ctx cancel",
			ctx:     ctxWIthcancel,
			flags:   map[string]string{"lang": "rust", "since": "weekly"},
			wantErr: false,
		},
		{
			name:    "return error because of wrong flag",
			ctx:     ctxWIthcancel,
			flags:   map[string]string{"": "", "since": ""},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Controller(tt.ctx, tt.flags); (err != nil) != tt.wantErr {
				t.Errorf("Controller() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
