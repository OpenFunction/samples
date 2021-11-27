### Prerequisites

1.Untaint master node: `kubectl taint node controlplane node-role.kubernetes.io/master-`{{execute}}

2.Setup OpenFunction **Builder** and **Serving**: `curl -sfL https://raw.githubusercontent.com/OpenFunction/OpenFunction/main/hack/deploy.sh | bash -s  -- --with-cert-manager --with-shipwright --with-openFuncAsync`{{execute}}