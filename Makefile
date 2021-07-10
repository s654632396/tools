.PHONY: echo test  test-coverage

UName := $(shell uname)
CurDate := $(shell date)
# argument pass by ENV
Arch ?= $(shell go env GOARCH)
GoVersion ?= $(shell go version)
GitVersion ?= $(shell git rev-parse --short HEAD || echo "GitNotFound")
BuildFlags ?= "-v"

echo:
	@echo ${UName}
	@echo ${Arch}
	@echo ${GoVersion}
	@echo "BuildFlags= "${BuildFlags}
	@echo "Git: "${GitVersion}
	@echo "Current Date: "${CurDate}

test-coverage:
	go test github.com/s654632396/seltree/seltree -cover

test: test-coverage
	go test github.com/s654632396/seltree/seltree -test.v



