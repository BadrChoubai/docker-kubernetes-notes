REPOSITORY = 685126695892.dkr.ecr.us-east-1.amazonaws.com
TAG = node-app-ecr-repo

all: # @HELP builds binaries for one platform ($OS/$ARCH)
all: build

build: # @HELP create docker image from Dockerfile
build:
	$(MAKE) docker-image

docker-image:
	docker build -t $(TAG) .

docker-push:
	$(MAKE) docker-tag
	docker push $(REPOSITORY)/$(TAG)

docker-run:
	docker run \
		-d --rm \
		--name node-esc-example--local \
		--publish 80:80 \
		--pull "always" \
		$(REPOSITORY)/$(TAG)

docker-tag:
	docker tag $(TAG):latest $(REPOSITORY)/$(TAG)

clean: # @HELP removes built binaries and temporary files
clean:

#	=========================================================================
#	HELP
#	=========================================================================

help: # @HELP prints this help message
help: help-docker
	echo "TARGETS:"
	grep -E '^.*: *# *@HELP' $(MAKEFILE_LIST)     \
	    | awk 'BEGIN {FS = ": *# *@HELP"}; { printf "  %-30s %s\n", $$1, $$2 };'

help-docker:
	echo "docker info:"
	echo "  BUILD_IMAGE = $(BUILD_IMAGE)"
	echo "  REGISTRY = $(REGISTRY)"
	echo "  TAG = $(TAG)"
	echo

.DEFAULT_GOAL := all
.SILENT: help help-docker
.PHONY: all build help help-docker