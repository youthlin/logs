package logs

import (
	"os"

	"github.com/youthlin/logs/pkg/callinfo"
)

func GetLogger(opts ...Option) Logger {
	return defaultFactory.GetLogger(append(opts, AddSkip(1))...)
}

var defaultFactory = NewSimpleFactory()

func SetFactory(f Factory) {
	Assert(f != nil, "Factory can not be nil")
	defaultFactory = f
}
func GetFactory() Factory  { return defaultFactory }
func SetAdaptor(a Adaptor) { defaultFactory.SetAdaptor(a) }
func SetConfig(c *Config)  { defaultFactory.SetConfig(c) }

var _ Factory = (*factory)(nil)

type factory struct {
	Adaptor
	config *Config
}

func NewSimpleFactory() Factory {
	return NewFactory(SimpleAdaptor(os.Stdout), LevelConfig(LevelDebug))
}

// NewFactory return a new logger factory.
func NewFactory(a Adaptor, c *Config) Factory {
	Assert(a != nil, "Adaptor can not be nil")
	Assert(c != nil, "Config can not be nil")
	return &factory{Adaptor: a, config: c}
}

func (f *factory) SetAdaptor(a Adaptor) {
	Assert(a != nil, "Adaptor can not be nil")
	f.Adaptor = a
}
func (f *factory) SetConfig(c *Config) {
	Assert(c != nil, "Config can not be nil")
	f.config = c
}

// GetLogger 获取 Logger
//
// Option:
//  - WithName 直接使用给定的名称；如有，会忽略 AddSkip
//  - AddSkip 获取直接调用本方法处再往上 n 层的地方的包名
func (f *factory) GetLogger(opts ...Option) Logger {
	lo := new(LoggerOpt)
	for _, opt := range opts {
		opt(lo)
	}
	name := ""
	if lo.Name != nil {
		name = *lo.Name
	} else {
		name = callinfo.Skip(lo.AddSkip + 1).PkgName
	}
	return newLogger(f, withName(name))
}

// WithName get logger with name
func WithName(name string) Option {
	return func(lo *LoggerOpt) { lo.Name = &name }
}

// AddSkip get logger with skip frames(relative to factory.GetLogger)
func AddSkip(skip int) Option {
	return func(lo *LoggerOpt) { lo.AddSkip += skip }
}
