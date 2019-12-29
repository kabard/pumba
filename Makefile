.Phony: all

default: check-quality test build

ALL_PACKAGES=$(shell go list ./... | grep -v /vendor)
WORKDIR=$(shell echo "${PWD}")
APPLICATION_YAML=$(shell echo "$(WORKDIR)/application.yml")
SOURCE_DIRS=$(shell go list ./... | grep -v /vendor | grep -v /out | cut -d "/" -f2 | uniq | grep -v ^pumba$)
APP_EXECUTABLE="out/pumba"

setup: 
	go get -u golang.org/x/tools/cmd/goimports
	go get -u golang.org/x/lint/golint
	go get -u github.com/mattn/goveralls

check-quality: lint fmt vet

lint:
	env GO111MODULE=on golint -set_exit_status $(ALL_PACKAGES)

fmt:
	@gofmt -l -s -w $(SOURCE_DIRS)

imports:
	@goimports -l -w -v $(SOURCE_DIRS)

vet:
	env GO111MODULE=on go vet ./...

test:
	GO111MODULE=on go clean -testcache ./... && go test -v ./...

build:
	mkdir -p out/
	GO111MODULE=on go build -o $(APP_EXECUTABLE) ./cmd/pumba

coverage:
	GO111MODULE=on ENVIRONMENT=test goveralls -service=travis-ci

test-cover-html:
	@echo "mode: count" > coverage-all.out

	$(foreach pkg, $(ALL_PACKAGES),\
	ENVIRONMENT=test go test -coverprofile=coverage.out -covermode=count $(pkg);\
	tail -n +2 coverage.out >> coverage-all.out;)
	go tool cover -html=coverage-all.out -o out/coverage.html