# Managing Application Data in Kubernetes: Persistent Volumes

Continuing on from the last lesson, introducing Volumes. This is a quick guide to setting up and using Persistent Volumes
inside of Kubernetes.

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

### Storage Class

A **Storage Class** in Kubernetes defines how storage should be provisioned within the cluster. It acts as a blueprint
that specifies the type of storage, its performance characteristics, and the method of provisioning that should be used
when a PVC requests storage. Storage Classes are highly flexible, allowing administrators to define multiple classes of
storage based on workload requirements, cost considerations, and performance needs.

**Key Concepts of Storage Classes**:

- **Provisioner**: The provisioner is a component in Kubernetes that knows how to create storage based on the
  specifications in the Storage Class. For example, in cloud environments, there are provisioners for AWS EBS, Google
  Cloud Persistent Disk, and Azure Disk. In on-premises environments, you might use provisioners for NFS, Ceph, or
  GlusterFS.
- **Parameters**: Storage Classes can define specific parameters, such as disk type (e.g., SSD or HDD), replication
  factor, and IOPS (input/output operations per second). These parameters tailor the storage to meet specific
  performance and durability needs.
- **Reclaim Policy**: When a PVC is deleted, the associated Persistent Volume can either be retained or deleted,
  depending on the reclaim policy specified in the Storage Class. This is useful for deciding what happens to data after
  it’s no longer in use by the application.
    - **Retain**: The PV is not deleted, allowing manual intervention to retain data.
    - **Delete**: The PV and the data it contains are automatically deleted when the PVC is deleted.
- **Binding Mode**: Determines when a PV should be bound to a PVC:
    - **Immediate Binding**: The PV is bound as soon as the PVC is created.
    - **WaitForFirstConsumer**: Binding is delayed until a pod using the PVC is scheduled, ensuring that the PV is
      allocated in the same zone as the pod.
- **Volume Expansion**: Some Storage Classes support volume expansion, allowing PVCs to request more storage without
  downtime. This is particularly useful for growing applications, such as databases that may require more space over
  time.

### Benefits of Using Storage Classes:

- **Customization**: Allows you to define multiple storage tiers, such as high-performance SSD-backed storage for
  critical workloads and cost-effective HDD storage for archival purposes.
- **Automation**: Enables dynamic provisioning, reducing the manual overhead for cluster administrators and ensuring
  that storage is allocated automatically based on the application’s needs.
- **Efficient Resource Allocation**: Delaying volume binding until a pod is scheduled ensures that storage resources are
  efficiently used, especially in multi-zone or multi-node clusters.

## Introduction to Persistent Volumes

1. Let's create a new file to define our persistent volume: `host-pv.yaml`:

   ```yaml
   apiVersion: v1
   kind: PersistentVolume
   metadata:
     name: host-pv
   spec:
     capacity:
       storage: 1Gi
     volumeMode: Filesystem
     storageClassName: standard
     accessModes:
       - ReadWriteOnce
     hostPath:
       path: /data
       type: DirectoryOrCreate
   ```

   > This defines a Persistent Volume which can now be used by any **Pod** in our Deployment, but we'll
   > now need to create a **Claim**

2. Let's create a new file to define our **Claim**: `host-pvc.yaml`:

   ```yaml
   apiVersion: v1
   kind: PersistentVolumeClaim
   metadata:
     name: host-pvc
   spec:
     volumeName: host-pv
     accessModes:
       - ReadWriteOnce
     storageClassName: standard
     resources:
       requests:
         storage: 1Gi
   ```

   > We still haven't configured our connection to a **Pod**

3. Connecting our **Pod** to the **Claim**:

   ```yaml
   # deployment.yaml
    spec:
    containers:
        - name: storage-demo-app
          ...
          volumeMounts:
            - mountPath: /app/story
              name: storage-app-demo-volume
    volumes:
      - name: storage-app-demo-volume
        persistentVolumeClaim:
          claimName: host-pvc 
   ```

4. After these additions and changes to our configuration, let's apply them:

   ```shell
   pushd app
   kubectl apply -f host-pv.yaml
   kubectl apply -f host-pvc.yaml
   kubectl apply -f deployment.yaml
   popd
   ```

   ---

   Running `kubectl get pv`, you should see a similar output:

   | NAME    | CAPACITY | CLAIM                | STORAGECLASS |
   |---------|----------|----------------------|--------------|
   | host-pv | 1Gi      | default/host-pvc     | standard     |
