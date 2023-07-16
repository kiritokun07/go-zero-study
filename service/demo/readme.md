## demo

#### demo-api
```shell
goctl api go -api './api/demo.api' -dir './api'
```

#### model by datasource
```shell
goctl model mysql datasource -url="root:root@tcp(localhost:3306)/test" -table="*" -dir="./model" --cache
```