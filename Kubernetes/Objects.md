# Objects in Kubernetes

Kubernetes objects represent the desired state of your cluster. These objects are used to define what containerized
applications are running, which nodes they run on, and what resources and policies are applied to those applications.
These objects are managed through the Kubernetes API and include Pods, Deployments, Services, and other resources.

**Sections**:

- [Pod](#pod)
- [Deployment](#deployment)
- [Service](#service)

## Pod

A **Pod** is the smallest and most basic deployable object in Kubernetes. It represents a single instance of a running
process in your cluster. Pods are designed to be ephemeral, meaning they can be created, destroyed, or replaced based on
the current needs of the application or cluster. Each Pod is tightly coupled to a specific node but is not guaranteed to
remain on that node after restarts or failures.

**Key Features**:

1. **Contains one or more containers**:
   - A Pod can run a single container (most common) or multiple containers that are tightly coupled and need to share
     resources. These containers inside the same Pod are managed together and share the same network namespace and
     storage.
   - For example, one container could run an application, and another could manage logging or monitoring. This approach
     is known as the **sidecar pattern**.

2. **Shared resources for all containers within a Pod**:
   - All containers within a Pod share certain resources such as:
      - **Storage volumes**: These can be ephemeral or persistent (e.g., volumes backed by persistent storage like NFS
        or cloud storage).
      - **Network**: All containers in a Pod share the same IP address and network port space. They can communicate
        with each other using `localhost`.
      - **Environment variables**: These can be configured and shared among the containers to provide configuration
        values.

3. **Pod lifecycle**:
   - Pods are inherently **ephemeral** and are designed to be created and destroyed as needed. If a Pod fails,
     Kubernetes may create a new Pod (potentially on a different node).
   - When managed by a higher-level controller (like a Deployment), Pods can be replaced automatically based on the
     desired state of the system.
   - Pods can be in different phases: Pending, Running, Succeeded, Failed, and Unknown.

4. **Cluster-internal IP address**:
   - Each Pod is assigned a unique IP address within the cluster's internal network. This IP is shared by all
     containers in the Pod, meaning they can communicate using `localhost` on the same ports.
   - Other Pods communicate with each other via the Pod IP, but this IP is dynamic. If a Pod is recreated, its IP
     address changes, which is why Services are often used to provide stable network addresses.

5. **Ephemeral by design**:
   - Pods are not designed for long-term persistence. Their state is meant to be transient, which is why external
     resources like **Persistent Volumes** (PV) and **Persistent Volume Claims** (PVC) are used when you need to retain
     data beyond the lifespan of a Pod.

## Deployment

A **Deployment** is a higher-level abstraction in Kubernetes designed to manage the lifecycle of **Pods**. It ensures
that a set of identical Pods (known as replicas) are running, and manages the creation, deletion, and updates of these
Pods. Deployments offer powerful mechanisms for maintaining the desired state of an application, allowing for features
like scaling, rolling updates, and rollbacks.

**Key Features**:

1. **Manages a set of Pods**:
   - A Deployment ensures that a specified number of **replicas** (identical Pods) are running at all times. If any Pod
     crashes or is deleted, the Deployment controller will create new Pods to replace them, ensuring the application
     remains available.
   - You define the desired state (how many replicas you want, which containers and versions to run), and Kubernetes
     continuously works to ensure the actual state matches the desired state.

2. **Declarative updates and rollbacks**:
   - Deployments support **rolling updates**, allowing for seamless updates of Pods without downtime. Kubernetes will
     gradually replace old Pods with new ones while keeping a minimum number of replicas running during the update
     process. This ensures that the application stays available during the update.
   - If something goes wrong during an update, the Deployment can be **rolled back** to a previous version. Kubernetes
     stores the history of Deployment changes, making it easy to revert to a stable state if needed.
   - Deployments can also be **paused** to temporarily stop making changes to the Pods while still maintaining the
     current state of the system.

3. **Scaling**:
   - Deployments can be **scaled dynamically** by changing the number of replicas. This can be done manually (e.g., by
     specifying a different number of replicas in the configuration) or automatically with a **Horizontal Pod
     Autoscaler** (HPA).
   - The HPA can monitor resource usage (e.g., CPU or memory) and automatically adjust the number of Pods to meet
     demand. This ensures efficient use of resources while maintaining performance and availability.

4. **Self-healing and fault-tolerance**:
   - Deployments ensure fault-tolerance by automatically replacing unhealthy Pods. If a node fails, or a Pod crashes,
     the Deployment controller will reschedule the Pods on healthy nodes, minimizing application downtime.
   - It also provides a **self-healing** capability by ensuring the correct number of replicas is always running. This
     makes Deployments a resilient option for managing stateless applications.

5. **Stateless by design**:
   - Although Deployments can manage any type of workload, they are typically used for **stateless applications**. In a
     stateless application, Pods do not maintain any local state between restarts. Any state is stored externally (
     e.g., in a database or persistent storage).
   - For stateful applications that require Pods to maintain identity or persistent storage, Kubernetes provides a
     different object called **StatefulSet**.

## Service

A **Service** in Kubernetes is an abstraction that defines a logical set of Pods and a policy by which to access them.
It provides a stable way to expose a group of Pods to other applications or users, whether within the cluster (internal)
or outside of it (external). Because Pods are ephemeral and their IP addresses change when they are recreated or
rescheduled, Services ensure that other components can consistently access the right set of Pods even as their IPs
change.

**Key Features**:

1. **Exposes Pods to the Cluster or Externally**:
    - By default, Pods are assigned dynamic, internal IP addresses that may change if a Pod is rescheduled or replaced.
      Services act as a stable interface for accessing Pods, even if their underlying Pods and IP addresses change.
    - Services provide both internal (within the cluster) and external (outside the cluster) access to Pods.
        - **ClusterIP**: The default type of Service, which makes the Pods accessible only within the Kubernetes
          cluster. Other Pods or Services in the cluster can access the Pods through a stable **ClusterIP**.
        - **NodePort**: Exposes the Service on a static port on each Node's IP, allowing access from outside the
          cluster.
        - **LoadBalancer**: Integrates with cloud providers to expose the Service externally using a cloud-based load
          balancer.

2. **Grouping Pods with a Shared IP Address**:
    - Services use **selectors** to group together Pods that match certain criteria (usually based on labels). This
      allows the Service to direct traffic to the appropriate set of Pods.
    - The Service creates a virtual IP address that remains consistent, and traffic sent to this IP is distributed
      across the group of Pods (usually by round-robin load balancing).
    - The stable IP address is assigned as a **ClusterIP**, and clients within the cluster can use this IP or a DNS
      name (provided by Kubernetes) to access the Pods behind the Service.

3. **Stable Networking via DNS**:
    - Kubernetes provides DNS resolution for Services, which makes them accessible by a domain name instead of just an
      IP address. For example, if a Service is named `my-service`, it can be accessed by
      `my-service.default.svc.cluster.local` within the cluster.
    - This is important because it enables other Pods or microservices to communicate with a Service by name, even if
      the Pods backing the Service are rescheduled or their IPs change.

4. **Service Discovery**:
    - Kubernetes provides automatic **service discovery**. When a Pod is added or removed from a Service (based on label
      selectors), the Service updates its list of available endpoints.
    - This dynamic nature ensures that even if the Pods behind the Service change, the client applications can always
      connect to the correct endpoints without needing to be reconfigured.

5. **Load Balancing**:
    - Kubernetes Services automatically distribute incoming traffic across the set of Pods associated with the Service.
      The load balancing is typically done in a round-robin fashion, although more complex setups can be implemented
      using ingress controllers or external load balancers.
    - This helps distribute traffic evenly, ensuring no single Pod is overwhelmed and can help improve fault tolerance.

6. **Different Types of Services** [Reference](#service-types-reference):
    - **ClusterIP** (default):
        - Exposes the Service on an internal IP in the cluster. This makes the Service only reachable from within the
          cluster.
        - Ideal for internal services such as databases or other components that don’t need to be accessed from outside
          the cluster.
    - **NodePort**:
        - Exposes the Service on each node's IP at a static port. A user can access the Service from outside the cluster
          using `<NodeIP>:<NodePort>`.
        - Typically used for development or when external traffic needs to hit specific nodes.
    - **LoadBalancer**:
        - Exposes the Service externally using a cloud provider's load balancer. It assigns a public IP that can be used
          to route external traffic to the Service.
        - This is commonly used in cloud environments (e.g., AWS, Google Cloud) where Kubernetes can automatically
          create and configure cloud load balancers.
    - **ExternalName**:
        - Maps the Service to the contents of an external DNS name. This allows the Service to provide access to
          services outside the cluster without needing to manually configure Pods.

7. **Endpoints Object**:
    - The **Endpoints** object is closely related to a Service. It stores the actual IP addresses of the Pods that are
      part of the Service.
    - The Service object doesn’t store information about the individual Pods directly; instead, it refers to the
      Endpoints object, which holds the list of Pod IPs. As Pods are added or removed, the Endpoints object is updated
      accordingly.

8. **External Access and Ingress**:
    - Services can work together with **Ingress** resources to provide more sophisticated routing and load balancing. An
      **Ingress** object manages external access to Services within a cluster, typically via HTTP or HTTPS, and can
      provide features like virtual hosting and SSL termination.
    - For example, an Ingress can route traffic to different Services based on the hostname or path.

---

### Service Types Reference:

| Service Type     | Use Case                                                           | External Access                     |
|:-----------------|:-------------------------------------------------------------------|:------------------------------------|
| **ClusterIP**    | Internal access within the cluster only                            | **No**                              |
| **NodePort**     | Expose the Service on a static port on each node                   | **Yes (via `<NodeIP>:<NodePort>`)** |
| **LoadBalancer** | Expose the Service externally via a cloud provider's load balancer | **Yes (via cloud LB)**              |
| **ExternalName** | Map a Service to an external DNS name                              | **No (external redirection)**       |
