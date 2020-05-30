package app

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/micnncim/repoconfig/pkg/github"
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
		fakeRepo                     *github.Repository
		fakeAskUpdateRepositoryInput func(survey.Surveyor, *github.Repository) (*github.Repository, error)
		want                         *github.Repository
		wantErr                      bool
	}{
		{
			name: "update repo",
			fakeRepo: &github.Repository{
				Description:      "fake description",
				Private:          false,
				AllowMergeCommit: true,
			},
			fakeAskUpdateRepositoryInput: func(survey.Surveyor, *github.Repository) (*github.Repository, error) {
				return &github.Repository{
					Description:      "new description",
					Private:          true,
					AllowMergeCommit: false,
				}, nil
			},
			want: &github.Repository{
				Description:      "new description",
				Private:          true,
				AllowMergeCommit: false,
			},
			wantErr: false,
		},
		{
			name: "not update repo",
			fakeRepo: &github.Repository{
				Description:      "fake description",
				Private:          false,
				AllowMergeCommit: true,
			},
			fakeAskUpdateRepositoryInput: func(survey.Surveyor, *github.Repository) (*github.Repository, error) {
				return nil, ErrRepositoryNoChange
			},
			want: &github.Repository{
				Description:      "fake description",
				Private:          false,
				AllowMergeCommit: true,
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

				githubClient := &github.FakeClient{
					Repos: make(map[string]*github.Repository),
				}
				if tt.fakeRepo != nil {
					githubClient.Repos[fakeRepo] = tt.fakeRepo
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
