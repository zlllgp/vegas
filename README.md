如何被其他模块依赖proto api , 发布tag
```
git tag v1.0.0
git push origin v1.0.0 
github release
```

```
go mod edit -droprequire=github.com/apache/thrift/lib/go/thrift
go mod edit -replace=github.com/apache/thrift=github.com/apache/thrift@v0.13.0
```