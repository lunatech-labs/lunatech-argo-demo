# Lunatech argo demo

Playground to learn argo workflows (and how to integrate them with scala).

## Installation

1. Turn on the local kubernetes cluster on docker for Mac if it's not yet running and switch to its context
2. Install [Minio](https://min.io/):

       helm repo add minio https://helm.min.io
	   helm repo update
	   helm install minio minio/minio --set service.type=LoadBalancer

   Create an alias for the argo artifacts:

       export MINIO_ACCESS_KEY=$(kubectl get secret minio --namespace default -o jsonpath="{.data.accesskey}" | base64 --decode)
       export MINIO_SECRET_KEY=$(kubectl get secret minio --namespace default -o jsonpath="{.data.secretkey}" | base64 --decode)
       mc alias set minio http://localhost:9000 "$MINIO_ACCESS_KEY" "$MINIO_SECRET_KEY" --api s3v4

   And now a bucket to host our files:
   
       mc mb minio/argo-artifacts

3. [Install Argo](https://argoproj.github.io/argo-workflows/quick-start/):

       kubectl create ns argo
       kubectl create ns demo
       helm repo add argo https://argoproj.github.io/argo-helm
	   helm repo update
       helm install argo argo/argo -n argo

   Download the argo CLI [there](https://github.com/argoproj/argo-workflows/releases), install it, then check your installation by running:

       argo submit https://raw.githubusercontent.com/argoproj/argo/master/examples/hello-world.yaml --watch
       argo list
       argo get @latest
       argo logs @latest
	   
   If you want to use the web UI, change its k8s type to LoadBalancer with
   
       kubectl patch svc argo-server -p '{"spec": {"type": "LoadBalancer"}}'
	   
   then visit http://localhost:2746.

4. [Install Go](https://golang.org/doc/tutorial/getting-started)
5. Run `make`, check that the docker images are built with `docker image ls | grep argo-demo`. You should see `lunatech.com/argo-demo-do` and `lunatech.com/argo-demo-setup`
6. Run these commands to see it working locally:

       docker run -v $(pwd):/app/data --rm lunatech.com/argo-demo-setup:latest
       docker run -v $(pwd):/app/data --rm lunatech.com/argo-demo-do:latest
	
   You should see `Writing to /app/data/log.txt from setup and from do` and the file `log.txt` should have this content.

7. Now submit the workflow and watch it run:

       argo -n submit demo --watch argo-demo.yaml

   You can now change the logfile parameter and see it taken into account in the workflow logs.

       argo -n demo submit argo-demo.yaml -p logfile=/app/data/toto.txt --watch
       argo -n demo logs @latest

## Notes

- To delete completed jobs, run `k -n demo delete pod -l workflows.argoproj.io/completed=true`


## To do

[ ] show how to use minio to manage artifacts



