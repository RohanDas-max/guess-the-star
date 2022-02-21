package controller

import (
	"context"
	"testing"
	"time"
)

func TestController(t *testing.T) {
	ctxWithTimeout, cancle := context.WithTimeout(context.Background(), 1*time.Second)
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
			name: "context deadline",
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
