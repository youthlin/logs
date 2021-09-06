# logs
logs is a logging package, which support logging level(diffrent package could has diffrent level).

## import
```shell
go get -u github.com/youthlin/logs
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
