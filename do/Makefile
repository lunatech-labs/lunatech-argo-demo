docker: do
	docker build . -t lunatech.com/argo-demo-do:latest --no-cache

do: do.go
	CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' -o do

clean:
	rm -f do log.txt
