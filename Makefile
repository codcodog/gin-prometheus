name=ginprometheus

.PHONY: build
build:
	@go build -o $(name) .

.PHONY: run
run: build
	@./$(name) 

.PHONY: tidy
tidy:
	@go mod tidy
