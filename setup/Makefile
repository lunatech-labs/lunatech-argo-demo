docker: setup
	docker build . -t lunatech.com/argo-demo-setup:latest --no-cache

setup: setup.go
	CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' -o setup

clean:
	rm -f setup log.txt
	docker image rm lunatech.com/argo-demo-setup:latest
