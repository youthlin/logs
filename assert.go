package logs

import "fmt"

func Assert(exp bool, fmtArgs ...interface{}) {
	if !exp {
		throw(fmtArgs...)
	}
}

func AssertThen(exp bool, ok, fail func(), fmtArgs ...interface{}) {
	if exp {
		if ok != nil {
			ok()
		}
	} else {
		if fail != nil {
			fail()
		}
		throw(fmtArgs...)
	}
}

func throw(fmtArgs ...interface{}) {
	if len(fmtArgs) > 0 {
		panic(fmt.Sprintf(fmt.Sprintf("%v", fmtArgs[0]), fmtArgs[1:]...))
	} else {
		panic("assert failed")
	}
}
