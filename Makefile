.PHONY: build clean deploy

build:
	
	env GOOS=linux go build -ldflags="-s -w" -o ./bin/api ./main.go
	env GOOS=linux go build -ldflags="-s -w" -o ./bin/sendmail ./sendmail/main.go

clean:
	rm -rf ../bin

deploy: clean build
	sls deploy --verbose
