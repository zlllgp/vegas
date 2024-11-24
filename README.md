# 权益
```
    配置 maker
      /        \
玩法 las -> 权益 vegas
```

如何被其他模块依赖proto api , 增加版本号发布tag
```
git tag v1.0.2
git push origin v1.0.2
github 页面上发布 release
```

解决依赖包问题
```
go mod edit -droprequire=github.com/apache/thrift/lib/go/thrift
go mod edit -replace=github.com/apache/thrift=github.com/apache/thrift@v0.13.0
```
# docker
docker export xxx > docker.tar
进入容器
docker exec -it 4e*** sh

# gomock
https://blog.csdn.net/xidianhuihui/article/details/131128948
go install github.com/golang/mock/mockgen@latest
go get github.com/golang/mock/mock
go get github.com/golang/mock/mockgen

# todo
redis
cache
log