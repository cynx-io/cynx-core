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


TAG := v0.0.33
COMMIT_MSG := "nil req"

.PHONY: publish
publish:
	git add .
	git commit -m $(COMMIT_MSG)
	git push origin main
	git tag -a $(TAG) -m $(COMMIT_MSG)
	git push origin $(TAG)
	buf push --label $(TAG)

