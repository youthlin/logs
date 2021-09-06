package callinfo

import (
	"runtime"
	"strings"
)

// CallInfo 调用信息，包含包名、文件名、函数名、行号
type CallInfo struct {
	PkgName,
	FileName,
	ShortFile,
	FuncName string
	Line int
}

func (c *CallInfo) ShortFileName() string {
	if c.ShortFile == "" {
		c.ShortFile = shortFileName(c.FileName)
	}
	return c.ShortFile
}

// Get 获取当前调用栈（调用本函数的地方）的调用信息，可能返回 nil
func Get() *CallInfo {
	return Skip(1)
}

// Skip 获取当前调用栈（调用本函数的地方）往上 skip 层的调用信息，可能返回 nil
func Skip(skip int) *CallInfo {
	pc, file, line, ok := runtime.Caller(skip + 1)
	if !ok {
		return nil
	}
	pkgName := ""
	funcName := runtime.FuncForPC(pc).Name()
	// funcName 两种形式
	// - github.com/youthlin/logs/pkg/core.GetCallInfoSkip
	// - main.main
	if index := strings.LastIndex(funcName, "/"); index >= 0 { // 有斜线的，先把斜线之前的模块路径提取
		pkgName = funcName[:index]
		funcName = funcName[index+1:]
	}
	if index := strings.Index(funcName, "."); index >= 0 { // 剩下函数名，第一个点号之前的部分也属于包名
		if pkgName == "" { // main.main
			pkgName = funcName[:index]
		} else { // module/path/to/pkg.func
			pkgName += "/" + funcName[:index]
		}
		funcName = funcName[index+1:]
	}
	return &CallInfo{
		PkgName:  pkgName,
		FileName: file,
		FuncName: funcName,
		Line:     line,
	}
}

func shortFileName(file string) string {
	short := file
	for i := len(file) - 1; i >= 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			break
		}
	}
	return short
}
