# Deploying to Production

This directory contains project code for the lessons on deploying Docker in
production onto a cloud provider: this course used **AWS**.

## Control vs. Ease of Use: Deployment Options

When deploying Docker containers to production, you must decide between two approaches: **managing your own
infrastructure** or using a **managed Docker hosting service**. Each has its trade-offs in terms of control, complexity,
and ease of use.

### 1. Provisioning Your Own Server

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

### 2. Using a Managed Service

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
