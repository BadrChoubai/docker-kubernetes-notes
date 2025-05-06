# 1. First-Steps

1. Perform the command to list all API resources in your Kubernetes cluster. Save the output to a file named `resources.csv`.
2. List the services on your Linux operating system the are associated with Kubernetes. Save the output to a file named `services.csv`.
3. List the status of the kubet service running on the Kubernetes node, output the result to a file named kubelet-status.txt, and save the file in the /tmp directory.
4. Use the declarative syntax to create a Pod from a YAML file in Kubernetes. Save the YAML file as `chap1-pod.yaml`. Use the `kubectl create` command to create the Pod.
5. Using the `kubectl` CLI tool, list all the Services created in your kubenetes cluster across all namespaces. Save the output of the command to a file name `all-k8s-services.txt`.

# 2. Kubernetes Cluster

1. Increase your efficiency with running `kubectl` commands by shortening `kubectl` and creating a shell alias to `k`.
2. Using the `kubectl` CLI tool, get the output of the Pods running in the `kube-system` namespace and show the Pod IP addresses. Save the output of the command to a file named `pod-ip-output.txt`
3. Using the CLI tool, which allows you to view the client certificate that the kubelet uses to authenticate to the Kubernetes API. Output the results to a file named `kubelet-config.txt`
4. Using the `etcdctl` CLI tool, back up the etcd datastore to a snapshot file named `etcdbackup1`. Once that backup is complete, send the output of the command `etcdctl snapshot status etcdbackup1` to a file named `snapshot-status.txt`.
5. Using the `etcdctl` CLI tool, restore the etcd datastore using the same `etcd-backup1` file from the previous question. When you complete the restore operation, `cat` the etcd YAML and save it to a file named `etcd-restore.yaml`.
6. Upgrade the control plane components using kubeadm. When completed, check that everything including kubelet and `kubectl`, is upgraded to 1.24.0.

# 3. Identity and Access Management

1. Create a new Service Account name `secure-sa`, and create a `Pod` that uses this Service Account. Make sure the token is not exposed to the `Pod`.
2. Create a new cluster `Role` name `acme-corp-role` that wille allow the `create` action on `Deployments`, `replicates`, and `DaemonSets`. Bind that cluster Role to the Service Account `secure-sa` and make sure the Service Account can only create the assigned resources within the default namespace and nowhere else. Use `auth can-i` to verify that the `secure-sa` Service Account cannot create `Deployments` in the `kube-system` namespace, and output the result of the command plus the command itself to a file and share that file.

# 4. Deploying Applications in Kubernetes

1. Create a `Pod` named `limited` with the image `httpd` and set the resource requests to `1Gi` gor CPU and `100Mi` for Memory.
2. Create a `ConfigMap` named `ui-data` with the key and value pairs as follows. Apply a `ConfigMap` to a `Pod` named `frontend` with the image `busybox:1.28` and pass it to the `Pod` via the following environment variables: `color.good=purple`, `color.bad=yallow`, `allow.textmode=true`, `how.nice.to.look=fairlyNice`.

# 5. Running Applications in Kubernetes

1. From a three-node cluster, cordon one of the worker nodes. Schedule a Pod without a `nodeSelector`. Uncordon the worker node and edit the Pod, applying a new node name to the `YAML` (set it to the node that was just uncordoned). After replacing the YAML, see if the Pod is scheduled to the recently uncordoned node.
2. Start a basic `nginx` Deployment; remove the taint from the control plane node so that Pods don't need a toleration to be scheduled to it. Add a `nodeSelector` to the Pod spec within the Deployment, and see if the Pod is now running on the control plane node.

# 6. Communication in a Kubernetes Cluster

1. Create a `Deployment` named `hello` using the image `nginxdemos/hello:plain-text` with the `kubectl` command line. Expose the `Deployment` to create a `ClusterIP` Service named `hello-svc` that can communicate over `port 80` using the `kubectl` command line. use the correct `kubectl` command to verify that it's a `ClusterIP` Service with the correct port exposed.
2. Change the `hello-svc` Service created in the previous exercise to a `NodePort` Service, where the `NodePort` should be `30000`. Be sure to edit the Service in place, without creating a new YAML or issuing a new imperative command. Communicate with the Pods within the `hello` Deployment via the `NodePort` Service using `curl`
3. Install an `Ingress` controller in the cluster using the command `k apply -f https://raw.githubusercontent.com/chadmcrowell/acing-the-cka-exam/refs/heads/main/ch_06/nginx-ingress-controller.yaml`. Change the `hello-svc` Service to a `ClusterIP` Service and create in `Ingress` resource that will route to the `hello-svc` Service when a client requests `hello.com`.
4. Create a new `kind` cluster without a CNI. Install the bridge CNI, followed by the Calico CNI. After installing the CNI, verify that the `CoreDNS` Pods are up and running and the nodes are in a ready state.

# 7. Storage in Kubernetes

1. Create a `Pod` name `two-web` with two containers. The first container will be named `httpd` and will use the image `httpd:alpine3.17`. The second container will be named `nginx` and will use the image `nginx:1.23.3-alpine`.
2. Both containers should access the same volume that is shared from local storage on the container itself.
3. `Container1` will mount the colume to `/var/www/html/` and `Container2` will mount the volume to `/usr/share/nginx/html/`.
4. Start up the Pod and ensure everything is mounted and shared correctly.

# 8. Troubleshooting Kubernetes

1. Move the file `kube-scheduler.yaml` to the `/tmp` directory with the command `mv /etc/Kubernetes/manifests/kube-scheduler.yaml /tmp/kube-scheduler.yaml`.
2. Create a Pod with the command `k run nginx -image nginx`. List the Pods and see if the Pod is in a running state.
3. Determine why the Pod is not starting by looking at the events and the logs. Detemine how to fix it and get the Pod back in a running state.
4. Run the command `curl https://raw.githubusercontent.com/chadmcrowell/acing-the-cka-exam/main/ch_08/10-kubeadm.conf --silent --output /etc/systemd/system/kubelet.service.d/10-kubeadm.conf; systemctl daemon-reload; systemctl restart kubelet`.
5. Check the status of the kubelet, and got through the troubleshooting steps to resolve the problem with the kubelet service.
