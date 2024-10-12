# Managing Application Data in Kubernetes

You should verify that your cluster is running with the following command before proceeding:

```shell
minikube status
```

You should see output similar to this:

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

## 
