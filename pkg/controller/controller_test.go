package controller

import (
	"context"
	"testing"
	"time"
)

func TestController(t *testing.T) {
	ctx, cncl := context.WithTimeout(context.Background(), 1*time.Second)
	defer cncl()
	ctx1, cncl1 := context.WithCancel(context.Background())
	defer cncl1()

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
				ctx:   ctx,
				flags: map[string]string{"lang": "go", "since": "weekly"},
			},
			wantErr: true,
		},
		{
			name: "pass",
			args: args{
				ctx:   ctx1,
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
