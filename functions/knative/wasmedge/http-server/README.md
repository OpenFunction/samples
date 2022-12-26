# Sample Function WasmEdge ( written in Rust )

## Prerequisites

### OpenFunction

You can refer to the [Installation Guide](https://github.com/OpenFunction/OpenFunction#install-openfunction) to setup OpenFunction.

## Run it locally

1. Build

```bash
cargo build --target wasm32-wasi --release
```

2. Run

```bash
wasmedge target/wasm32-wasi/release/wasmedge_hyper_server.wasm
```

3. Test

Run the following from another terminal.

```bash
$ curl http://localhost:8080/echo -X POST -d "WasmEdge"
WasmEdge
```

## Deployment

> To setup `WasmEdge` workload runtime in kubernetes cluster and push images to a container registry,
> please refer to the [prerequisites](../../getting-started/Quickstarts/prerequisites) section for more info.

1. Create function

```shell
kubectl apply -f wasmedge-http-server.yaml
 ```

2. Check the function status

```shell
kubectl get functions.core.openfunction.io -w
NAME                   BUILDSTATE   SERVINGSTATE   BUILDER         SERVING         ADDRESS                                                      AGE
wasmedge-http-server   Succeeded    Running        builder-4p2qq   serving-lrd8c   http://wasmedge-http-server.default.svc.cluster.local/echo   12m
  ```

3. Access function

Once the `BUILDSTATE` becomes `Succeeded` and the `SERVINGSTATE` becomes `Running`, you can access this function through the address in the `ADDRESS` field:
You can observe the process of a function with the following command:

```shell
kubectl run curl --image=radial/busyboxplus:curl -i --tty
curl http://wasmedge-http-server.default.svc.cluster.local/echo  -X POST -d "WasmEdge"
WasmEdge
```
