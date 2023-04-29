## short_url

#### short-url-api

```shell
goctl api go -api './api/shorturl.api' -dir './api''
```

#### short-url-rpc
```shell
goctl rpc protoc 'rpc/transform/transform.proto' --go_out='./rpc/transform' --go-grpc_out='./rpc/transform' --zrpc_out='./rpc/transform'
```