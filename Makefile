PROJECT_NAME := "testServer" 

dev:
	go run main.go

lint:
	golint -set_exit_status .

race: dep ## Run data race detector
	@go test -race -short ./...

msan: dep ## Run memory sanitizer
	@go test -msan -short ./...

dep: ## Get the dependencies
	@go get -v -d -t ./...

build:
	docker build -t test_server .

run:
	docker run -ti -p 80:80 --rm --name test-api test_server