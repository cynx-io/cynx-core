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


TAG := v0.0.38
COMMIT_MSG := "aws SES email"

publish_proto:
	buf push --label $(TAG)

.PHONY: publish
publish: tidy
	git add .
	git commit -m $(COMMIT_MSG)
	git push origin main
	git tag -a $(TAG) -m $(COMMIT_MSG)
	git push origin $(TAG)
	buf push --label $(TAG)

