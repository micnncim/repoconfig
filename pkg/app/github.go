package app

import (
	"github.com/google/go-github/v31/github"

	pkggithub "github.com/micnncim/repoconfig/pkg/github"
)

const (
	surveyKeyName                = "Name"
	surveyKeyDescription         = "Description"
	surveyKeyHomepage            = "Homepage"
	surveyKeyPrivate             = "Private"
	surveyKeyVisibility          = "Visibility"
	surveyKeyHasIssues           = "HasIssues"
	surveyKeyHasProjects         = "HasProjects"
	surveyKeyHasWiki             = "HasWiki"
	surveyKeyDefaultBranch       = "DefaultBranch"
	surveyKeyAllowSquashMerge    = "AllowSquashMerge"
	surveyKeyAllowMergeCommit    = "AllowMergeCommit"
	surveyKeyAllowRebaseMerge    = "AllowRebaseMerge"
	surveyKeyDeleteBranchOnMerge = "DeleteBranchOnMerge"
	surveyKeyArchived            = "Archived"
)

var surveyUpdateRepositoryOptions = []string{
	surveyKeyName,
	surveyKeyDescription,
	surveyKeyHomepage,
	surveyKeyPrivate,
	surveyKeyVisibility,
	surveyKeyHasIssues,
	surveyKeyHasProjects,
	surveyKeyHasWiki,
	surveyKeyDefaultBranch,
	surveyKeyAllowSquashMerge,
	surveyKeyAllowMergeCommit,
	surveyKeyAllowRebaseMerge,
	surveyKeyDeleteBranchOnMerge,
	surveyKeyArchived,
}

type updateRepositoryOptions struct {
	Name                bool
	Description         bool
	Homepage            bool
	Private             bool
	Visibility          bool
	HasIssues           bool
	HasProjects         bool
	HasWiki             bool
	DefaultBranch       bool
	AllowSquashMerge    bool
	AllowMergeCommit    bool
	AllowRebaseMerge    bool
	DeleteBranchOnMerge bool
	Archived            bool
}

func askUpdateRepositoryInput(currentRepo *github.Repository) (*pkggithub.UpdateRepositoryInput, error) {
	resp, err := askMultiSelect("What would you like to update?", surveyUpdateRepositoryOptions)
	if err != nil {
		return nil, err
	}

	opts := &updateRepositoryOptions{
		Name:                contains(surveyKeyName, resp),
		Description:         contains(surveyKeyDescription, resp),
		Homepage:            contains(surveyKeyHomepage, resp),
		Private:             contains(surveyKeyPrivate, resp),
		Visibility:          contains(surveyKeyVisibility, resp),
		HasIssues:           contains(surveyKeyHasIssues, resp),
		HasProjects:         contains(surveyKeyHasProjects, resp),
		HasWiki:             contains(surveyKeyHasWiki, resp),
		DefaultBranch:       contains(surveyKeyDefaultBranch, resp),
		AllowSquashMerge:    contains(surveyKeyAllowSquashMerge, resp),
		AllowMergeCommit:    contains(surveyKeyAllowMergeCommit, resp),
		AllowRebaseMerge:    contains(surveyKeyAllowRebaseMerge, resp),
		DeleteBranchOnMerge: contains(surveyKeyDeleteBranchOnMerge, resp),
		Archived:            contains(surveyKeyArchived, resp),
	}

	input := &pkggithub.UpdateRepositoryInput{
		Name:                currentRepo.GetName(),
		Description:         currentRepo.GetDescription(),
		Homepage:            currentRepo.GetHomepage(),
		Private:             currentRepo.GetPrivate(),
		Visibility:          currentRepo.GetVisibility(),
		HasIssues:           currentRepo.GetHasIssues(),
		HasProjects:         currentRepo.GetHasProjects(),
		HasWiki:             currentRepo.GetHasWiki(),
		DefaultBranch:       currentRepo.GetDefaultBranch(),
		AllowSquashMerge:    currentRepo.GetAllowSquashMerge(),
		AllowMergeCommit:    currentRepo.GetAllowMergeCommit(),
		AllowRebaseMerge:    currentRepo.GetAllowRebaseMerge(),
		DeleteBranchOnMerge: currentRepo.GetDeleteBranchOnMerge(),
		Archived:            currentRepo.GetArchived(),
	}

	if opts.Name {
		input.Name, err = askInput(surveyKeyName)
		if err != nil {
			return nil, err
		}
	}
	if opts.Description {
		input.Description, err = askInput(surveyKeyDescription)
		if err != nil {
			return nil, err
		}
	}
	if opts.Homepage {
		input.Homepage, err = askInput(surveyKeyHomepage)
		if err != nil {
			return nil, err
		}
	}
	if opts.Private {
		var v string
		v, err = askSelect(surveyKeyPrivate, surveyBoolOptions)
		if err != nil {
			return nil, err
		}
		input.Private = v == "true"
	}
	if opts.Visibility {
		input.Visibility, err = askInput(surveyKeyDescription)
		if err != nil {
			return nil, err
		}
	}
	if opts.HasIssues {
		var v string
		v, err = askSelect(surveyKeyHasIssues, surveyBoolOptions)
		if err != nil {
			return nil, err
		}
		input.HasIssues = v == "true"
	}
	if opts.HasProjects {
		var v string
		v, err = askSelect(surveyKeyHasProjects, surveyBoolOptions)
		if err != nil {
			return nil, err
		}
		input.HasProjects = v == "true"
	}
	if opts.HasWiki {
		var v string
		v, err = askSelect(surveyKeyHasWiki, surveyBoolOptions)
		if err != nil {
			return nil, err
		}
		input.HasWiki = v == "true"
	}
	if opts.DefaultBranch {
		input.DefaultBranch, err = askInput(surveyKeyDefaultBranch)
		if err != nil {
			return nil, err
		}
	}
	if opts.AllowSquashMerge {
		var v string
		v, err = askSelect(surveyKeyAllowSquashMerge, surveyBoolOptions)
		if err != nil {
			return nil, err
		}
		input.AllowSquashMerge = v == "true"
	}
	if opts.AllowMergeCommit {
		var v string
		v, err = askSelect(surveyKeyAllowMergeCommit, surveyBoolOptions)
		if err != nil {
			return nil, err
		}
		input.AllowMergeCommit = v == "true"
	}
	if opts.AllowRebaseMerge {
		var v string
		v, err = askSelect(surveyKeyAllowRebaseMerge, surveyBoolOptions)
		if err != nil {
			return nil, err
		}
		input.AllowRebaseMerge = v == "true"
	}
	if opts.DeleteBranchOnMerge {
		var v string
		v, err = askSelect(surveyKeyDeleteBranchOnMerge, surveyBoolOptions)
		if err != nil {
			return nil, err
		}
		input.DeleteBranchOnMerge = v == "true"
	}
	if opts.Archived {
		var v string
		v, err = askSelect(surveyKeyArchived, surveyBoolOptions)
		if err != nil {
			return nil, err
		}
		input.Archived = v == "true"
	}

	return input, nil
}

func contains(target string, list []string) bool {
	for _, v := range list {
		if target == v {
			return true
		}
	}
	return false
}
