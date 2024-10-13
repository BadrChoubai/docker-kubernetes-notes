eval "$(minikube docker-env)"
docker build -t storage-demo-app:1.0 .
