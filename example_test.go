package logs_test

import (
	"context"
	"os"
	"testing"

	"github.com/youthlin/logs"
	"github.com/youthlin/logs/pkg/kv"
)

var log = logs.GetLogger()

func TestLogs(t *testing.T) {
	log.Trace("trace log default not print")
	log.Info("Hello, %s", "World")
	// [2021-09-03 17:53:26.507| Info|github.com/youthlin/logs_test|example_test.go:15]	Hello, World
	foo(context.Background())
	// [2021-09-03 17:53:26.507|Debug|github.com/youthlin/logs_test|example_test.go:26]	func=foo Foo
	// [2021-09-03 17:53:26.507|Debug|github.com/youthlin/logs_test|example_test.go:27]	func=foo key=42 with key-value
	log.Debug("no key-value")
	// [2021-09-03 17:53:26.508|Debug|github.com/youthlin/logs_test|example_test.go:20]	Debug
}

func foo(ctx context.Context) {
	var log = log.With("func", "foo") // new var
	log.Debug("Foo")
	log.Ctx(kv.Add(ctx, "key", 42)).Warn("with key-value") // with context
}

func TestDiscard(t *testing.T) {
	factory := logs.NewFactory(logs.DiscardAdaptor(), logs.LevelConfig(logs.None), 0)
	log := factory.GetLogger()
	log.Info("this message will discard")
	// <nothing output>
}

func TestSimaple(t *testing.T) {
	factory := logs.NewFactory(logs.SimpleAdaptor(os.Stdout), logs.LevelConfig(logs.Info), 0)
	log := factory.GetLogger()
	log.Info("Hello")
}

func TestConfig(t *testing.T) {
	logs.SetConfig(&logs.Config{
		Root: logs.Warn,
		Loggers: map[string]logs.Level{
			"github.com/youthlin/logs": logs.Debug,
		},
	})
	log := logs.GetLogger()
	log.Trace("not print")
	log.Debug("debug level")
}
