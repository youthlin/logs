package logs

import (
	"fmt"
	"strings"
)

const (
	Unset Level = iota
	All
	Trace
	Debug
	Info
	Warn
	Error
	None
)

func (lvl Level) String() string {
	switch lvl {
	case Unset:
		return "Unset"
	case All:
		return "All"
	case Trace:
		return "Trace"
	case Debug:
		return "Debug"
	case Info:
		return "Info"
	case Warn:
		return "Warn"
	case Error:
		return "Error"
	case None:
		return "None"
	}
	return fmt.Sprintf("Level(%d)", lvl)
}

func (lvl Level) MarshalText() (text []byte, err error) {
	return []byte(lvl.String()), nil
}

func (lvl *Level) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {
	case "all":
		*lvl = All
	case "trace":
		*lvl = Trace
	case "debug":
		*lvl = Debug
	case "info":
		*lvl = Info
	case "warn":
		*lvl = Warn
	case "error":
		*lvl = Error
	case "none":
		*lvl = None
	default:
		*lvl = Unset
	}
	return nil
}
