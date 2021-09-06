package logs

import (
	"context"

	"github.com/youthlin/logs/pkg/kv"
)

var _ Logger = (*logger)(nil)

type logger struct {
	*factory
	name string
	ctx  context.Context
	kvs  []interface{}
}

type newLoggerOpt func(*logger)

func withName(name string) newLoggerOpt {
	return func(l *logger) { l.name = name }
}

func withOpt(opts ...Option) newLoggerOpt {
	lo := new(LoggerOpt)
	for _, opt := range opts {
		opt(lo)
	}
	return func(l *logger) {
		if lo.Name != nil {
			l.name = *lo.Name
		}
	}
}

func newLogger(f *factory, opts ...newLoggerOpt) *logger {
	Assert(f != nil)
	l := &logger{factory: f, ctx: context.Background()}
	for _, opt := range opts {
		opt(l)
	}
	return l
}

func (l *logger) clone() *logger {
	return &logger{
		factory: l.factory,
		name:    l.name,
		ctx:     l.ctx,
		kvs:     append(make([]interface{}, 0, len(l.kvs)), l.kvs...),
	}
}

func (l *logger) Name() string {
	return l.name
}

func (l *logger) Ctx(ctx context.Context) Logger {
	copy := l.clone()
	copy.ctx = ctx
	return copy
}

func (l *logger) GetCtx() context.Context {
	return l.ctx
}

func (l *logger) With(kvs ...interface{}) Logger {
	copy := l.clone()
	copy.kvs = append(copy.kvs, kv.TrimOdd(kvs)...)
	return copy
}

func (l *logger) GetKVs() []interface{} {
	return l.kvs
}

func (l *logger) Trace(fmt string, args ...interface{}) { l.Log(Trace, fmt, args...) }
func (l *logger) Debug(fmt string, args ...interface{}) { l.Log(Debug, fmt, args...) }
func (l *logger) Info(fmt string, args ...interface{})  { l.Log(Info, fmt, args...) }
func (l *logger) Warn(fmt string, args ...interface{})  { l.Log(Warn, fmt, args...) }
func (l *logger) Error(fmt string, args ...interface{}) { l.Log(Error, fmt, args...) }

func (l *logger) Log(lvl Level, format string, args ...interface{}) {
	if adaptor := l.factory.Adaptor; adaptor != nil {
		if lvl >= GetLevel(l.name, l.config) {
			adaptor.Log(NewMsg(l, lvl, format, args...))
		}
	}
}

func GetLevel(loggerName string, c *Config) Level {
	lvl := c.Trie().Search(loggerName)
	return lvl.(Level)
}
