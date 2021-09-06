package logs

import (
	"context"
	"time"
)

var _ Message = (*Msg)(nil)

type Msg struct {
	name  string
	level Level
	time  time.Time
	ctx   context.Context
	kvs   []interface{}
	fmt   string
	args  []interface{}
}

func NewMsg(l Logger, lvl Level, fmt string, args ...interface{}) *Msg {
	Assert(l != nil)
	return &Msg{
		name:  l.Name(),
		level: lvl,
		time:  time.Now(),
		ctx:   l.GetCtx(),
		kvs:   l.GetKVs(),
		fmt:   fmt,
		args:  args,
	}
}

func (m *Msg) LoggerName() string {
	return m.name
}
func (m *Msg) Level() Level {
	return m.level
}
func (m *Msg) Time() time.Time {
	return m.time
}
func (m *Msg) Ctx() context.Context {
	return m.ctx
}
func (m *Msg) Kvs() []interface{} {
	return m.kvs
}
func (m *Msg) Msg() string {
	return m.fmt
}
func (m *Msg) Args() []interface{} {
	return m.args
}
