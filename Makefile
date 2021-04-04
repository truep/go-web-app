app_name = web-app-truep
src_path = ./cmd/webapp/*.go
bin_path = ./bin/
build_flags = -ldflags "-s -w" 

build:
	go mod tidy && go mod vendor
	go build $(build_flags) -o $(bin_path)$(app_name) $(src_path)

.PHONY: clean

test: 
	go test ./...

clean:
	rm -rf $(bin_path)

all: test build