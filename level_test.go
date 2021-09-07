package logs_test

import (
	"encoding/json"
	"testing"

	"github.com/youthlin/logs"
)

func TestLevel(t *testing.T) {
	lvl := logs.LevelTrace
	b, err := json.Marshal(lvl)
	logs.AssertThen(len(b) > 0, func() {
		t.Logf("to json: %s", b)
	}, nil)
	logs.Assert(err == nil)
	var level logs.Level
	err = json.Unmarshal(b, &level)
	logs.AssertThen(err == nil, nil, func() {
		t.Logf("err=%+v", err)
	})
	logs.Assert(level == lvl)
}
