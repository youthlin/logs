package logs

import "fmt"

func Assert(exp bool, fmtArgs ...interface{}) {
	if !exp {
		if len(fmtArgs) > 0 {
			panic(fmt.Sprintf(fmt.Sprintf("%v", fmtArgs[0]), fmtArgs[1:]...))
		} else {
			panic("assert failed")
		}
	}
}
