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