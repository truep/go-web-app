build:
	go mod vendor
	go build -ldflags "-s -w" -o ./bin/web-app-truep ./cmd/webapp/main.go

.PHONY: clean

clean:
	rm -rf bin/*
