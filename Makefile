tidy:
	go mod tidy
	go fmt ./src/...
	fieldalignment -fix ./src/...
	go vet ./src/...
	golangci-lint run --fix ./src/...
	staticcheck ./src/...

.PHONY: proto
proto:
	buf dep update
	buf generate

.PHONY: publish
publish: proto
	@echo "Publishing to GitHub..."
	git add .
	git commit -m "proto"
	git push origin main
	git tag -a v0.0.20 -m ""
	git push origin v0.0.20

