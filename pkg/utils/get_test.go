package utils

import (
	"context"
	"net/http"
	"testing"
)

func TestGet(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	tests := []struct {
		name    string
		ctx     context.Context
		url     string
		wantErr bool
	}{
		{
			name:    "test1",
			url:     "http://google.com",
			ctx:     ctx,
			wantErr: false,
		},
		{
			name:    "test2 empty url",
			url:     "",
			ctx:     ctx,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Get(tt.ctx, tt.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("get()=%v, wantErr=%v", err, tt.wantErr)
			}
		})
	}

}

func TestCheckStatus(t *testing.T) {
	resp, _ := http.Get("http://google.com")

	tests := []struct {
		name    string
		h       *http.Response
		wantErr bool
	}{
		{
			name:    "pass with http response",
			h:       resp,
			wantErr: false,
		},
		{
			name:    "fail with empty http response",
			h:       &http.Response{},
			wantErr: true,
		},
	}
	defer resp.Body.Close()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := checkStatus(tt.h)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}

}
