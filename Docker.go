# To build a container image in Go, you can use the docker package, which is a Go client library for the Docker API. This package provides an API for interacting with the Docker daemon and can be used to programmatically create, manage, and build container images. Here's an example of how to use the docker package to build a container image:

package main

import (
    "context"
    "io"
    "log"

    "github.com/docker/docker/api/types"
    "github.com/docker/docker/client"
)

func main() {
    cli, err := client.NewClientWithOpts(client.FromEnv)
    if err != nil {
        log.Fatal(err)
    }

    buildCtx, cancel := context.WithCancel(context.Background())
    defer cancel()

    buildOptions := types.ImageBuildOptions{
        Tags:       []string{"image-name:latest"},
        Remove:     true,
        Dockerfile: "path/to/Dockerfile",
    }

    buildResponse, err := cli.ImageBuild(buildCtx, nil, buildOptions)
    if err != nil {
        log.Fatal(err)
    }
    defer buildResponse.Body.Close()

    _, err = io.Copy(os.Stdout, buildResponse.Body)
    if err != nil {
        log.Fatal(err)
    }
}



# This code first imports the docker package and creates a new client. Then, it creates a context for the build process and sets the build options, including the image name and the path to the Dockerfile. The ImageBuild function is then called with the context, options and a reader as arguments. It returns a response that includes the build output in the Body field. The body is then copied to the standard output so that the user can see the output of the build.

It's worth noting that the docker package requires that the Docker daemon is running and that the user running the Go code has permission to access the Docker daemon.

You can also use other libraries such as docker-go-api or go-dockerclient to build container images as well.
