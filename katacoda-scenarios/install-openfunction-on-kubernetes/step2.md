### Upgrade kubernetes version

The latest version of OpenFunction requires that you have a Kubernetes cluster with version ``>=1.19.0``.

We should upgrade kubernetes version to v1.19.0 first.

1.Upgrade kubeadm,kubelet,kubectl
```
apt-get update -y && \
apt-get install -y --allow-change-held-packages kubelet=1.19.0-00 kubectl=1.19.0-00 kubeadm=1.19.0-00
```{{execute}}

2.Download image: `kubeadm config images pull --kubernetes-version 1.19.0`{{execute}}

3.Upgrade kubernetes controlplane: `kubeadm upgrade apply v1.19.0 --yes --ignore-preflight-errors=all`{{execute}}

4.Upgrade kubernetes worker node: `ssh node01 "apt-get install -y kubelet=1.19.0-00 kubectl=1.19.0-00"`{{execute}}

if this command print error message like:
```
E: Could not get lock /var/lib/dpkg/lock-frontend - open (11: Resource temporarily unavailable)
E: Unable to acquire the dpkg frontend lock (/var/lib/dpkg/lock-frontend), is another process using it?
```

You can just try waiting before unlocking and run the command again.

Or run commands:
```
ssh node01 "killall apt apt-get"
# If none of the above works, remove the lock files. Run in terminal:
ssh node01 "rm /var/lib/apt/lists/lock"
ssh node01 "rm /var/cache/apt/archives/lock"
ssh node01 "rm /var/lib/dpkg/lock*"
ssh node01 "dpkg --configure -a"
```

5.Verify kubernetes status:
```
kubectl version --short && \
kubectl get nodes && \
kubectl get componentstatus && \
kubectl cluster-info
```{{execute}}