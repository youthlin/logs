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

## api
```go
// ---------- use global function ----------
logs.Name() // return the loggers name, default package name
logs.Ctx(ctx) // return a logger with ctx, which can add some key-value pairs via kv.Add(ctx)
logs.With(kvs...) // retuan a logger with key-value pairs
logs.Trace(fmt, args...)
logs.Debug(fmt, args...)
logs.Info(fmt, args...)
logs.Warn(fmt, args...)
logs.Error(fmt, args...)

// ---------- use Logger interface ----------
var log = logs.GetLogger() // get a logger, which name is package name
log.Ctx(ctx).
    With(kvs...).
    Debug(fmt, args...)

// ----------  set log level for each package ----------
logs.SetConfig(&logs.Config{
    Root: logs.LevelError,
    Loggers: map[string]logs.Level{
        "github.com": logs.Info,
        "github.com/youthlin": logs.Debug,
    },
})
// abc -> use root level: error
logs.GetLogger(logs.WithName("abc")).
    Warn("not call adaptor.Log so won't print")
logs.GetLogger(logs.WithName("abc")).
    Error("to adaptor")

// github.com/some -> use github.com level: info
logs.GetLogger(logs.WithName("github.com/some")).
    Debug("debug not print")
logs.GetLogger(logs.WithName("github.com/some")).
    Info("info message")

// github.com/youthlin -> debug
logs.GetLogger(logs.WithName("github.com/youthlin")).
    Debug("debug message")


// ---------- use adaptor ----------
logs.SetAdaptor(logs.DiscardAdaptor)
logs.Info("log message was process by adaptor")
// a dicard adaptor not ptint any log

// import _ "github.com/youthlin/z"
logs.Info("import github.com/youthlin/z will register a ZapAdaptor, so this message would log by zap")

// or use zap via custom config(Encoder/WriteSyncer/Core)
zapConfig := z.DefaultConfig() // or config from yaml/json file
z.SetConfig(zapConfig)
logs.With(kvs...).Ctx(ctx).Info("format", args...)

// custom your adaptor
type Adaptor interface(){
    Log(Message)
}
```
