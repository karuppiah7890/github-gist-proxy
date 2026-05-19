# Contributing to this project

## How to run it locally without using Containers or Kubernetes

### Building this project from source

This project has been built and tested with Go (Golang) version `1.26.3`

Please download and install Golang from https://go.dev/dl - preferably version `1.26.3` or newer but with same major version `1` - with newer minor or patch version

You can build this project by simply running

```bash
go build -v
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
