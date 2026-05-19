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
