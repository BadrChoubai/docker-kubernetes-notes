# Docker and Kubernetes

This project contains source code, artifacts, and written notes created while learning the foundations of Docker and
Kubernetes.

**Detailed Course Notes**:

- [Docker](./Docker/README.md)
- [Kubernetes](./Kubernetes/README.md)

**Sections**:

1. [Introduction to Docker](#introduction-to-docker)
2. [Introduction to Kubernetes](#introduction-to-kubernetes)

---

## Introduction to Docker

Conceptually, Docker is used to manage two things: **Containers** and **Images**. Containers isolate and run
applications in a consistent, portable environment, while images serve as the templates
from which these containers are created.

### 1. Containers

- A **container** is an isolated, lightweight, and portable package that includes the application code, runtime,
  libraries, and dependencies needed to run a piece of software. Containers package software in a way that makes it easy
  to run consistently across different environments, from local development machines to cloud servers.

- **Isolated**: Containers are isolated from the underlying host system and from each other. They have their own
  filesystem, processes, network interfaces, and resource limits, which ensures that changes inside a container don’t
  affect the host or other containers.

- **Single-task focus**: Containers are designed to run **one task** or **one service**. For example, a container might
  run a web server, a database, or a background worker. This single-task design aligns with the microservices
  architecture, where applications are broken down into small, independently deployable services.

- **Shareable and reproducible**: Containers are portable, meaning they can be shared and run on any system that
  supports Docker. This ensures that if a container works on a developer’s laptop, it will work the same way in
  production, providing consistency and eliminating "it works on my machine" issues.

- **Stateless**: Containers are designed to be **stateless**, meaning they don’t persist any data by default once they
  stop running. To manage persistent data (such as databases or user uploads), you need to use **volumes** or **bind
  mounts** to store data outside the container.
    - Example: If a container running a database stops, the data will be lost unless a volume is used to persist it.
    - Stateless containers are easier to scale, as new instances can be spun up without needing to worry about state
      synchronization.

### 2. Images

- A **Docker image** is a **blueprint** used to create containers. It contains everything the container needs to run,
  including the application code, runtime, environment variables, libraries, and system tools. Think of an image as a
  **snapshot** of an environment at a specific point in time.

- **Code + environment**: A Docker image bundles the application code along with its runtime environment. This includes
  the OS libraries, configurations, and any dependencies the application needs to run. This bundling ensures that the
  application runs consistently, regardless of the environment in which it is deployed.

- **Read-only**: Docker images are **read-only**, meaning they cannot be modified once created. When a container is
  created from an image, Docker adds a writable layer on top of the read-only image where changes can be made. These
  changes, however, do not affect the underlying image and are discarded when the container stops, unless they are saved
  or persisted (for example, using volumes).

- **Does not run**: An image is **static** and does not execute by itself. To run an application, Docker takes an image
  and creates a running instance of it, called a **container**.

- **Built + shared**: Docker images can be created (or **built**) using a `Dockerfile`. A `Dockerfile` is a script that
  defines the instructions for assembling the image, such as installing dependencies, copying files, and setting up the
  environment. Once built, images can be shared through a **Docker registry** (like Docker Hub or a private registry),
  allowing other users or systems to pull the image and create containers from it.
    - Example: An image of a Python web app might include the Python runtime, the app’s code, and the necessary
      libraries.
    - Docker images can be tagged with versions (e.g., `my-app:1.0`, `my-app:latest`), allowing different versions of
      the same application to be stored and shared.

---

## Introduction to Kubernetes

When you're managing a large-scale application with numerous containers, manually handling each of them becomes
unmanageable due to several factors:

- **Scalability Issues**: Manually spinning up new containers, configuring them, or managing their lifecycles is a
  complex task when the number of containers grows. This leads to errors like incorrect configuration settings, missing
  environment variables, and version mismatches.

- **Repetitive Tasks**: Tasks such as updating a container, deploying new versions, or restarting a crashed instance are
  repetitive. Performing them manually wastes time and increases the chance for human error.

- **Environment Consistency**: Manually deploying containers in different environments (dev, test, prod) can result in
  inconsistencies due to differences in configuration and dependencies. Ensuring that a container works in different
  environments without issues is tough to manage without automation.

- **Complex Dependencies**: Modern applications are often made up of microservices, with each microservice running in
  its own container. Managing inter-service communication, ensuring the proper order of service deployment, and handling
  dependencies manually is error-prone.

- **Version Management**: Deploying different versions of the same service (e.g. for A/B testing or canary releases)
  manually requires extra effort to track which version is running where and how the versions interact with other
  services.

Kubernetes introduces automation to the deployment, scaling, and management of containerized applications, significantly
reducing the chance of error while improving the overall efficiency of managing large-scale container environments.