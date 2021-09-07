package logs

import (
	"fmt"
	"strings"
)

const (
	LevelUnset Level = iota
	LevelAll
	LevelTrace
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelNone
)

func (lvl Level) String() string {
	switch lvl {
	case LevelUnset:
		return "Unset"
	case LevelAll:
		return "All"
	case LevelTrace:
		return "Trace"
	case LevelDebug:
		return "Debug"
	case LevelInfo:
		return "Info"
	case LevelWarn:
		return "Warn"
	case LevelError:
		return "Error"
	case LevelNone:
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
		*lvl = LevelAll
	case "trace":
		*lvl = LevelTrace
	case "debug":
		*lvl = LevelDebug
	case "info":
		*lvl = LevelInfo
	case "warn":
		*lvl = LevelWarn
	case "error":
		*lvl = LevelError
	case "none":
		*lvl = LevelNone
	default:
		*lvl = LevelUnset
	}
	return nil
}
