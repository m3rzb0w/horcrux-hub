build:
	@npm --prefix front/ run build
	@go build -o bin/app .
run: build
	@./bin/app