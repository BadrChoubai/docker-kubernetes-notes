# Managing Application Data in Kubernetes

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

## Project Structure

Inside of this directory we have a simple application which reads and writes from the file inside of `story/` inside 
the `app` directory.

```text
app
├── app.js
├── docker-compose.yaml
├── Dockerfile
├── package.json
└── story
    └── text.txt
```

## Deploy the App to Our Cluster

These instructions are not detailed as we've already covered them in the first two lessons, simply run the below script
to scaffold the infrastructure for our application. You can view the configuration for our **Deployment** and
**Service** inside of `/app/deployment.yaml`.

```shell
( # Build our container image
  pushd app
  ./build-image.sh
  popd
)

( 
    pushd app
    kubectl apply -f deployment.yaml
    popd
)
```

- **Checking our Deployment**:

    ```shell
    kubectl get deployments
    ```

    ---

  | NAME                    | READY |
  |:------------------------|:-----:|
  | storage-demo-deployment |  1/1  |

- **Opening our Application**:

    ```shell
    minikube service storage-app-service
    ```

## Our Application's Functionality

Our app has two endpoints exposed:

1. `GET /story`:
    - Reads the content from the file (`text.txt`) inside of the `story` directory
    - Returns the content in a JSON response:

        ```json
        { "story": "" } 
        ```
      
2. `POST /story`:
    - Reads the incoming JSON request body, i.e.:

        ```json
        { "story":  "Once upon a time,..." }
        ```

    - Appends it to the (`text.txt`) inside of the `story` directory

### The Problem

Because we haven't setup any **Volume** in our Kubernetes cluster, if our application
crashes we will lose any of the data that a user generates.

These volume types provide flexibility for managing different kinds of data storage within Kubernetes, depending on use
cases like temporary storage, access to host files, or integration with third-party storage providers.

