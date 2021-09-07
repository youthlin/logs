package logs_test

import (
	"context"
	"fmt"
	"strings"
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

	var sb strings.Builder
	logs.SetAdaptor(logs.SimpleAdaptor(&sb))
	logs.With(42, "haha").
		Ctx(kv.Add(context.Background(), "haha", 37)).
		Error("err=%v", arg.ErrJSON("%+v", fmt.Errorf("error msg")))
	// t.Log(sb.String())
	// [2021-09-07 20:22:57.973|Error|github.com/youthlin/logs_test|global_test.go:33]	42=haha haha=37 err="error msg"
	logs.Assert(strings.Contains(sb.String(), `|Error|github.com/youthlin/logs_test|global_test.go:33]	42=haha haha=37 err="error msg"`))
}
