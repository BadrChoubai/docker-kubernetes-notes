# Objects in Kubernetes

Kubernetes objects represent the desired state of your cluster. These objects are used to define what containerized
applications are running, which nodes they run on, and what resources and policies are applied to those applications.
These objects are managed through the Kubernetes API and include Pods, Deployments, Services, and other resources.

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