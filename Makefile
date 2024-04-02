test:
	go test -race -shuffle=on ./...

run:
	go run ./app/main.go