package app

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-github/v31/github"

	pkggithub "github.com/micnncim/repoconfig/pkg/github"
	"github.com/micnncim/repoconfig/pkg/survey"
)

func Test_askUpdateRepositoryInput(t *testing.T) {
	fakeRepo := &github.Repository{
		Name:             github.String("fake-repo"),
		Description:      github.String("fake description"),
		Private:          github.Bool(true),
		AllowSquashMerge: github.Bool(false),
	}

	tests := []struct {
		name                       string
		fakeAskInputMessages       map[string]string
		fakeAskSelectMessages      map[string]string
		fakeAskMultiSelectMessages map[string][]string
		want                       *pkggithub.UpdateRepositoryInput
		wantErr                    bool
	}{
		{
			name: "get updated input",
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
			want: &pkggithub.UpdateRepositoryInput{
				Name:             "fake-repo",
				Description:      "new description",
				Private:          false,
				AllowSquashMerge: true,
			},
			wantErr: false,
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
			if err != nil {
				t.Errorf("err: %v", err)
				return

			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Fatalf("(-want +got):\n%s", diff)
			}
		})
	}
}
