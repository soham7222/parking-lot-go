run_local:
	go run main.go

install_deps:
	go get ./...
	go mod download
	go mod tidy
	go mod vendor

build_local:
	go build .

test_local:
	go test ./... -p=1 -count=1

test_local_with_coverage:
	go test ./... -p=1 -count=1 -coverprofile=cover.out

