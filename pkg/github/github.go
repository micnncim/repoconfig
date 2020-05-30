package github

// Repository represents a GitHub repository.
type Repository struct {
	Name                string `json:"name"`
	Description         string `json:"description"`
	Homepage            string `json:"homepage"`
	Private             bool   `json:"private"`
	Visibility          string `json:"visibility"`
	HasIssues           bool   `json:"has_issues"`
	HasProjects         bool   `json:"has_projects"`
	HasWiki             bool   `json:"has_wiki"`
	DefaultBranch       string `json:"default_branch"`
	AllowMergeCommit    bool   `json:"allow_merge_commit"`
	AllowRebaseMerge    bool   `json:"allow_rebase_merge"`
	AllowSquashMerge    bool   `json:"allow_squash_merge"`
	DeleteBranchOnMerge bool   `json:"delete_branch_on_merge"`
	Archived            bool   `json:"archived"`
	// IsTemplate          bool
}
