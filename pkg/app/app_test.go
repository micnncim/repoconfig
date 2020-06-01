package app

import (
	"context"
	"io/ioutil"
	"log"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/spf13/cobra"

	"github.com/micnncim/repoconfig/pkg/github"
	"github.com/micnncim/repoconfig/pkg/spinner"
	"github.com/micnncim/repoconfig/pkg/survey"
)

func Test_app_run(t *testing.T) {
	log.SetOutput(ioutil.Discard)

	fakeOwner := "fake-owner"

	orgAskUpdateRepositoryInput := askUpdateRepositoryInput
	defer func() {
		askUpdateRepositoryInput = orgAskUpdateRepositoryInput
	}()

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
				Name:             "fake-repo-1",
				Description:      "fake description",
				Private:          false,
				AllowMergeCommit: true,
			},
			fakeAskUpdateRepositoryInput: func(survey.Surveyor, *github.Repository) (*github.Repository, error) {
				return &github.Repository{
					Name:             "fake-repo-1",
					Description:      "new description",
					Private:          true,
					AllowMergeCommit: false,
				}, nil
			},
			want: &github.Repository{
				Name:             "fake-repo-1",
				Description:      "new description",
				Private:          true,
				AllowMergeCommit: false,
			},
			wantErr: false,
		},
		{
			name: "not update repo",
			fakeRepo: &github.Repository{
				Name:             "fake-repo-2",
				Description:      "fake description",
				Private:          false,
				AllowMergeCommit: true,
			},
			fakeAskUpdateRepositoryInput: func(survey.Surveyor, *github.Repository) (*github.Repository, error) {
				return nil, ErrRepositoryNoChange
			},
			want: &github.Repository{
				Name:             "fake-repo-2",
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
			askUpdateRepositoryInput = tt.fakeAskUpdateRepositoryInput

			fakeGithubClient := github.NewFakeClient()
			if tt.fakeRepo != nil {
				fakeGithubClient.SetRepository(tt.fakeRepo.Name, tt.fakeRepo)
			}

			a := &app{
				githubClient: fakeGithubClient,
				spinner:      spinner.New(spinner.WithWriter(ioutil.Discard)),
			}

			if err := a.run(&cobra.Command{}, []string{fakeOwner, tt.fakeRepo.Name}); (err != nil) != tt.wantErr {
				t.Errorf("err = %v, wantErr %v", err, tt.wantErr)
			}

			got, err := fakeGithubClient.GetRepository(context.Background(), fakeOwner, tt.fakeRepo.Name)
			if err != nil {
				t.Errorf("err: %v", err)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("(-want +got):\n%s", diff)
			}
		})
	}
}
