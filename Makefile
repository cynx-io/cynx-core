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


TAG := v0.0.24
COMMIT_MSG := "typo"

.PHONY: publish
publish: proto
	git add .
	git commit -m $(COMMIT_MSG)
	git push origin main
	git tag -a $(TAG) -m $(COMMIT_MSG)
	git push origin $(TAG)

