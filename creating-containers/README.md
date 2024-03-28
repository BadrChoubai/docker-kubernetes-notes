# Creating Containers

Docker lets you build your own container images or use pre-built
ones. Images are the "blueprints" for creating containers, which run and
execute your code.

## Images

Some notes on images:

1. Images are immutable meaning that once you've created one from running `docker build`,
you will need to re-build your image each time you'd like an external change to
be picked up.
2. Images are layer-based meaning that if you build or re-build an image only 
the instructions that detect changes will be executed inside your `Dockerfile`.


### Friendly Manuals

1. [Layers](https://docs.docker.com/build/guide/layers/)
2. [Multi-Stage Builds](https://docs.docker.com/build/guide/multi-stage/)
