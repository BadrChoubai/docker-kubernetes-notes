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

Kubernetes has many [types of Volumes](https://kubernetes.io/docs/concepts/storage/volumes/#volume-types) to choose from
this lesson will focus on three:

- **[emptyDir](https://kubernetes.io/docs/concepts/storage/volumes/#emptydir)**
- **[hostPath](https://kubernetes.io/docs/concepts/storage/volumes/#hostpath)**
- **[Container Storage Interface (CSI)](https://kubernetes.io/blog/2019/01/15/container-storage-interface-ga/)**

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