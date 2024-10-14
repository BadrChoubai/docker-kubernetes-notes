# Networking in Kubernetes: Hands-On

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

Also, double-check that no existing resources exist except for the default `Kubernetes: ClusterIP`:

```shell
kubectl get deployments
kubectl get services
```

[Kubernetes Tools](../TOOLS.md)

## Our Application

1. **Multiple APIs**: We are working with three different APIs which we would like to be able to establish connections
   inside our Kubernetes Cluster:
    - **User API**: This API is takes an incoming request to create a new user
    - **Auth API**: This API creates an authentication token for a new user
    - **Tasks API**: This API enables the creating and reading of tasks for a given user

2. **Beginning Architecture**:

   ![App Architecture Diagram](../../../.attachments/Network-project-diagram.png)

3. Our Goals:

   1. Deploy each application into our Kubernetes cluster:

      - [ ] All APIs running in the same Cluster
      - [x] Auth and Users API in same Pod
      - [ ] Tasks API in own Pod
        
   2. Allow communication between Users API and Auth API
   3. Ensure that only the Users API and Tasks API are accessible by an API Client

### Build and Push Our Container Images

```shell
eval $(minikube -p minikube docker-env)

(
  pushd users-api || exit
  docker build -t cc-users-api:1.0 .
  popd || exit
)
  
(
  pushd auth-api || exit
  docker build -t cc-auth-api:1.0 .
  popd || exit
)
```
   
### Deploy Our Users App to the Cluster

```shell
pushd infrastructure/users
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
popd
```

## Evolving the Architecture

**Goal Architecture**:

   ![App Architecture Diagram](../../../.attachments/2nd-Network-project-diagram.png)

```shell
pushd infrastructure/users
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
popd

pushd infrastructure/auth
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
popd
```
