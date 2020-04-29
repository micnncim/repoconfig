package http

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/google/go-cmp/cmp"
	"go.uber.org/zap"
)

func TestClient_DoRequest(t *testing.T) {
	type args struct {
		ctx    context.Context
		method string
		path   string
		body   interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			}))

			endpoint, err := url.Parse(srv.URL)
			if err != nil {
				t.Fatalf("failed to parse url: %v", err)
			}
			c := &Client{
				endpoint: endpoint,
				logger:   zap.NewNop(),
			}

			got, err := c.DoRequest(context.Background(), tt.args.method, tt.args.path, map[string]string{}, tt.args.body)
			if err != nil {
				t.Fatalf("err: %v", err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("(-want +got):\n%s", diff)
			}
		})
	}
}
