package app

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-github/v31/github"

	pkggithub "github.com/micnncim/repoconfig/pkg/github"
	"github.com/micnncim/repoconfig/pkg/survey"
	"github.com/spf13/cobra"
)

func Test_app_run(t *testing.T) {
	var (
		fakeOwner = "fake-owner"
		fakeRepo  = "fake-repo"
	)

	tests := []struct {
		name                         string
		fakeRepo                     *pkggithub.FakeRepository
		fakeAskUpdateRepositoryInput func(survey.Surveyor, *github.Repository) (*pkggithub.UpdateRepositoryInput, error)
		want                         *pkggithub.FakeRepository
		wantErr                      bool
	}{
		{
			name: "update repo",
			fakeRepo: &pkggithub.FakeRepository{
				Description: "fake description",
				Private:     false,
			},
			fakeAskUpdateRepositoryInput: func(survey.Surveyor, *github.Repository) (*pkggithub.UpdateRepositoryInput, error) {
				return &pkggithub.UpdateRepositoryInput{
					Description: "new description",
					Private:     true,
				}, nil
			},
			want: &pkggithub.FakeRepository{
				Description: "new description",
				Private:     true,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			func() {
				orgAskUpdateRepositoryInput := askUpdateRepositoryInput
				askUpdateRepositoryInput = tt.fakeAskUpdateRepositoryInput
				defer func() {
					askUpdateRepositoryInput = orgAskUpdateRepositoryInput
				}()

				githubClient := &pkggithub.FakeClient{
					Repos: map[string]*pkggithub.FakeRepository{
						fakeRepo: tt.fakeRepo,
					},
				}

				a := &app{
					githubClient: githubClient,
				}

				if err := a.run(&cobra.Command{}, []string{fakeOwner, fakeRepo}); (err != nil) != tt.wantErr {
					t.Errorf("err = %v, wantErr %v", err, tt.wantErr)
				}

				got := githubClient.Repos[fakeRepo]
				if diff := cmp.Diff(tt.want, got); diff != "" {
					t.Fatalf("(-want +got):\n%s", diff)
				}
			}()
		})
	}
}
