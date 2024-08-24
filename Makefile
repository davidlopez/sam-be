ARTIFACT_NAME=sam-be

build:
	@go build -o bin/${ARTIFACT_NAME}/${ARTIFACT_NAME} cmd/${ARTIFACT_NAME}/main.go

run:
	@go run cmd/${ARTIFACT_NAME}/main.go