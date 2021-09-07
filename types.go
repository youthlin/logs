package logs

import (
	"context"
	"time"

	"github.com/youthlin/logs/pkg/trie"
)

type (
	// Level log level 日志级别
	Level int

	// Message log message 日志实体
	Message interface {
		LoggerName() string
		Level() Level
		Time() time.Time
		Skip() int
		Ctx() context.Context
		Kvs() []interface{}
		Msg() string
		Args() []interface{}
	}
	// Adaptor used to handle log message. 实现该接口来实际打日志
	Adaptor interface {
		Log(Message)
	}
	Config struct {
		Root    Level
		Loggers map[string]Level
		trie    *trie.Tire
	}
	// LoggerOpt logger option. 日志配置
	LoggerOpt struct {
		Name    *string
		AddSkip int
	}
	// Option used to get logger. 目的是获取 Logger Name
	Option func(*LoggerOpt)

	// Factory used to create logger. 获取 Logger 的日志工厂
	Factory interface {
		SetAdaptor(Adaptor)
		SetConfig(*Config)
		GetLogger(opts ...Option) Logger
	}
	// Logger is a interface to logging. 打日志接口
	Logger interface {
		// Name return the Logger's name.
		// 不同名称的 Logger 可以有不同的日志级别
		Name() string
		// Ctx attach context to this Logger. see kv.Add(ctx, kvs...).
		// 在 Logger 上附加 ctx, 可以使用 kv.Add 为 ctx 附带 kv
		Ctx(ctx context.Context) Logger
		GetCtx() context.Context
		// With attach key-values to this Logger. 在 Logger 上附带 kv
		With(kvs ...interface{}) Logger
		GetKVs() []interface{}
		// Skip relative to below methods(Trace/Debug...)
		// 相对于调用以下 Trace/Debug 再往上跳过 skip 层堆栈
		// 如果 Adaptor 会打印堆栈，应该带上这个 skip
		Skip(skip int) Logger
		GetSkip() int

		Trace(fmt string, args ...interface{})
		Debug(fmt string, args ...interface{})
		Info(fmt string, args ...interface{})
		Warn(fmt string, args ...interface{})
		Error(fmt string, args ...interface{})
	}
)

// WithName get logger with name
func WithName(name string) Option {
	return func(lo *LoggerOpt) { lo.Name = &name }
}

// AddSkip get logger with skip frames(relative to factory.GetLogger)
func AddSkip(skip int) Option {
	return func(lo *LoggerOpt) { lo.AddSkip += skip }
}
