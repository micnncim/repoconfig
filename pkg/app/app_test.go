package app

import (
	"reflect"
	"testing"
)

func Test_parseArgs(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		args    args
		want    []repository
		wantErr bool
	}{
		{
			name: "only owner",
			args: args{
				args: []string{"micnncim"},
			},
			want:    []repository{{owner: "micnncim"}},
			wantErr: false,
		},
		{
			name: "set of owner and repo",
			args: args{
				args: []string{"micnncim/repoconfig"},
			},
			want:    []repository{{owner: "micnncim", repo: "repoconfig"}},
			wantErr: false,
		},
		{
			name: "multiple set of owner and repo",
			args: args{
				args: []string{"micnncim/repoconfig", "micnncim/label-exporter"},
			},
			want: []repository{
				{owner: "micnncim", repo: "repoconfig"},
				{owner: "micnncim", repo: "label-exporter"},
			},
			wantErr: false,
		},
		{
			name: "mix of only owner and set of owner and repo",
			args: args{
				args: []string{"actions-ecosystem", "micnncim/repoconfig"},
			},
			want: []repository{
				{owner: "actions-ecosystem"},
				{owner: "micnncim", repo: "repoconfig"},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := parseArgs(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
