# Intermediate Kubernetes

## Setting Up

To get started with Kubernetes, ensure you have the necessary tooling installed. For this course, we’ll be using
`kubectl` to interact with the Kubernetes API and a local cluster using `minikube`.

After you’ve installed `kubectl` and `minikube`, start the local cluster by running:

```shell
minikube start
```

You can verify that your cluster is running with the following command:

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

## Creating a Deployment

### Project Structure

Inside of this directory we have the same application as before inside the `app` directory, with the inclusion
of an `infrastructure` directory where a configuration for a **Service** and **Deployment** can be viewed:

```text
app
├── .dockerignore
├── app.js
├── Dockerfile
└── package.json

infrastructure
├── apply.sh
├── deployment.yaml
├── service.yaml
└── teardown.sh
```

To deploy this application inside Kubernetes, we need to create a Deployment configuration file.

### Deployment Configuration

1. Create the Deployment using the below script:

   ```shell
   pushd infrastructure
   kubectl apply -f deployment.yaml
   popd
   ```

   > Check the status of the Pods with:

   ```shell
   kubectl get pods
   ```

### Exposing the Application

To make our application accessible from outside the cluster, we need to create a Service. Here's the configuration for
our Service:

```yaml
apiVersion: v1
kind: Service
metadata:
   name: backend # Name of the service.
spec:
   selector:
      app: node-app # Selects Pods with the label "app: node-app".
   ports:
      - protocol: TCP # Protocol for the service.
        port: 80 # Port exposed by the service.
        targetPort: 8080 # Port on the Pod that the service forwards to.
   type: LoadBalancer # Exposes the service externally.
```

1. Create the Service with `kubectl`:

   ```shell
   pushd infrastructure
   kubectl apply -f service.yaml
   popd
   ```

2. Verify the Service status:

   ```shell
   kubectl get services
   ```

   You should see output similar to:

   | NAME           | TYPE         | CLUSTER-IP   | EXTERNAL-IP |    PORT(S)     |   AGE |
   |:---------------|:-------------|:-------------|:-----------:|:--------------:|------:|
   | backend        | LoadBalancer | 10.99.23.244 | `<pending>` | 80:31155/TCP   |   63s |

3. To access the application, run:

   ```shell
   minikube service backend
   ```

   This will open your application in the default web browser.

## Updating and Deleting Resources

### Updating Resources

If you need to update your Deployment, modify the `deployment.yaml` file and reapply the changes:

```shell
kubectl apply -f deployment.yaml
```

You can also scale your deployment easily. For example, to scale to 3 replicas:

1. Update `deployment.yaml:spec.replicas`

   ```yaml
   spec:
      replicas: 3
   ```

2. Run the same `kubectl` command as earlier:

   ```shell
   kubectl apply -f deployment.yaml
   ```

   | NAME                                 | READY | STATUS  |
   |--------------------------------------|-------|---------|
   | node-app-deployment-6bd546888c-27s6d | 1/1   | Running |
   | node-app-deployment-6bd546888c-8lktc | 1/1   | Running |
   | node-app-deployment-6bd546888c-tksbz | 1/1   | Running |

### Deleting Resources

To delete your Deployment and Service, use the following commands:

```shell
kubectl delete deployment node-app-deployment
kubectl delete service backend
```

## Understanding Labels and Selectors

Labels and selectors are essential for organizing and managing resources in Kubernetes.

### Using Labels

You can define labels in your Deployment and Service configuration to facilitate resource management:

```yaml
selector:
   matchLabels:
      app: node-app # Matches Pods with this label.
   matchExpressions:
      - { key: key, operator: In, values: [ first, second ] } # More complex matching criteria.
```

### Practical Example

1. To get all resources with a specific label:

   ```shell
   kubectl get pods -l app=node-app
   ```

2. To view the labels of your Pods, run:

   ```shell
   kubectl get pods --show-labels
   ```

This allows you to easily filter and manage your resources based on specific attributes.

## Liveness Probes

Liveness probes are a critical feature in Kubernetes that help ensure the health and availability of your applications.
They allow Kubernetes to detect and respond to situations where an application may be running but is unresponsive or
stuck, ensuring that it can automatically recover.

### What are Liveness Probes?

A liveness probe is a mechanism that Kubernetes uses to periodically check if your application is running as expected.
If the liveness probe fails, Kubernetes will automatically restart the container, helping to maintain the application's
availability.

### Configuring Liveness Probes

To configure a liveness probe, you can add it to your Deployment configuration at
`spec.template.spec.containers[0].livenessProbe`. Here’s an example of how to set up a
simple HTTP liveness probe:

```yaml
containers:
  - name: node-app-v1
    image: kub-first-app:1.0
    ports:
      - containerPort: 8080
    livenessProbe:
      httpGet:
        path: /health # Endpoint to check for liveness
        port: 8080
      initialDelaySeconds: 5 # Delay before the probe starts
      periodSeconds: 10 # How often to perform the probe
```

### Steps to Add a Liveness Probe

1. **Define the Endpoint**: Ensure your application has a health check endpoint, such as `/health`, which returns a
   success response when the application is healthy.

2. **Modify the Deployment Configuration**: Add the `livenessProbe` section as shown above in your `deployment.yaml`
   file.

3. **Apply the Changes**:

   ```shell
   pushd infrastructure
   kubectl apply -f deployment.yaml
   popd
   ```

4. **Monitor the Probes**: You can check the status of your Pods and see if the liveness probe is functioning correctly:

   ```shell
   kubectl get pods
   ```

### Conclusion

By implementing liveness probes, you can ensure that Kubernetes automatically manages the health of your applications.
This feature helps to maintain high availability and resilience in your deployments, as it allows for quick recovery
from unresponsive states.
