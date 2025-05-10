run:
	go run main.go

test:
	go test ./... -count=1

testv:
	go test -v ./... -count=1