1. **[emptyDir](https://kubernetes.io/docs/concepts/storage/volumes/#emptydir)**

    - An `emptyDir` volume is initially empty when a pod is created, and it lives as long as the pod is running. It is
      typically used to store temporary data generated during the container's lifecycle. The data in an `emptyDir`
      volume is deleted if the pod is terminated. This type of volume is useful for scenarios like sharing scratch data
      between containers in a pod or caching temporary files.

    - **Key points**:
        - Data is stored on the node’s local storage.
        - Cleared upon pod termination (but not container restarts).
        - Used for sharing temporary or intermediate data between containers in the same pod.

2. **[hostPath](https://kubernetes.io/docs/concepts/storage/volumes/#hostpath)**

    - A `hostPath` volume mounts a file or directory from the host node's filesystem into a pod. This allows the pod to
      access or persist data on the underlying node. However, it introduces concerns with data portability and security,
      since different nodes may have different filesystem structures. It is often used for tasks like logging, exposing the
      host’s storage for certain privileged applications, or sharing configuration files.

    - **Key points**:

        - Binds a specific host directory or file to a pod.
        - Can be used to expose host system functionality or share host files.
        - Limited portability and potential security risks due to tight coupling with host nodes.

3. **[Container Storage Interface (CSI)](https://kubernetes.io/blog/2019/01/15/container-storage-interface-ga/)**

    - The Container Storage Interface (CSI) is a standardized way for Kubernetes to integrate with different storage
      systems. It allows third-party storage providers to expose their storage systems in a consistent manner. CSI
      drivers are developed and deployed independently of Kubernetes, which allows for more flexibility and extensibility in
      managing storage resources like block storage, network-attached storage, or cloud-based solutions.

    - **Key points**:

        - CSI provides a unified interface for storage providers.
        - Allows external storage systems to be integrated with Kubernetes.
        - Enables more advanced storage features like dynamic provisioning, snapshots, and cloning.

## Modifying our Configuration: Adding a Volume

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

## We've Introduced a Different Issue...

In Kubernetes, `emptyDir` volumes are **local to the node** where the pod is running. When you have multiple replicas of
a pod (e.g., in a deployment), each replica gets its **own separate `emptyDir` volume** on the specific node where it is
scheduled. These volumes are not shared across replicas and are tied to the lifecycle of each individual pod.

### Behavior with Multiple Replicas Using `emptyDir`:

1. **Isolation per replica**: Each pod replica gets its own distinct `emptyDir` volume. The data stored in one pod’s
   `emptyDir` is not accessible by other pods, even if they are replicas of the same application.

2. **No replication or persistence**: Data in an `emptyDir` volume is local to the node and is not replicated to other
   nodes. If a pod or node fails, the data in that volume is lost and cannot be recovered by another pod.

3. **Scheduling across nodes**: Kubernetes may schedule replicas across different nodes. Each `emptyDir` volume resides
   on the node where its corresponding pod is running, with no automatic synchronization between the volumes of
   different replicas.

#### Example:

Let’s update our deployment to scale up to three replicas:

```yaml
spec:
  replicas: 3
```

Now we have a deployment with three replicas of a pod using `emptyDir`. Each pod will have:

- Its own `emptyDir` on its assigned node.
- No data sharing between pods.
- Separate temporary storage that is deleted when the pod terminates or the node shuts down.

`emptyDir` is great for cases where each pod requires independent temporary storage, but it is unsuitable for data that
needs to be shared or persisted across replicas.

#### Behavior During Node Crashes:

1. **Data Loss on Crash**: If a node crashes, any `emptyDir` volume stored on that node is lost. The data is not
   persisted or replicated, and no recovery is possible from another node.
2. **No Volume Migration**: `emptyDir` volumes are tied to the node where the pod runs. When a node crashes and
   Kubernetes reschedules the pod to a different node, a new `emptyDir` volume is created on the new node, but the data
   from the original node's `emptyDir` is lost.
3. **Independent Recovery**: Each pod replica that is rescheduled on a new node will have its own fresh, empty
   `emptyDir`. There is no data synchronization between pods, even if they are replicas of the same application. The new
   pod starts with an empty volume, regardless of what was stored in the previous pod's `emptyDir` before the crash.

## Modifying our Configuration: Using a Different Volume

Clearly, the behavior of `emptyDir` isn't what we'd expect for our application. Let's change our **Volume** to use the
`hostPath` type instead:

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
      hostPath:
        path: /data
        type: DirectoryOrCreate
```

We've modified our configuration, let's apply the changes:

```shell
pushd app
kubectl apply -f deployment.yaml
popd
```

## What Changed?

In Kubernetes, `hostPath` volumes mount a file or directory from the **host node's filesystem** into the pod. Unlike
`emptyDir`, which is ephemeral, `hostPath` allows pods to access files or directories on the underlying node, making it
suitable for scenarios where pods need to interact with the node's filesystem directly. However, the data is **tied to
the specific node** where the pod is scheduled, and sharing data across nodes requires extra configuration.

### Behavior with Multiple Replicas Using `hostPath`:

1. **Shared access to host files**: Each pod replica can mount the same host directory or file using `hostPath`, but
   since this directory is specific to the node, only replicas scheduled on the same node can access the same data. If
   replicas are spread across multiple nodes, each replica accesses its own host's directory, resulting in potentially
   different data across replicas.

2. **No replication or portability**: The contents of a `hostPath` volume are local to the node. Data is not
   automatically synchronized or replicated between nodes. If a pod or node fails, the data stored in the `hostPath`
   volume on that node is not accessible by pods scheduled on other nodes.

3. **Scheduling across nodes**: If Kubernetes schedules pod replicas across different nodes, each replica will mount the
   `hostPath` on its respective node. There is no inherent mechanism to synchronize or share data between different
   nodes’ filesystems.

**With our three replicas of a pod using `hostPath`, each pod will**:

- Access the specified directory or file on the host node where it is scheduled.
- Have potential differences in data if replicas are spread across multiple nodes, as each pod can only access the
  `hostPath` on its own node.
- Maintain data persistence as long as the node and the host directory are available.

`hostPath` is useful for scenarios where pods need access to specific files or directories on the host node, but it is
not suitable for portable or replicated data across nodes.

#### Behavior During Node Crashes:

1. **Data Persistence on Host**: Unlike `emptyDir`, data stored in a `hostPath` volume can persist on the node after a
   pod terminates. However, if the node crashes, the data becomes inaccessible until the node is restored. The data
   itself remains on the host's filesystem if the node can recover.

2. **No Volume Migration**: `hostPath` volumes are tied to the node where the pod runs. If a node crashes and Kubernetes
   reschedules the pod to a different node, the pod will mount the `hostPath` directory on the new node, which may be
   empty or different from the data on the crashed node. The data on the original node remains on its filesystem but is
   not automatically migrated.

3. **Node-Dependent Data**: Since `hostPath` volumes are tied to the specific host, the pod rescheduled on a different
   node will have its own host directory. If the application relies on shared data across replicas or nodes, additional
   mechanisms (like a networked file system) are needed to ensure data consistency.

## Persistent Volumes

...
