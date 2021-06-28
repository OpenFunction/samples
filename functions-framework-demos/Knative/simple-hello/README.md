# HTTP request

Start the service.

```go
go mod tidy
go run main.go
```

> The service uses port `8080` by default, if you want to customize the port, you can set it using the environment variable `PORT`.
>
> ```shell
> PORT=8081 go run main.go
> ```

You can see the following.

```shell
2021/06/29 10:10:48 Knative Function serving http: listening on port 8080
```

Now you can access the service via `curl` command.If everything works, you can see the following.

```shell
~# curl http://localhost:8080
Hello, World!
```

