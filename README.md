# go_zero_demo

检测环境
```shell
goctl env check --install --verbose --force
```

创建 http 服务
```shell
./goctl api new http

# 生成 http 代码
cd http/
    
../goctl api go -api=http.api -dir=./

```


创建 rpc 服务
```shell
cd rpc/

../goctl rpc new hello

# 生成 rpc 代码

cd hello

../../goctl rpc protoc hello.proto --go_out=./ --go-grpc_out=./ --zrpc_out=.
```


服务依赖环境
```shell
（此代码开发主机在 172.16.9.103 上面，各种外部依赖服务在主机 172.16.9.116 上面）

先启动 docker-compose-env.yml（需要修改 docker_data/prometheus/prometheus.yml 里面监控服务地址）

docker compose -f docker-compose-env.yml up -d

```



