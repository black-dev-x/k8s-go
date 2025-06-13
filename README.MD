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

You do not access a pod directly, you access a service that redirects you
just for the sake of learning, the go app is running at 8080
the service is running at 8081
the host port then is 8082
```
kubectl apply -f k8s/service.yaml
kubectl get svc
kubectl port-forward svc/server 8082:8081
```
