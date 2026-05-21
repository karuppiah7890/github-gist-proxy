# Contributing to this project

## Useful Extensions for VS Code

If you are using VS Code, then please consider downloading some of the following useful extensions (plugins) for ease of development. These extensions are also mentioned as recommendations in the repo inside `.vscode/extensions.json`. VS Code will automatically read `.vscode/extensions.json` and recommend the extensions to you when you open the repo in VS Code. Just in case that gets missed due to some reason, here are the list of useful extensions for VS Code

For Spell Check: https://marketplace.visualstudio.com/items?itemName=streetsidesoftware.code-spell-checker

For Docker related development (Dockerfile, Docker Compose): https://marketplace.visualstudio.com/items?itemName=docker.docker

For Golang Development: https://marketplace.visualstudio.com/items?itemName=golang.Go

## How to run it locally without using Containers or Kubernetes

### Building this project from source

This project has been built and tested with Go (Golang) version `1.26.3`

Please download and install Golang from https://go.dev/dl - preferably version `1.26.3` or newer but with same major version `1` - with newer minor or patch version

You can build this project by simply running

```bash
go build -v
```

### Running tests

To run tests quietly with short output, run this -

```bash
go test ./...
```

To run tests with verbose logging which has debug logs, run this -

```bash
go test -v ./...
```

### Running this project

After building it from source, you can run this project by simply running

```bash
./proxy
```

## How to run it locally with just Containers

There is a `Dockerfile` present for the service

You can build the container image for the service using a container build tool like `docker`. You can also choose other container build tools like `podman` etc

An example using `docker` CLI and Docker daemon -

Build `proxy` like this -

```bash
docker build --tag proxy .

# OR for verbose details -

DOCKER_BUILDKIT=0 docker build --tag proxy .
```

And if the build is cached, use below to create a new build without cache

```bash
docker build --no-cache --tag proxy .

# OR for verbose details -

DOCKER_BUILDKIT=0 docker build --no-cache --tag proxy .
```


Run `proxy` like this -

```bash
docker run --rm --publish 8080:8080 proxy
```

## How to run it locally with just Containers using Docker Compose

> [!NOTE]
> The `develop` feature in the Docker Compose configuration is available only
> Docker Compose Version `2.22.0` and later. So verify your Docker Compose
> Version, for example using `docker compose version`. If you use an older
> version than `2.22.0`, then please remove the `develop` section in the
> Docker Compose Configuration or else it will throw an error saying
> `services.proxy Additional property develop is not allowed`

The below command will build the two images and also run them for you

```bash
docker compose up --build --detach
```

The proxy is exposed to the host through 8080 port and is accessible only through the loopback network interface - `localhost` or `127.0.0.1` and not all network interfaces for security reasons - so as to not expose the service to outside devices that can access the host through other network interfaces

To run the containers without building the image, just run

```bash
docker compose up --no-build --detach
```

If you want to just build the two images, just run

```bash
docker compose build
```

And if the build is cached, use below to create a new build without cache

```bash
docker compose build --no-cache
```

If you want to develop the service and also rebuild the image whenever there are changes in the source code, then the `develop` feature in the Docker Compose Configuration will come in handy. You just need to run this

```bash
docker compose up --watch --detach
```

This will ensure that whenever there are source code changes in the service, the container image will be rebuilt and then the new container image will be started

To check the service status, use this

```bash
docker compose ps
```

To get the service logs, use this

```bash
docker compose logs
```

To follow the logs in real time, do this

```bash
docker compose logs -f
```

## Debugging Docker builds

To debug what's being sent as part of the build context, please use the `Dockerfile.debug` file like this -

```bash
DOCKER_BUILDKIT=0 docker build --file Dockerfile.debug
```

Accordingly modify the `.dockerignore` file

And if the build is cached, use below to create a new build without cache

```bash
DOCKER_BUILDKIT=0 docker build --no-cache --file Dockerfile.debug
```

## How to run it locally with just Containers using Kubernetes

We'll be using `helm` tool to deploy (install) and manage our service. Management means - get deployment information, upgrade our service, delete our service

Please install `helm` by following the official Helm website https://helm.sh/docs/intro/install or from the official Helm releases - https://github.com/helm/helm/releases

Once installed, also check if a tool like `minikube` or `kind` or similar is installed to run local Kubernetes clusters

You can install `minikube` or `kind` by following the instructions here - https://kubernetes.io/docs/tasks/tools/

We'll be using `minikube` with a driver like `docker` for example

```bash
minikube start
```

And once the Kubernetes cluster is ready, ensure that the container images are available in the worker node's container runtime. For example, for `minikube`, you can access the container runtime like this -

First get details to connect to the container daemon

```bash
minikube docker-env
```

For a specific minikube profile, you can do this -

```bash
minikube --profile <profile-name> docker-env
```

Then run the commands that it gives in your shell. You can also do this -

```bash
eval $(minikube docker-env)
```

Once you do this, you can see the list of containers running in the daemon like this -

```bash
docker ps
```

Now just build the image using `docker build` or `docker compose build`. We'll use `docker compose build` to build the image, like this -

```bash
docker compose build
```

Once done, check if the container daemon has the image of `github-gist-proxy` service named `proxy` using this -

```bash
docker images
```

You should see `proxy:latest`

We have a generic helm chart that can run any web service in general. By default it runs `nginx` for you. But you can pass in configuration to run any kind of simple web service

Now you can use the helm chart to run the service

```bash
helm install proxy helm-chart --set image.repository=proxy --set image.tag=latest --set service.port=8080 --set livenessProbe.httpGet.path=/livez --set readinessProbe.httpGet.path=/livez
```

To run it easily with lesser command line arguments, you can use the helm values yaml files for the service like this -

```bash
helm install proxy helm-chart --values helm-values.yaml
```

You can check the helm releases to see that they are installed and you can also follow the instructions in the post release notes to port forward the pod's container's port to the host to connect to it

```bash
helm ls
```

Next we'll be using `kubectl` to access the Kubernetes Cluster resources

Install `kubectl` from https://dl.k8s.io or by following https://kubernetes.io/docs/tasks/tools/

Or you can use `minikube kubectl` command to run `kubectl`

```bash
kubectl get deployments

kubectl get pods

kubectl get services
```
