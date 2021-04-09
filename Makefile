.PHONY: all
all:
	make -C setup
	make -C do

clean:
	make -C setup clean
	make -C do clean
	docker image rm lunatech.com/argo-demo-setup:latest
	docker image rm lunatech.com/argo-demo-do:latest
