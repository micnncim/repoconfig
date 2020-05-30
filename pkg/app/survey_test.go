package app

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/micnncim/repoconfig/pkg/github"
	"github.com/micnncim/repoconfig/pkg/survey"
)

func Test_askUpdateRepositoryInput(t *testing.T) {
	fakeRepo := &github.Repository{
		Name:             "fake-repo",
		Description:      "fake description",
		Private:          true,
		AllowSquashMerge: false,
	}

	tests := []struct {
		name                       string
		fakeAskInputMessages       map[string]string
		fakeAskSelectMessages      map[string]string
		fakeAskMultiSelectMessages map[string][]string
		want                       *github.Repository
		wantErr                    bool
	}{
		{
			name: "get updated repo",
			fakeAskInputMessages: map[string]string{
				surveyKeyDescription: "new description",
			},
			fakeAskSelectMessages: map[string]string{
				surveyKeyPrivate:          "false",
				surveyKeyAllowSquashMerge: "true",
			},
			fakeAskMultiSelectMessages: map[string][]string{
				askUpdateRepositoryInputMessage: {surveyKeyDescription, surveyKeyPrivate, surveyKeyAllowSquashMerge},
			},
			want: &github.Repository{
				Name:             "fake-repo",
				Description:      "new description",
				Private:          false,
				AllowSquashMerge: true,
			},
			wantErr: false,
		},
		{
			name:                       "not get updated repo",
			fakeAskInputMessages:       map[string]string{},
			fakeAskSelectMessages:      map[string]string{},
			fakeAskMultiSelectMessages: map[string][]string{},
			want:                       nil,
			wantErr:                    true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := &survey.FakeSurveyor{
				AskInputMessages:       tt.fakeAskInputMessages,
				AskSelectMessages:      tt.fakeAskSelectMessages,
				AskMultiSelectMessages: tt.fakeAskMultiSelectMessages,
			}

			got, err := askUpdateRepositoryInput(s, fakeRepo)
			if (err != nil) != tt.wantErr {
				t.Errorf("err = %v, wantErr %v", err, tt.wantErr)
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("(-want +got):\n%s", diff)
			}
		})
	}
}
