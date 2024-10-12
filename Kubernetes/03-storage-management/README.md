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

## Modifying our Configuration

Let's create a volume inside of our **Deployment** by modifying our `deployment.yaml` file.

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