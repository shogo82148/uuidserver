.PHONY: build
build:
	sam build

.PHONY: deploy
deploy:
	sam deploy

.PHONY: build-UuidFunction
build-UuidFunction:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o $(ARTIFACTS_DIR)/bootstrap .
