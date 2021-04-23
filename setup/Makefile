version := 0.0.1

docker: setup
	docker build . -t lunatech-labs/argo-demo-setup:${version} --no-cache

setup: setup.go
	CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' -o setup

clean:
	rm -f setup log.txt
