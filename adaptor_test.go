package logs_test

import (
	"strings"
	"testing"

	"github.com/youthlin/logs"
)

func TestDiscardAdaptor(t *testing.T) {
	a := logs.DiscardAdaptor()
	a.Log(logs.NewMsg(logs.GetLogger(), logs.LevelDebug, "hello %s", "world"))
}

func TestSimpleAdaptor(t *testing.T) {
	var sb strings.Builder
	a := logs.SimpleAdaptor(&sb)
	func() { // app: logs.Debug()
		func() { // logger: logger.Log
			a.Log(logs.NewMsg(logs.GetLogger(), logs.LevelDebug, "hello %s", "world"))
			a.Log(logs.NewMsg(logs.GetLogger().With("key", 42), logs.LevelDebug, "with key-value pairs"))
		}()
	}()
	content := sb.String()
	t.Log(content) // [2021-09-06 13:31:39.475|Debug|github.com/youthlin/logs_test|adaptor_test.go:23]	hello world
	logs.Assert(strings.Contains(content, "|Debug|github.com/youthlin/logs_test|adaptor_test.go:23]	hello world"))
	logs.Assert(strings.Contains(content, "key=42 with key-value pairs"))

}
