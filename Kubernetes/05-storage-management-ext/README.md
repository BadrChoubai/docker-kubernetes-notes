# Managing Application Data in Kubernetes: Quick Bits

Continuing on from our lesson on Persistent Volumes, let's quickly go over Environment Variables and
`ConfigMap` in Kubernetes.

## Environment Variables

Environment variables are a key part of configuring applications in Kubernetes. They allow you to provide dynamic
configuration to containers without hardcoding values into the application's code. In Kubernetes, environment variables
can be injected into pods at runtime, making it easy to manage settings such as file paths, API keys, or external
services in a more flexible and secure manner.

Kubernetes lets you define these environment variables directly within your **Pod** or **Deployment** YAML file, so they
are accessible to your application. This approach ensures that your application can access dynamic or configurable
values without being tightly coupled to the underlying infrastructure.

### Using Environment Variables

You can set environment variables inside the pod configuration, and the application can read them just like any other
environment variable defined at runtime. Here's an example of how to define and use environment variables in a
Kubernetes deployment.

1. Let's modify our app to use an environment variable somewhere:

    ```javascript
    const filePath = path.join(__dirname, process.env.STORY_FOLDER, 'text.txt');
    ```

2. Let's make sure we introduce that variable in our **Deployment** `config.yaml`:

    ```yaml
    containers:
    - name: storage-demo-app
      env:
        - name: STORY_FOLDER
          value: 'story'
    ```

3. After those changes to our configuration, let's apply them:

   ```shell
   pushd app
   kubectl apply -f deployment.yaml
   popd
   ```

## ConfigMap

A **ConfigMap** in Kubernetes is a way to manage environment variables and other configuration data separate from the
container images. Instead of embedding configuration directly into your application code or deployment files, you can
store it in a ConfigMap. This makes your applications more portable, as they can be reconfigured easily without
modifying the actual application code or rebuilding containers.

With ConfigMaps, you can externalize key-value pairs, making it easier to manage multiple environment-specific
configurations or sensitive data. This decouples your configuration data from your app logic, leading to more manageable
and maintainable Kubernetes setups.

ConfigMaps can be used to pass environment variables into pods, as well as to populate configuration files or
command-line arguments in containers. Here's an example of how to create and use a ConfigMap in your Kubernetes
deployment.

### Using Config Map

By using a ConfigMap, you can separate your configuration data from your application, making it easier to update or
change environment-specific settings without needing to modify the container or deployment definitions themselves.

1. Let's create a file named: `environment.yaml`:

    ```yaml
    apiVersion: v1
    kind: ConfigMap
    metadata:
      name: data-storage-env
    data:
      folder: 'story'
    ```

2. Let's now use this `configMap` and access the values from it inside our configuration:

   ```yaml
    containers:
    - name: storage-demo-app
      env:
        - name: STORY_FOLDER
          valueFrom:
            configMapKeyRef:
              name: data-store-env
              key: folder
   ```

3. After those changes to our configuration, let's apply them:

   ```shell
   pushd app
   kubectl apply -f config-map.yaml
   kubectl apply -f deployment.yaml
   popd
   ```

4. Let's verify the changes:
    - Running `kubectl get configmap` should output something that looks like:

      | NAME             | DATA | AGE |
      |------------------|------|-----|
      | data-storage-env | 1    | 14s |

   
