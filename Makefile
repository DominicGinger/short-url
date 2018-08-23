default:
	go run main.go

build:
	rm -rf dist
	mkdir dist
	GOOS=linux go build main.go
	mv main dist/
	upx --brute dist/main

deploy:
	cp Dockerfile dist/
	cd dist && now --public
