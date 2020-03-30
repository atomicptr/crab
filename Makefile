.PHONY = build

meta_path = github.com/atomicptr/crab/pkg/meta
git_commit := $(shell git rev-list -1 HEAD)
git_version := $(shell git --no-pager tag --points-at HEAD | head -n 1)

build:
	go build -o bin/crab \
		-ldflags "\
			-X $(meta_path).Version=$(git_version) \
			-X $(meta_path).GitCommit=$(git_commit)" \
		cmd/crab/main.go
test:
	go test -v ./...
release-dryrun:
	GIT_COMMIT=$(git_commit) GIT_VERSION=$(git_version) goreleaser --snapshot --skip-publish --rm-dist
