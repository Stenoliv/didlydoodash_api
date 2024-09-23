run: build
	@./bin/discgolf.exe

dev:
	go run src/cmd/api/main.go

build:
	@go build -o bin/discgolf.exe src/cmd/api/main.go

database:
	@go run src/cmd/migrate/main.go

drop:
	@go run src/cmd/drop/main.go