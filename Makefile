default:
	go run main.go

build:
	GOOS=linux go build main.go
	docker build . -t short

