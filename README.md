# go_zero_demo

检测环境
```shell
goctl env check --install --verbose --force
```

创建 http 服务
```shell
./goctl api new http

```
生成 http 代码
```shell

cd http/
    
../goctl api go -api=http.api -dir=./

```

创建 rpc 服务
```shell
cd rpc/

../goctl rpc new hello
```





