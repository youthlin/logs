package logs

import (
	"os"

	"github.com/youthlin/logs/pkg/callinfo"
)

func GetLogger(opts ...Option) Logger { return defaultFactory.GetLogger(opts...) }

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
	config      *Config
	pkgNameSkip int
}

func NewSimpleFactory() Factory {
	return NewFactory(SimpleAdaptor(os.Stdout), LevelConfig(Debug), 1)
}

// NewFactory return a new logger factory.
// @param a is a logger adaptor, which really emit the log
// @param pkgNameSkip used to get caller's pkg name(as logger name),
// it's value means frames to `GetLogger`
func NewFactory(a Adaptor, c *Config, pkgNameSkip int) Factory {
	Assert(a != nil, "Adaptor can not be nil")
	Assert(c != nil, "Config can not be nil")
	Assert(pkgNameSkip >= 0, "pkgNameSkip can not less than 0")
	return &factory{Adaptor: a, pkgNameSkip: pkgNameSkip, config: c}
}

func (f *factory) SetAdaptor(a Adaptor) {
	Assert(a != nil, "Adaptor can not be nil")
	f.Adaptor = a
}
func (f *factory) SetConfig(c *Config) {
	Assert(c != nil, "Config can not be nil")
	f.config = c
}
func (f *factory) GetLogger(opts ...Option) Logger {
	pkgName := callinfo.Skip(f.pkgNameSkip + 1).PkgName
	return newLogger(f, withName(pkgName), withOpt(opts...))
}
