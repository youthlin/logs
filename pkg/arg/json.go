package arg

import (
	"encoding/json"
	"fmt"
)

type Arg struct {
	data interface{}
}

func JSON(data interface{}) *Arg {
	return &Arg{data}
}

func ErrJSON(verb string, err error) *Arg {
	return &Arg{fmt.Sprintf(verb, err)}
}

func (a *Arg) String() string {
	b, err := json.Marshal(a.data)
	if err != nil {
		return fmt.Sprintf("!(BADJSON|err=%+v|data=%#v)", err, a.data)
	}
	return string(b)
}
