## shorturl

before use the script,you should `cd` to the readme.md path like `cd ./service/shorturl`

#### short-url-api

```shell
goctl api go -api './api/shorturl.api' -dir './api'
```

#### short-url-rpc
```shell
goctl rpc protoc 'rpc/transform/transform.proto' --go_out='./rpc/transform' --go-grpc_out='./rpc/transform' --zrpc_out='./rpc/transform'
```

#### model by ddl
```shell
goctl model mysql ddl -src="./model/ddl/shorturl.sql" -dir="./model" --cache
```

#### model by datasource
```shell
goctl model mysql datasource -url="root:root@tcp(localhost:3306)/test" -table="shorturl" -dir="./model" --cache
```
