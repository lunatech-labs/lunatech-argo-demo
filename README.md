# Lunatech argo demo

Playground to learn argo workflows (and how to integrate them with scala).

## Installation

1. Turn on the local kubernetes cluster on docker for Mac if it's not yet running and switch to its context
2. [Install Argo](https://argoproj.github.io/argo-workflows/quick-start/):

       kubectl create ns argo
       kubectl create ns demo
       helm repo add argo https://argoproj.github.io/argo-helm
	   helm repo update
       helm install argo argo/argo -n argo

   Download the argo CLI [there](https://github.com/argoproj/argo-workflows/releases), install it, then check your installation by running:

       argo -n demo submit https://raw.githubusercontent.com/argoproj/argo-workflows/master/examples/workflow-template/hello-world.yaml --watch
       argo -n demo list
       argo -n demo get @latest
       argo -n demo logs @latest
	   
   If you want to use the web UI, change its k8s type to LoadBalancer with
   
       kubectl -n argo patch svc argo-server -p '{"spec": {"type": "LoadBalancer"}}'
	   
   then visit http://localhost:2746.

## Usage

### Submitting a workflow

Submit the workflow and watch it run:

	argo -n demo submit --watch argo-demo.yaml

You can now change the `logfile` parameter and see it taken into account in the workflow logs.

	argo -n demo submit argo-demo.yaml -p logfile=/app/data/toto.txt --watch
	argo -n demo logs @latest

### Templates

We can change the kind of this Workflow to a WorkflowTemplate one, to reuse it. First create it in the cluster:

    argo -n demo template create argo-demo-workflow-template.yaml
	argo template list -A

We can now submit a small workflow containing only a reference to this WorkflowTemplate:

    argo -n demo submit argo-demo-using-template-no-params.yaml --watch

We can change the parameters in two ways: from within a workflow definition or as CLI arguments:

	argo -n demo submit argo-demo-using-template.yaml --watch
	argo -n demo logs @latest
	argo -n demo submit argo-demo-using-template-no-params.yaml --watch -p logfile="/app/data/template-workflow-cli-params.txt"
	argo -n demo logs @latest

### Using the REST API

First let's clean the namespace

	argo template list -A
	argo -n demo template delete argo-demo-workflow-template
	kubectl -n demo delete pod -l workflows.argoproj.io/completed=true`

Now let's re-add the template and use it in the same workflows

    curl -H 'Content-Type: application/json' http://localhost:2746/api/v1/workflow-templates/demo -d @argo-demo-workflow-template.json 
	curl -H 'Content-Type: application/json' http://localhost:2746/api/v1/workflows/demo -d @argo-demo-using-template-no-params.json 
	argo -n demo watch @latest
	argo -n demo logs @latest
    curl -H 'Content-Type: application/json' http://localhost:2746/api/v1/workflows/demo -d @argo-demo-using-template.json
	argo -n demo logs @latest

    curl http://localhost:2746/api/v1/workflows/demo | jq .items[].metadata.name

### Using the GUI

TODO

## Notes

- To delete completed jobs, run `k -n demo delete pod -l workflows.argoproj.io/completed=true`

- To install and use [Minio](https://min.io/):

       helm repo add minio https://helm.min.io
	   helm repo update
       helm install minio minio/minio --set service.type=LoadBalancer

   Create an alias for the argo artifacts:

       export MINIO_ACCESS_KEY=$(kubectl get secret minio --namespace default -o jsonpath="{.data.accesskey}" | base64 --decode)
       export MINIO_SECRET_KEY=$(kubectl get secret minio --namespace default -o jsonpath="{.data.secretkey}" | base64 --decode)
       mc alias set minio http://localhost:9000 "$MINIO_ACCESS_KEY" "$MINIO_SECRET_KEY" --api s3v4

   And now a bucket to host our files:
   
       mc mb minio/argo-artifacts

## To build locally

1. [Install Go](https://golang.org/doc/tutorial/getting-started)
2. Run `make`, check that the docker images are built with `docker image ls | grep argo-demo`. You should see `lunatech-labs/argo-demo-do` and `lunatech-labs/argo-demo-setup`
3. Run these commands to see it working locally:

       docker run -v $(pwd):/app/data --rm lunatech-labs/argo-demo-setup:latest
       docker run -v $(pwd):/app/data --rm lunatech-labs/argo-demo-do:latest
	
   You should see `Writing to /app/data/log.txt from setup and from do` and the file `log.txt` should have this content.

## To do

[ ] show how to use minio to manage artifacts


