# repoconfig

[![actions-workflow-test][actions-workflow-test-badge]][actions-workflow-test]
[![release][release-badge]][release]
[![pkg.go.dev][pkg.go.dev-badge]][pkg.go.dev]
[![license][license-badge]][license]

A CLI to execute a bulk update for GitHub repository configs.

Some configs that `repoconfig` can set are below.

![screenshot](docs/assets/screenshot.png)

## Prerequisites

Create a GitHub token and set it via the environment variable `GITHUB_TOKEN`.

## Installation

```console
$ go get github.com/micnncim/repoconfig/cmd/repoconfig
```

## Usage

If you feel worried about how `repoconfig` works, setting `--dry-run` at first is recommended.
With `--dry-run`, `repoconfig` doesn't update repository configs at all.

The form of arguments is `OWNER` or `OWNER/REPO`.
Multiple arguments are acceptable.

**Caution**: If `OWNER` is specified, rather than `OWNER/REPO`, `repoconfig` will update **all** the repositories of the `OWNER`.

The following help describes the flags.

```console
$ repoconfig --help
CLI to update repository configs

Usage:
  repoconfig [flags]

Flags:
      --allow-merge-commit       Whether to allow merging pull requests with a merge commit (default true)
      --allow-rebase-merge       Whether to allow rebase-merging pull requests (default true)
      --allow-squash-merge       Whether to allow allow squash-merging pull requests (default true)
      --default-branch string    The default branch for a repository (default "master")
      --delete-branch-on-merge   Whether to allow automatically deleting head branches when pull requests are merged
      --dry-run                  Whether user dry-run mode
      --has-issues               Whether a repository has issues (default true)
      --has-projects             Whether a repository has projects (default true)
      --has-wiki                 Whether a repository has wiki (default true)
      -h, --help                 help for repoconfig
```

## Example

```console
$ repoconfig micnncim micnncim/github-lab micnncim/github-actions-lab --delete-branch-on-merge=true
2020-04-30T02:58:31.558+0900    INFO    app.github      successfully updated repository {"owner": "micnncim", "repo": "github-lab", "update_repository_options": {"has_issues":true,"has_projects":true,"has_wiki":true,"default_branch":"master","allow_squash_merge":true,"allow_merge_commit":true,"allow_rebase_merge":true,"delete_branch_on_merge":true}, "dry_run": false}
2020-04-30T02:58:32.138+0900    INFO    app.github      successfully updated repository {"owner": "micnncim", "repo": "github-actions-lab", "update_repository_options": {"has_issues":true,"has_projects":true,"has_wiki":true,"default_branch":"master","allow_squash_merge":true,"allow_merge_commit":true,"allow_rebase_merge":true,"delete_branch_on_merge":true}, "dry_run": false}
```

```console
$ repoconfig monalisa --delete-branch-on-merge=true
2020-04-30T02:58:31.558+0900    INFO    app.github      successfully updated repository {"owner": "monalisa", "repo": "monalisa-repo1", "update_repository_options": {"has_issues":true,"has_projects":true,"has_wiki":true,"default_branch":"master","allow_squash_merge":true,"allow_merge_commit":true,"allow_rebase_merge":true,"delete_branch_on_merge":true}, "dry_run": false}
2020-04-30T02:58:32.138+0900    INFO    app.github      successfully updated repository {"owner": "monalisa", "repo": "monalisa-repo2", "update_repository_options": {"has_issues":true,"has_projects":true,"has_wiki":true,"default_branch":"master","allow_squash_merge":true,"allow_merge_commit":true,"allow_rebase_merge":true,"delete_branch_on_merge":true}, "dry_run": false}
```

## References

