package github

// UpdateRepositoryRequest is a request for updating a repository.
// https://developer.github.com/v3/repos/#update-a-repository
// The fields intentionally commented out are dangerous for bulk update.
type UpdateRepositoryOptions struct {
	// Name                string
	// Description         string
	// Homepage            string
	// Private             bool
	// Visibility          bool
	HasIssues   bool `json:"has_issues"`
	HasProjects bool `json:"has_projects"`
	HasWiki     bool `json:"has_wiki"`
	// IsTemplate         bool
	DefaultBranch       string `json:"default_branch"`
	AllowSquashMerge    bool   `json:"allow_squash_merge"`
	AllowMergeCommit    bool   `json:"allow_merge_commit"`
	AllowRebaseMerge    bool   `json:"allow_rebase_merge"`
	DeleteBranchOnMerge bool   `json:"delete_branch_on_merge"`
	// Archived            bool
}
