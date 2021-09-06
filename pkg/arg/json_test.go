package arg_test

import (
	"fmt"
	"testing"

	"github.com/youthlin/logs/pkg/arg"
)

func TestJSON(t *testing.T) {
	type S struct {
		Age int
	}
	for _, tCase := range []struct {
		data interface{}
		want string
	}{
		{1, `1`},
		{"1", `"1"`},
		{S{20}, `{"Age":20}`},
		{nil, `null`},
	} {
		if got := fmt.Sprintf("%v", arg.JSON(tCase.data)); got != tCase.want {
			t.Errorf("got= %q want = %q", got, tCase.want)
		}
	}
}

func TestErrJSON(t *testing.T) {
	e := arg.ErrJSON("%+v", fmt.Errorf("error message"))
	if fmt.Sprintf("%v", e) != `"error message"` {
		t.Error()
	}
}
