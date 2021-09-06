package logs

import (
	"fmt"
	"io"
	"strings"

	"github.com/youthlin/logs/pkg/callinfo"
	"github.com/youthlin/logs/pkg/kv"
)

// DiscardAdaptor returns an Adaptor which discard all log messages.
func DiscardAdaptor() Adaptor {
	return discard(0)
}

// SimpleAdaptor returns an Adaptor which print to a specified io.Writer
func SimpleAdaptor(out io.Writer) Adaptor {
	return &simple{out}
}

var _ Adaptor = discard(0)

type discard int

func (discard) Log(Message) {}

var _ Adaptor = (*simple)(nil)

type simple struct {
	io.Writer
}

func (s *simple) Log(m Message) {
	kvs := m.Kvs()
	kvs = append(kvs, kv.Get(m.Ctx())...)
	var sbKV strings.Builder
	for _, kv := range kv.ToString(kvs) {
		sbKV.WriteString(kv[0])
		sbKV.WriteString("=")
		sbKV.WriteString(kv[1])
		sbKV.WriteString(" ")
	}
	// simple.Log <- logger.Log <- logger.Info
	//      1            2            3
	call := callinfo.Skip(3)
	fmt.Fprintf(
		s.Writer,
		fmt.Sprintf(
			// [2006-01-02 15:04:05.000| Info|github.com/youthlin/logs/pkg/adaptor|file:line]	k=v Hello, World
			"[%s|%5s|%s|%s:%d]\t%s%s\n",
			m.Time().Format("2006-01-02 15:04:05.000"),
			m.Level(),
			m.LoggerName(),
			call.ShortFileName(),
			call.Line,
			sbKV.String(),
			m.Msg(),
		),
		m.Args()...,
	)
}
