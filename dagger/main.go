package main

import (
	"context"
	"fmt"
	"math"
	"math/rand"

	"dagger/dagger/internal/dagger"
)

type Dagger struct{}

// Publish the application container after building and testing it on-the-fly
func (m *Dagger) Publish(
	ctx context.Context,
	source *dagger.Directory,
	// +optional
	testCmd string,
	// +optional
	buildCmd string) (string, error) {
	_, err := m.Run(ctx, source, testCmd)
	if err != nil {
		return "", err
	}

	address, err := m.Build(source, buildCmd).
		Publish(ctx, fmt.Sprintf("ttl.sh/hello-dagger-%.0f", math.Floor(rand.Float64()*10000000))) //#nosec

	if err != nil {
		return "", err
	}

	return address, nil
}

// Build the application container
func (m *Dagger) Build(source *dagger.Directory, command string) *dagger.Container {

	build := dag.Node(dagger.NodeOpts{Ctr: m.BuildEnv(source)}).
		Commands().
		Run(buildCommandToSlice(command)).
		Directory("./dist")

	return dag.Container().From("nginx:1.25-alpine").
		WithDirectory("/usr/share/nginx/html", build).
		WithExposedPort(80)
}

// Return the result of running unit tests
func (m *Dagger) Run(ctx context.Context, source *dagger.Directory, command string) (string, error) {
	return dag.Node(dagger.NodeOpts{Ctr: m.BuildEnv(source)}).
		Commands().
		Run(testCommandToSlice(command)).
		Stdout(ctx)
}

// Build a ready-to-use development environment
func (m *Dagger) BuildEnv(source *dagger.Directory) *dagger.Container {
	return dag.Node(dagger.NodeOpts{Version: "21"}).
		WithNpm().
		WithSource(source).
		Install().
		Container()
}
