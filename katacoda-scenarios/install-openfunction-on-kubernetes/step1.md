### Kubernetes Cluster

For this scenario, Katacoda has just started a fresh Kubernetes cluster for you. Verify that it's ready for your use:
```
kubectl version --short && \
kubectl get nodes && \
kubectl get componentstatus && \
kubectl cluster-info
```{{execute}}

It should list a 2-node cluster and the control plane components should be reporting Healthy. If it's not healthy, try again in a few moments. If it's still not functioning refresh the browser tab to start a fresh scenario instance before proceeding.