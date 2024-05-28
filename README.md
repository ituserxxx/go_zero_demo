# go_zero_demo

检测环境
```shell
goctl env check --install --verbose --force
```

http 服务
```shell
创建服务
./goctl api new http

# 生成 http 代码
cd http/
    
../goctl api go -api=http.api -dir=./

# 启动 http （项目根目录下执行）
go run http/http.go -f http/etc/http-api.yaml 
```


rpc 服务
```shell
# 创建 rpc 服务
cd rpc/

../goctl rpc new hello

# 生成 rpc 代码

cd rpc/hello

../../goctl rpc protoc hello.proto --go_out=./ --go-grpc_out=./ --zrpc_out=.

# 启动 rpc 服务 （项目根目录下执行）
 go run rpc/hello/hello.go -f rpc/hello/etc/hello.yaml
```


服务依赖环境
```shell
（此代码开发主机在 172.16.9.103 上面，各种外部依赖服务在主机 172.16.9.116 上面）

先启动 docker-compose-env.yml（需要修改 docker_data/prometheus/prometheus.yml 里面监控服务地址）

docker compose -f docker-compose-env.yml up -d

```



