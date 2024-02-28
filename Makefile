build-docker:
	docker build -t golang-ci:${VERSION} . --build-arg VERSION=${VERSION}
build:
	go build  -o app main.go
run:
	./app || go run main.go
