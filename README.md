# Lunatech argo demo

Playground to learn argo workflows (and how to integrate them with scala).

## Installation

1. Turn on the local kubernetes cluster on docker for Mac if it's not yet running and switch to its context
2. [Install Argo](https://argoproj.github.io/argo-workflows/quick-start/):

	```helm repo add argo https://argoproj.github.io/argo-helm
       helm install argo argo/arg```

   Download the argo CLI [there](https://github.com/argoproj/argo-workflows/releases), install it, then check your installation by running:

    ```argo submit https://raw.githubusercontent.com/argoproj/argo/master/examples/hello-world.yaml --watch
    argo list
    argo get @latest
    argo logs @latest```

3. [Install Go](https://golang.org/doc/tutorial/getting-started)
4. Run `make`, check that the docker images are built with `docker image ls | grep argo-demo`. You should see `lunatech.com/argo-demo-do` and `lunatech.com/argo-demo-setup`
5. Run these commands to see it working locally:

	```docker run -v $(pwd):/app/data --rm lunatech.com/argo-demo-setup:latest
	docker run -v $(pwd):/app/data --rm lunatech.com/argo-demo-do:latest```
	
	You should see `Final message: Writing from setup and from do` and the file `log.txt` should have this content.