- [Repositories | GitHub Developer Guide](https://developer.github.com/v3/repos/#update-a-repository)

<!-- badge links -->

[actions-workflow-test]: https://github.com/micnncim/repoconfig/actions?query=workflow%3ATest
[actions-workflow-test-badge]: https://img.shields.io/github/workflow/status/micnncim/repoconfig/Test?label=Test&style=for-the-badge&logo=github

[release]: https://github.com/micnncim/repoconfig/releases
[release-badge]: https://img.shields.io/github/v/release/micnncim/repoconfig?style=for-the-badge&logo=github

[pkg.go.dev]: https://pkg.go.dev/github.com/micnncim/repoconfig?tab=overview
[pkg.go.dev-badge]: https://img.shields.io/badge/pkg.go.dev-reference-02ABD7?style=for-the-badge&logoWidth=25&logo=data%3Aimage%2Fsvg%2Bxml%3Bbase64%2CPHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHZpZXdCb3g9Ijg1IDU1IDEyMCAxMjAiPjxwYXRoIGZpbGw9IiMwMEFERDgiIGQ9Ik00MC4yIDEwMS4xYy0uNCAwLS41LS4yLS4zLS41bDIuMS0yLjdjLjItLjMuNy0uNSAxLjEtLjVoMzUuN2MuNCAwIC41LjMuMy42bC0xLjcgMi42Yy0uMi4zLS43LjYtMSAuNmwtMzYuMi0uMXptLTE1LjEgOS4yYy0uNCAwLS41LS4yLS4zLS41bDIuMS0yLjdjLjItLjMuNy0uNSAxLjEtLjVoNDUuNmMuNCAwIC42LjMuNS42bC0uOCAyLjRjLS4xLjQtLjUuNi0uOS42bC00Ny4zLjF6bTI0LjIgOS4yYy0uNCAwLS41LS4zLS4zLS42bDEuNC0yLjVjLjItLjMuNi0uNiAxLS42aDIwYy40IDAgLjYuMy42LjdsLS4yIDIuNGMwIC40LS40LjctLjcuN2wtMjEuOC0uMXptMTAzLjgtMjAuMmMtNi4zIDEuNi0xMC42IDIuOC0xNi44IDQuNC0xLjUuNC0xLjYuNS0yLjktMS0xLjUtMS43LTIuNi0yLjgtNC43LTMuOC02LjMtMy4xLTEyLjQtMi4yLTE4LjEgMS41LTYuOCA0LjQtMTAuMyAxMC45LTEwLjIgMTkgLjEgOCA1LjYgMTQuNiAxMy41IDE1LjcgNi44LjkgMTIuNS0xLjUgMTctNi42LjktMS4xIDEuNy0yLjMgMi43LTMuN2gtMTkuM2MtMi4xIDAtMi42LTEuMy0xLjktMyAxLjMtMy4xIDMuNy04LjMgNS4xLTEwLjkuMy0uNiAxLTEuNiAyLjUtMS42aDM2LjRjLS4yIDIuNy0uMiA1LjQtLjYgOC4xLTEuMSA3LjItMy44IDEzLjgtOC4yIDE5LjYtNy4yIDkuNS0xNi42IDE1LjQtMjguNSAxNy05LjggMS4zLTE4LjktLjYtMjYuOS02LjYtNy40LTUuNi0xMS42LTEzLTEyLjctMjIuMi0xLjMtMTAuOSAxLjktMjAuNyA4LjUtMjkuMyA3LjEtOS4zIDE2LjUtMTUuMiAyOC0xNy4zIDkuNC0xLjcgMTguNC0uNiAyNi41IDQuOSA1LjMgMy41IDkuMSA4LjMgMTEuNiAxNC4xLjYuOS4yIDEuNC0xIDEuN3oiLz48cGF0aCBmaWxsPSIjMDBBREQ4IiBkPSJNMTg2LjIgMTU0LjZjLTkuMS0uMi0xNy40LTIuOC0yNC40LTguOC01LjktNS4xLTkuNi0xMS42LTEwLjgtMTkuMy0xLjgtMTEuMyAxLjMtMjEuMyA4LjEtMzAuMiA3LjMtOS42IDE2LjEtMTQuNiAyOC0xNi43IDEwLjItMS44IDE5LjgtLjggMjguNSA1LjEgNy45IDUuNCAxMi44IDEyLjcgMTQuMSAyMi4zIDEuNyAxMy41LTIuMiAyNC41LTExLjUgMzMuOS02LjYgNi43LTE0LjcgMTAuOS0yNCAxMi44LTIuNy41LTUuNC42LTggLjl6bTIzLjgtNDAuNGMtLjEtMS4zLS4xLTIuMy0uMy0zLjMtMS44LTkuOS0xMC45LTE1LjUtMjAuNC0xMy4zLTkuMyAyLjEtMTUuMyA4LTE3LjUgMTcuNC0xLjggNy44IDIgMTUuNyA5LjIgMTguOSA1LjUgMi40IDExIDIuMSAxNi4zLS42IDcuOS00LjEgMTIuMi0xMC41IDEyLjctMTkuMXoiLz48L3N2Zz4=

[license]: LICENSE
[license-badge]: https://img.shields.io/github/license/micnncim/repoconfig?style=for-the-badge
