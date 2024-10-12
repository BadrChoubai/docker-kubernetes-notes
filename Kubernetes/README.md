# Kubernetes

Kubernetes introduces automation to the deployment, scaling, and management of containerized applications, significantly
reducing the chance of error while improving the overall efficiency of managing large-scale container environments.

The following sections describe some problems with multi-container deployments and the solution provided by Kubernetes 
to overcoming them:

1. [Handling Container Crashes and Replacements](#handling-container-crashes-and-replacements)
2. [Scaling Application Resources during High-Traffic Usage](#scaling-application-resources-during-high-traffic-usage)
3. [Incoming Traffic Should Be Properly Load-Balanced](#incoming-traffic-should-be-properly-load-balanced)

**Core Concepts**:

- [Kubernetes Architecture](./ARCHITECTURE.md)
- [Kubernetes Objects](./Objects.md)

---

## Handling Container Crashes and Replacements

Containers are lightweight and ephemeral in nature. They can fail due to hardware issues, software bugs, or crashes
related to external dependencies. If a container crashes or shuts down, manually restarting it can introduce delays and
risks in maintaining the availability of the application. Some reasons for container failure might include:

- **Resource Constraints**: A container may run out of CPU, memory, or disk space.
- **Networking Failures**: A network issue may cause a container to lose communication with other services, causing it
  to shut down.
- **Code Failures**: A bug in the code might cause the container to crash or become unresponsive.

Kubernetes solves this problem with self-healing capabilities, specifically:

- **Health Checks**: Kubernetes allows you to define liveness and readiness probes. Liveness probes periodically check
  whether a container is running, and readiness probes check if the container is ready to handle traffic. If a probe
  fails, Kubernetes can take action, such as restarting the container.

- **Auto-Restart**: If a container crashes or fails a health check, Kubernetes automatically restarts it to ensure the
  service continues running without manual intervention.

- **Rescheduling**: If a node (a worker machine) in the Kubernetes cluster fails, Kubernetes reschedules the container
  on a different healthy node, ensuring continuous availability of the service.

---

## Scaling Application Resources During High Traffic Usage

In a dynamic production environment, application usage can fluctuate. For example, during peak hours or promotional
events, you may need to scale up your application to handle the surge in traffic. Manually scaling your containers in
response to varying workloads would require constant monitoring and intervention. Some issues include:

- **Predicting Traffic Spikes**: Manually predicting traffic and scaling ahead of time is difficult. Scaling too late
  can cause performance issues or downtime, while scaling too early wastes resources.

- **Horizontal Scaling Complexity**: Managing horizontal scaling, where multiple instances of a container are deployed
  to handle increased load, requires balancing traffic between containers and ensuring that they can all access shared
  resources.

Kubernetes offers automated scaling solutions:

- **Horizontal Pod Autoscaling (HPA)**: Kubernetes can automatically adjust the number of container instances (pods)
  based on CPU utilization or other custom metrics like memory usage or network traffic. It ensures that the right
  number of containers is running to handle current traffic.

- **Vertical Pod Autoscaling**: Kubernetes can also automatically adjust the resources allocated to a container (e.g.
  CPU or memory) based on real-time usage.

- **Cluster Autoscaling**: If the workload increases beyond the capacity of the current cluster nodes, Kubernetes can
  automatically provision new nodes to handle the increased load, allowing seamless scalability.

---

## Incoming Traffic Should Be Properly Load-Balanced

In a multi-container setup, directing traffic to the correct container instances is crucial for performance, redundancy,
and fault tolerance. Manually managing the distribution of traffic across containers has several challenges:

- **Traffic Overload on Specific Instances**: Without proper load balancing, certain instances of a container might get
  overloaded with traffic, while others remain idle. This can lead to performance bottlenecks and decreased reliability.

- **Handling Container Failures**: If a container fails, incoming traffic must be rerouted to healthy instances to avoid
  downtime or errors.

- **Service Discovery**: As new containers are spun up, other services need a way to discover them and direct traffic
  accordingly. Manually managing this service discovery would be cumbersome.

Kubernetes introduces a built-in **service** resource and various load-balancing mechanisms to manage traffic:

- **Service**: Kubernetes services define a stable IP and DNS name for a set of pods (containers). As containers are
  added or removed, the service ensures that traffic is distributed evenly across available containers, without the need
  for manual intervention.

- **Load Balancers**: Kubernetes supports different types of load balancers:
    - **ClusterIP**: The default type, which routes traffic within the cluster to the right pods.
    - **NodePort**: Exposes the service on each node's IP, allowing external traffic to reach the service.
    - **External Load Balancer**: In cloud environments (e.g. AWS, GCP), Kubernetes can automatically provision
      external load balancers to distribute traffic to the service across multiple nodes.

- **Ingress Controllers**: Kubernetes also supports Ingress resources, which manage external access to the cluster
  (e.g. HTTP/S traffic). An ingress controller handles routing rules and SSL termination, providing more advanced
  load-balancing and traffic-routing features.
