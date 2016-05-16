build:
	go build

docker:
	docker run --rm -v "$(GOPATH)":/go -w /go/src/tantalic.com/flickr-service blang/golang-alpine go build -v
	docker build -t flickr-service .
	rm flickr-service 

