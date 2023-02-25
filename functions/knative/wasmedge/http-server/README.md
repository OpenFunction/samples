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

## Deployment the wasm function into Kubernetes

You can refer [this](https://openfunction.dev/docs/concepts/wasm_functions/#build-and-run-wasm-functions) guide on how to deploy this wasm function into Kubernetes.