# Docker Core Concepts

**Sections**:

- [Docker Commands](#key-docker-commands)
- [Data, Volumes, and Networking](#data-volumes-and-networking)
- [Docker vs. Docker Compose](#docker-vs-docker-compose)
- [Local vs. Remote Development](#local-vs-remote-development-using-docker)
- [Deployment Considerations](#deployment-considerations)

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
  mounts** to store data outside of the container.
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

## Key Docker Commands

### 1. `docker build`

This command creates a Docker image from a Dockerfile and the specified build context (usually the current directory).
The image can then be run as a container.

```sh
docker build -t NAME:TAG .
```

- `-t NAME:TAG`: Assigns a **name** and **tag** (version) to the image being built. The tag is optional but useful for
  versioning the image (e.g., `my-app:1.0`).
    - Example: `docker build -t my-app:1.0 .` would create an image named `my-app` with a version tag of `1.0`.
- `.`: The **build context**, which is the directory containing the `Dockerfile` and any files referenced in the image
  build (like source code or configuration). Here, `.` represents the current directory.
    - Docker uses this context to access files needed to assemble the image.

Additional options for `docker build` include:

- `--no-cache`: Build the image without using any cached layers, forcing Docker to rebuild everything from scratch.
    - Example: `docker build --no-cache -t my-app:1.0 .`
- `-f Dockerfile.custom`: Use a Dockerfile other than the default one in the build context.
    - Example: `docker build -t my-app:1.0 -f Dockerfile.custom .`


### 2. `docker run`

This command runs a container from a Docker image. It creates and starts the container and can be customized with
various options.

```sh
docker run --name NAME --rm -d IMAGE
```

- `--name NAME`: Specifies the **name** of the container. This makes it easier to reference the container for
  management (e.g., stopping or viewing logs). If not provided, Docker assigns a random name.
    - Example: `docker run --name my-container my-image` will run a container named `my-container` from the image
      `my-image`.

- `--rm`: Automatically removes the container when it stops. This is useful for temporary or one-off containers where
  you don’t need to keep any data after they exit.
    - Example: `docker run --rm my-image` will delete the container once it stops.

- `-d`: Runs the container in **detached mode**, meaning it runs in the background without tying up the terminal. To
  interact with or view the logs from the container, you would need to use commands like `docker logs` or `docker exec`.
    - Example: `docker run -d my-image` starts the container in the background, and you can later check its logs with
      `docker logs`.

Additional useful options:

- `-p HOST_PORT:CONTAINER_PORT`: Maps a port from the host to a port in the container, allowing external access to
  services running inside the container.
    - Example: `docker run -p 8080:80 my-image` exposes port 80 in the container to port 8080 on the host.

- `-e "ENV_VAR=value"`: Passes an environment variable into the container.
    - Example: `docker run -e "ENV=prod" my-image` sets an environment variable `ENV` to `prod` inside the container.


### 3. `docker push`

This command pushes a locally built image to a Docker registry, such as Docker Hub or a private registry, making it
available for others or for use in production environments.

```sh
docker push REGISTRY_URL/NAME:TAG
```

- `REGISTRY_URL`: The URL of the Docker registry you are pushing to. For Docker Hub, you can omit the registry URL, and
  for private registries, it might look something like `my-registry.com`.
    - Example: `docker push my-registry.com/my-app:1.0`

- `NAME:TAG`: The **name** and **tag** of the image to push. The tag allows you to specify a version or other
  identifier (like `latest` or `v2.0`).
    - Example: `docker push my-app:1.0` pushes the image tagged `1.0`.

If you're using a private registry, make sure you're logged in using `docker login` before pushing images.


### 4. `docker pull`

This command retrieves (pulls) an image from a Docker registry, allowing you to use it to run containers on your local
machine.

```sh
docker pull REGISTRY_URL/NAME:TAG
```

- `REGISTRY_URL`: The URL of the registry from which you're pulling the image. For Docker Hub, this can be omitted, and
  it will default to pulling from Docker Hub.
    - Example: `docker pull my-registry.com/my-app:1.0`

- `NAME:TAG`: The **name** and **tag** of the image to pull. If no tag is provided, Docker defaults to pulling the
  `latest` tag.
    - Example: `docker pull my-app:1.0` pulls the version tagged `1.0`.
    - Example: `docker pull my-app` pulls the image tagged `latest`.

Pulled images are stored locally and can be viewed using the `docker images` command.

---

## Data, Volumes, and Networking

By default, containers are **isolated** and **stateless**.

- **Isolation** means that containers don’t share processes, files, or networks with the host or other containers by
  default.
- **Statelessness** means that any data created inside a container is lost once the container stops, unless steps are
  taken to persist it (such as using volumes or bind mounts).

To persist data or share files between the host and containers, Docker provides **Bind Mounts** and **Volumes**.

### Bind Mounts

Bind mounts allow you to mount a file or directory from your host system into a container. This is useful in development
environments where you want changes made on your host to immediately reflect in the container, and vice versa. However,
bind mounts are tied to the host filesystem, making them less portable.

- Example:

  ```sh
  docker run -v /host/path:/container/path <image>
  ```

  In this example, the directory `/host/path` from your host machine is mounted inside the container at
  `/container/path`.

- You can also mount directories as read-only by adding `:ro` at the end:

  ```sh
  docker run -v /host/path:/container/path:ro <image>
  ```

### Volumes

Volumes are a better option for data persistence, especially in production environments. Unlike bind mounts, Docker
manages volumes, and they are stored in a Docker-controlled part of the host filesystem (usually
`/var/lib/docker/volumes/`). Volumes are useful because they are portable, easier to back up, and can be shared across
containers. Docker automatically handles the lifecycle of volumes, making them easier to work with at scale.

- Example:

  ```sh
  docker volume create my-volume
  docker run -v my-volume:/container/path <image>
  ```

  In this example, the named volume `my-volume` is created and mounted into the container at `/container/path`.

- Alternatively, you can use the `--mount` option, which provides more flexibility and clarity in the syntax:

  ```sh
  docker run --mount type=volume,source=my-volume,target=/container/path <image>
  ```

Volumes are typically preferred for production environments because they are decoupled from the host’s filesystem and
offer better portability and management options.

### When to Use Bind Mounts vs. Volumes

- **Bind Mounts**: Ideal for **development** when you need a live connection between your local files and the container,
  allowing immediate reflection of file changes.
- **Volumes**: Preferred in **production** environments for **data persistence**, as Docker manages their lifecycle and
  they are more portable and easier to back up.

### Networking

By default, Docker containers run in isolated networks. Each container gets its own virtual network interface, and
containers can communicate with each other or the host using Docker's networking capabilities. You can configure
different types of networks based on your needs:

- **Bridge Network**: The default network for containers, allowing them to communicate with each other via container
  names.
- **Host Network**: Allows the container to use the host’s network directly (useful for performance but removes network
  isolation).
- **Custom Networks**: You can create custom networks for more advanced setups, allowing for better control over how
  containers communicate with each other and external systems.

Example of creating and using a custom network:

```sh
docker network create my-network
docker run --network=my-network <image>
```

---

## Docker vs. Docker Compose

As applications grow more complex, especially when building **multi-container applications** (e.g., a web server,
database, and caching service), managing individual containers with basic Docker commands can become **cumbersome** and
difficult to orchestrate. For example, manually running and linking multiple containers would require several
`docker run` commands, as well as managing network configurations and environment variables.

**Docker Compose** simplifies this process by allowing you to define, build, and manage multiple containers in a single
configuration file, making it easier to orchestrate complex environments.

Docker Compose uses a `docker-compose.yaml` (or `docker-compose.yml`) file to describe the services, networks, and
volumes that your application needs, and can start everything with one simple command.

### Key Differences:

- **Docker**: Good for running single containers or manually controlling individual containers. For simple, one-off
  tasks, Docker’s CLI commands (`docker run`, `docker build`, etc.) are sufficient.
- **Docker Compose**: Ideal for **multi-container** applications where several services need to work together. Instead
  of manually starting, stopping, and linking multiple containers, Docker Compose can handle all of this through a YAML
  configuration file.

### Docker Compose Workflow

With Docker Compose, you write a **YAML configuration file** (`docker-compose.yml`) that defines all the services (
containers) your application requires, including the build context, network configuration, environment variables, and
dependencies.

For example, a `docker-compose.yml` file for a web app with a database might look like this:

```yaml
version: "3"
services:
  web:
    build: .
    ports:
      - "5000:5000"
    volumes:
      - .:/app
    depends_on:
      - db
  db:
    image: postgres
    environment:
      POSTGRES_USER: user
      POSTGRES_PASSWORD: password
```

This file defines two services:

- **web**: A container built from the current directory that runs a web application, exposes port `5000`, and has a
  volume mapping the host directory to `/app` inside the container.
- **db**: A container running a Postgres database image, with environment variables for user and password.

With this single YAML file, you can now manage both containers as a group, simplifying the process of launching and
linking services.

### Key Docker Compose Commands

1. `docker compose up`

   ```sh
   docker compose up
   ```

    - **Builds** missing images and **starts all containers** defined in the `docker-compose.yml` file.
    - If any images are missing (for example, if the image has not been built locally), Docker Compose will
      automatically build them using the instructions provided in the `build` section of the YAML file.
    - **Starts the services**: All containers (services) defined in the YAML file will be started together. Docker
      Compose also ensures the correct startup order of services (for instance, the `web` service may depend on the `db`
      service being up first).
    - Example: Running `docker compose up` in the example file would start the web app and PostgreSQL database in
      separate containers, with the web app automatically linked to the database.

   Additional options:
    - `-d`: Runs the services in **detached mode** (in the background), similar to running `docker run -d`.
      ```sh
      docker compose up -d
      ```

    - `--build`: Forces a rebuild of the images, even if they already exist. This is useful if you’ve made changes to
      the `Dockerfile` or source code.
      ```sh
      docker compose up --build
      ```

2. `docker compose down`

   ```sh
   docker compose down
   ```

    - **Stops** all running containers started by `docker compose up` and **removes them**, along with any networks or
      volumes defined by the `docker-compose.yml` file.
    - This ensures a clean shutdown of all containers, clearing up system resources without leaving unused containers or
      networks behind.
    - Example: Running `docker compose down` will stop both the web and database containers, removing any associated
      networks.

   Additional options:
    - `-v`: Removes the **volumes** associated with the containers. This is helpful if you want to clear out any data
      stored in volumes and start with a fresh environment the next time.
      ```sh
      docker compose down -v
      ```

### When to Use Docker Compose

- **Multi-container applications**: If your application requires multiple services (like a web server, database, cache,
  etc.), Docker Compose makes it easy to manage and link those services together.
- **Environment configuration**: Docker Compose allows you to define environment variables, volumes, and network
  configurations in a structured YAML file, ensuring consistency across different environments (development, testing,
  production).
- **Simplified management**: With one command (`docker compose up`), you can start all your services, and with another (
  `docker compose down`), you can shut them all down, making it easier to manage complex setups.

### Benefits of Docker Compose

- **Orchestration**: Docker Compose orchestrates the startup order of containers, ensuring that dependencies (like a
  database) are started before other services that depend on them (like a web app).
- **Reproducibility**: The `docker-compose.yml` file ensures your environment is consistent, making it easy to share and
  replicate setups across different machines and teams.
- **Networking**: Docker Compose automatically sets up networks, allowing containers to communicate with each other
  using simple container names. You don’t need to manually configure links or network settings.
- **Volume Management**: Compose helps manage persistent data by defining volumes in the YAML file, ensuring data
  persists even when containers are restarted or removed.

---

## Local vs. Remote Development Using Docker

Docker provides a consistent environment for developing, testing, and running applications, whether locally on a
developer’s machine or remotely on a production server. This portability ensures that what works in your local
environment will work identically in a remote or production environment.

### Local Development with Docker

When working **locally**, Docker helps developers create isolated, reproducible environments that eliminate common
issues such as dependency conflicts or the need to install software globally. This results in a more efficient and
consistent development experience.

- **Isolated, encapsulated, and reproducible environments**: Docker allows you to create a containerized environment for
  each project, completely isolated from other applications running on your machine. This means that each project can
  have its own versions of programming languages, libraries, and tools without interfering with other projects.
    - Example: You could have one container running a Python 3.9 app and another running a Python 2.7 app, without
      causing any conflicts or requiring changes to your local system.

- **No dependency or software clashes**: Because Docker containers are isolated from each other and the host system, you
  don’t need to worry about version mismatches or conflicting dependencies between different projects. All necessary
  dependencies are bundled within the container, ensuring that your local environment remains clean and consistent.
    - Example: Instead of installing Node.js or PostgreSQL globally, Docker containers can include those dependencies,
      ensuring that each project gets the correct versions without affecting your system.

- **Faster onboarding and setup**: Using Docker in local development makes it easy for new team members to get started.
  They only need Docker and Docker Compose installed to run the project’s containerized environment, eliminating complex
  setup processes or manual installation of dependencies.

### Remote Host / Production with Docker

When moving to **remote environments**, such as staging, testing, or production servers, Docker simplifies deployment by
ensuring that the same containers used in local development can be used in these environments. Docker provides a
seamless transition from local to remote environments, minimizing deployment risks and reducing the "it works on my
machine" problem.

- **What worked locally will work on a remote environment as well**: Docker containers ensure that your code runs in the
  same environment regardless of where it's executed. Since the environment and dependencies are encapsulated within the
  container, there are no surprises when moving from local development to production. This consistency simplifies
  debugging and reduces the risk of errors due to environmental differences.
    - Example: If a container running locally includes Node.js version 14 and MongoDB, the same container will run with
      the same dependencies in production, regardless of the host server’s OS or installed software.

- **Easy updates and rollbacks**: Docker makes updating production environments much simpler. Instead of manually
  updating software or dependencies on a remote server, you can simply **replace the existing container** with an
  updated one. If an issue arises, rolling back to a previous version is as simple as redeploying the old container.
    - Example: If you need to update your web app, you can build a new image and deploy it by stopping the old container
      and starting a new one. Rolling back is just as easy by restarting the previous container version.

- **Scalability and consistency**: In production, Docker containers are highly scalable. Tools like Docker Swarm or
  Kubernetes can be used to orchestrate and scale Docker containers across multiple servers, ensuring high availability
  and load balancing. Since Docker containers behave consistently across all environments, scaling becomes a
  straightforward process of replicating containers.

---

## Deployment Considerations

When moving from development to production deployment with Docker, there are several important considerations to ensure
your application is performant, secure, and scalable. While Docker simplifies many aspects of deployment, certain
strategies and best practices should be followed to optimize for production environments.

### 1. Replace Bind Mounts with **Volumes** or `COPY`

- **Bind Mounts** are often used during local development to link a directory from your host system to a directory
  inside the container, allowing live changes to reflect immediately in the container. However, this approach is **not
  ideal for production** because it depends on the host’s filesystem and could lead to **security risks** and
  **inconsistent behavior** in different environments.

- **Volumes**: In production, use **Docker volumes** for persistent data storage. Volumes are managed by Docker and are
  more secure, performant, and portable across different environments (local, staging, production). Volumes allow data
  to be decoupled from the container lifecycle, ensuring persistence even if the container is restarted or replaced.
    - Example: A database container might use a volume to store its data so that it persists across container restarts.

  ```sh
  docker run -v my_volume:/data my_container
  ```

- **`COPY`**: For production images, avoid bind mounts for copying application code. Instead, use the `COPY` instruction
  in the `Dockerfile` to copy the necessary files into the image during the build process. This ensures that the
  application code is bundled inside the container and not dependent on the host filesystem.
    - Example in `Dockerfile`:
      ```Dockerfile
      COPY . /app
      ```

  This method makes your Docker image self-contained, reducing the chances of discrepancies between the local and
  production environments.

### 2. Multiple Containers Might Need Multiple Hosts

- For **larger, more complex applications**, a single server might not be enough to run all your containers efficiently.
  As your application scales, you may need to distribute containers across multiple hosts for **load balancing** and 
 **high availability**.

    - **Multi-container, multi-host setups**: For applications that rely on several services (e.g., a frontend, backend,
      database, and caching layer), you might want to run containers on different servers to avoid overloading a single
      host or to improve redundancy.
    - **Orchestration tools**: To manage multiple hosts and containers at scale, consider using orchestration tools like
      **Docker Swarm** or **Kubernetes**. These platforms help you schedule, manage, and scale containers across
      clusters of hosts, ensuring that containers are distributed efficiently.

    - Example: A web app might have its web server running on one host, its database on another, and its cache (e.g.,
      Redis) on a third host. Docker Swarm or Kubernetes can automatically distribute and manage these services.

### 3. Multi-stage Builds

- **Multi-stage builds** are a Docker feature that allows you to create lean, optimized production images by **splitting
  the build process into multiple stages**. This technique helps you reduce image size, which leads to **faster
  deployments** and **smaller attack surfaces** in production.

    - During development, you may need a lot of dependencies, build tools, and debugging features that are unnecessary
      in production. With multi-stage builds, you can **separate the build environment** from the runtime environment.

    - Example: In a multi-stage build for a Go application, the first stage compiles the code with all necessary build
      dependencies, and the second stage creates a lightweight image with only the compiled binary.

      ```Dockerfile
      # First stage: Build the application
      FROM golang:1.16 as builder
      WORKDIR /app
      COPY . .
      RUN go build -o myapp
  
      # Second stage: Copy the built binary to a minimal image
      FROM alpine:latest
      WORKDIR /root/
      COPY --from=builder /app/myapp .
      CMD ["./myapp"]
      ```

  This results in a smaller final image (based on `alpine`, a lightweight base image) without the build dependencies,
  leading to faster startup and reduced image bloat.

### Control vs. Ease of Use: Deployment Options

When deploying Docker containers to production, you must decide between two approaches: **managing your own
infrastructure** or using a **managed Docker hosting service**. Each has its trade-offs in terms of control, complexity,
and ease of use.

#### 1. Provisioning Your Own Server

Provisioning your own server for Docker gives you **full control** over the infrastructure, but it comes with additional
responsibilities:

- **Complete control**: By setting up and managing your own server (e.g., using a cloud service like AWS EC2, Google
  Compute Engine, or a physical server), you have complete control over how Docker is installed, configured, and run.

    - You can configure resource limits, optimize performance, set up security policies, and manage networking exactly
      how you need.

    - You’re responsible for **monitoring**, **scaling**, and **maintaining** the server, including managing security
      updates, backups, and system monitoring.

    - **Greater flexibility**: You can configure the server for special requirements, such as custom networking or
      firewall rules, that may not be available in a managed service.

  However, this approach requires a deeper understanding of system administration and can be time-consuming to manage at
  scale.

    - Example: Running a **self-hosted Docker Swarm** or **Kubernetes cluster** on your own servers.

#### 2. Using a Managed Service

A managed Docker hosting service offloads much of the complexity of managing the infrastructure, providing ease of use
at the cost of some control. Managed services handle most of the infrastructure work for you, making them ideal if you
want to focus more on application development than on managing servers.

- **Less control, easier to manage**: Managed services, such as **AWS Elastic Container Service (ECS)**, **Google
  Kubernetes Engine (GKE)**, **Azure Kubernetes Service (AKS)**, or **DigitalOcean App Platform**, handle the 
 **provisioning**, **scaling**, and **monitoring** of Docker containers. This reduces the need for deep operational
  knowledge.

    - **Automatic scaling**: These platforms automatically scale your application as traffic increases and ensure that
      containers are balanced across multiple hosts.

    - **Built-in monitoring and logging**: Managed services provide monitoring tools to track the health of your
      containers and alert you to any issues. Many platforms offer integrations with monitoring and logging services,
      such as **AWS CloudWatch**, **Google Cloud Monitoring**, or **Datadog**.

    - **Simplified networking**: These platforms automatically handle networking and load balancing between your
      containers and services.

    - **Security updates and patches**: Managed services take care of security updates for the underlying
      infrastructure, reducing the burden of managing server vulnerabilities.

  The trade-off here is that you may have less flexibility in certain configurations, and you're relying on a
  third-party provider for availability and control over the underlying infrastructure.

    - Example: Deploying a **containerized web app** on AWS ECS, which handles resource scaling, security updates, and
      load balancing automatically.
