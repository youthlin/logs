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
		Name *string
	}
	// Option used to get logger
	Option func(*LoggerOpt)

	// Factory used to create logger. 获取 Logger 的日志工厂
	Factory interface {
		SetAdaptor(Adaptor)
		SetConfig(*Config)
		GetLogger(opts ...Option) Logger
	}
	// Logger is a interface to logging. 打日志接口
	Logger interface {
		Name() string
		Ctx(ctx context.Context) Logger
		GetCtx() context.Context
		With(kvs ...interface{}) Logger
		GetKVs() []interface{}
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
