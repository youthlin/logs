package logs

import (
	"context"

	"github.com/youthlin/logs/pkg/kv"
)

var _ Logger = (*logger)(nil)

type logger struct {
	*factory
	name string
	skip int
	ctx  context.Context
	kvs  []interface{}
}

type newLoggerOpt func(*logger)

func withName(name string) newLoggerOpt {
	return func(l *logger) { l.name = name }
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
		skip:    l.skip,
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
func (l *logger) Skip(skip int) Logger {
	copy := l.clone()
	copy.skip = skip
	return copy
}
func (l *logger) GetSkip() int {
	return l.skip
}

func (l *logger) Trace(fmt string, args ...interface{}) { l.Log(LevelTrace, fmt, args...) }
func (l *logger) Debug(fmt string, args ...interface{}) { l.Log(LevelDebug, fmt, args...) }
func (l *logger) Info(fmt string, args ...interface{})  { l.Log(LevelInfo, fmt, args...) }
func (l *logger) Warn(fmt string, args ...interface{})  { l.Log(LevelWarn, fmt, args...) }
func (l *logger) Error(fmt string, args ...interface{}) { l.Log(LevelError, fmt, args...) }

func (l *logger) Log(lvl Level, format string, args ...interface{}) {
	if adaptor := l.factory.Adaptor; adaptor != nil {
		if lvl >= l.level() {
			adaptor.Log(NewMsg(l, lvl, format, args...))
		}
	}
}

func (l *logger) level() Level {
	lvl := l.config.trie.Search(l.name)
	return lvl.(Level)
}
