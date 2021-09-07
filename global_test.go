package logs_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/youthlin/logs"
	"github.com/youthlin/logs/pkg/arg"
	"github.com/youthlin/logs/pkg/callinfo"
	"github.com/youthlin/logs/pkg/kv"
)

func TestName(t *testing.T) {
	pkg := callinfo.Get().PkgName
	loggerName := logs.Name()
	log.Info("pkg name = %s | loggerName = %s", pkg, loggerName)
	logs.Assert(loggerName == pkg)
}

func TestGlobal(t *testing.T) {
	logs.Trace("trace")
	logs.Debug("debug")
	logs.Info("info")
	logs.Warn("warn")
	logs.Error("error")
	logs.With(42, "haha").
		Ctx(kv.Add(context.Background(), "haha", 37)).
		Error("err=%v", arg.ErrJSON("%+v", fmt.Errorf("error msg")))
}
