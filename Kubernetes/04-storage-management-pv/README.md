# Managing Application Data in Kubernetes: Persistent Volumes

You should verify that your cluster is running with the following command before proceeding:

```shell
minikube status
```

You should see output similar to this, if not run `minikube start`:

```text
minikube
type: Control Plane
host: Running
kubelet: Running
apiserver: Running
kubeconfig: Configured
```

[Kubernetes Tools](../TOOLS.md)

## Persistent Volumes

In Kubernetes, storage is a critical component, especially for stateful applications that require data persistence even
after a pod or node is deleted or rescheduled. Persistent Volumes (PVs) offer a solution to overcome the limitations of
regular volumes, providing a more durable and flexible storage system within a Kubernetes cluster.

**Challenges with Regular Volumes**:

- **Ephemeral Nature**: Regular volumes such as `emptyDir` or volumes bound directly to a pod (like `hostPath`) are
  ephemeral. When the pod is deleted or recreated (e.g., due to scaling, upgrades, or node failures), the volume and any
  data stored within it are destroyed as well. This ephemeral behavior is fine for temporary data but becomes
  problematic for applications requiring long-term storage, such as databases or logging systems.
- **Scaling and Data Accessibility**: When you scale a deployment, new pods do not automatically inherit data from
  previous instances, as regular volumes are specific to individual pods. This leads to data fragmentation and loss of
  continuity, as different instances cannot access shared data unless specifically configured with a persistent
  mechanism.
- **Multi-node Clusters and HostPath Limitations**: In multi-node Kubernetes clusters, volumes like `hostPath`, which
  use the local file system of a node, do not work across nodes. If a pod is rescheduled to a different node, it loses
  access to the data stored in the previous node’s local filesystem, since data is not automatically replicated across
  nodes.

**Pod and Node Independence**:

- **Regular Volumes and Node Binding**: Volumes like `emptyDir` are created when a pod is scheduled and exist only for
  the lifetime of the pod on a specific node. `hostPath` mounts the host’s local filesystem but is also bound to the
  lifecycle of the pod and node. As a result, when a pod or its host node is terminated, any data stored in these
  volumes is either lost or becomes inaccessible.
- **Problem for Stateful Applications**: Applications that require persistent storage, such as databases (MySQL,
  Postgres), messaging queues (RabbitMQ, Kafka), or file storage services (NFS, GlusterFS), cannot rely on regular
  volumes because data loss or inaccessibility upon pod rescheduling is unacceptable.

**Persistent Volumes (PVs)**:

- **Persistent Storage Solution**: Kubernetes introduces **Persistent Volumes (PVs)** to decouple the storage lifecycle
  from the pod lifecycle. A Persistent Volume is a storage resource in the cluster that exists independently of any
  specific pod. This means that even if a pod is terminated or rescheduled to a different node, the data in the PV
  remains intact and can be reattached to a new pod.
- **Data Durability and Survivability**: PVs ensure that data can survive beyond the lifespan of a pod, making it a
  critical feature for stateful applications. Whether a pod is rescheduled due to scaling or node failure, the
  underlying storage remains persistent, ensuring no data is lost during transitions.
- **Node-Agnostic Storage**: Unlike `hostPath`, PVs are not bound to specific nodes. This means that pods running on any
  node in the cluster can access the data stored in the PV, depending on the underlying storage backend (e.g., networked
  storage, cloud block storage). This solves the issue of data availability in multi-node environments.
- **Dynamic and Static Provisioning**: Kubernetes offers two ways to manage persistent storage:
    - **Static Provisioning**: In this model, a cluster administrator manually provisions PVs, which are then claimed by
      pods via **Persistent Volume Claims (PVCs)**.
    - **Dynamic Provisioning**: Kubernetes can dynamically create PVs on-demand when a PVC is requested by a pod. This
      eliminates the need for cluster admins to pre-provision volumes and allows for more flexibility, especially in
      cloud environments with elastic storage backends.
- **Storage Classes and Flexibility**: PVs can be configured with **Storage Classes**, which define the type and
  characteristics of storage (e.g., SSDs, HDDs, network-attached storage) that can be dynamically provisioned. Storage
  classes allow Kubernetes to integrate with various storage backends, such as AWS EBS, GCE Persistent Disks, NFS, or
  Ceph, providing flexibility in how storage is allocated and managed.

### Benefits of Persistent Volumes:

1. **Pod Resiliency**: PVs enable pods to restart or be rescheduled across nodes without losing their data, making them
   ideal for high-availability setups.
2. **Centralized Data Access**: Applications can access centralized, durable storage regardless of which node they run
   on, allowing for seamless data sharing between pods and state persistence.
3. **Separation of Concerns**: PVs abstract storage management away from individual pod configurations, allowing storage
   to be handled independently by administrators or dynamic provisioning systems.
4. **Cross-platform Compatibility**: PVs can be backed by various storage systems, including local storage,
   network-attached storage, or cloud storage solutions, making them highly adaptable to different infrastructures.

## Modifying our Configuration: Creating a Persistent Volume

Let's create a volume inside our **Deployment** by modifying our `deployment.yaml` file.

```yaml
spec:
  containers:
    - name: storage-demo-app
      ...
      volumeMounts:
          - mountPath: /app/story
            name: storage-app-demo-volume
  volumes:
    - name: storage-app-demo-volume
      emptyDir: {}
```

We've modified our configuration, let's apply the changes:

```shell
pushd app
kubectl apply -f deployment.yaml
popd
```

