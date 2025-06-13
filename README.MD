DWARF = Debugging With Arbitrary Record Format

To push this image to the dockerhub
```
docker login
docker build -t blackdevx/k8s-go:latest -f Dockerfile.prod .
docker push blackdevx/k8s-go:latest
```

```
kind create cluster --name=goexpert
kubectl cluster-info --context kind-goexpert
kubectl get nodes
kubectl apply -f k8s/deployment.yaml 
kubectl get pods
```
