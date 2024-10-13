# Managing Application Data in Kubernetes: Quick Bits

Continuing on from our lesson on Persistent Volumes, let's quickly go over Environment Variables and
`ConfigMap` in Kubernetes.

## Environment Variables

### Using Environment Variables

[//]: # (Introduce this section)

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

[//]: # (Introduce this section)

### Using Config Map

1. Let's not place our environment variables inside of configuration files, instead let's create a file named:
   `environment.yaml`,
   which will store them in a `configMap`:

    ```yaml
    apiVersion: v1
    kind: ConfigMap
    metadata:
      name: data-storage-env
    data:
      folder: 'story'
    ```

2. Let's now utilize this `configMap` and access the values from it inside of our configuration:

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

   
