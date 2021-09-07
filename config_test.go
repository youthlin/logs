package logs_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/youthlin/logs"
)

func TestUnmarshalConfig(t *testing.T) {
	var c logs.Config
	err := json.Unmarshal([]byte(`{
		"root": "Warn",
		"loggers": {
			"github.com/youthlin": "Info"
		}
	}`), &c)
	logs.Assert(err == nil)
	t.Log(c)
	logs.Assert(c.Root == logs.LevelWarn)
	logs.Assert(reflect.DeepEqual(c.Loggers, map[string]logs.Level{"github.com/youthlin": logs.LevelInfo}))
}
