# logs
logs is a logging facade, which supports logging level(diffrent package could has diffrent level),
and it supports any logging implementation(std log/zap, etc) by Adaptor interface.
支持为每个包设置日志级别的一个日志门面，可以通过 Adaptor 接口对接标准 log/zap 等任意日志实现。

## import
```shell
go get -u github.com/youthlin/logs
# 国内镜像
go mod edit -replace github.com/youthlin/logs@latest=gitee.com/youthlin/logs@latest&&go mod tidy
```

## examples

### simplest
```go
var log = logs.GetLogger()
func main(){
    log.Debug("Hello %s", "World")
}
```

### args print as json
```go
func foo(){
    myLogArg, err := ...
    log.Info("some value = %s | err = %s", arg.JSON(myLogArg), arg.ErrJSON("%+v", err))
}
```

### with key-value
```go
var log = logs.GetLogger()
func main(){
    log.With("key", 42).Debug("I have a prefix `key=42`")
    foo()
    log.Info("This msg do NOT have any key-value prefies")
}
func foo(){
    var log = log.With("key", "value") // new var
    log.Debug("my prefix is key=value")
    log.Debug("my prefix is key=value, too")
    bar(context.Background())
}
func bar(ctx context.Context){
    ctx = kv.Add(ctx, "key", "value")
    log.Ctx(ctx).Info("key-value may from ctx")
}
```

### package-level
```go
var log logs.Logger
func init(){
    logs.SetConfig(logs.Config{
        Root: logs.Warn,
        Loggers: map[string]logs.Level{
            "github.com/youthlin/app": logs.Debug,
        },
    })
    log = logs.GetLogger()
}
func main(){
    log.Debug("debug level")
}
```

### log Adaptor

```go
var log logs.Logger
func init(){
    // Adaptor is a interface which contains only one method: Log(logs.Message)
    logs.SetAdaptor(logs.DiscardAdaptor())
    log = logs.GetLogger()
}
func main(){
    log.Debug("use custom adaptor to logging")
}
```
