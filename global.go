package logs

import "context"

func getLogger() Logger {
	// callinfo.Skip <- [0]defaultFactory.GetLogger <- [1]logs.getLogger <- [2]logs.Debug <- [3]app
	//                        AddSkip+1
	// AddSkip=2
	return defaultFactory.GetLogger(AddSkip(2))
}
func Name() string {
	return getLogger().Name()
}
func Ctx(ctx context.Context) Logger {
	return getLogger().Ctx(ctx)
}
func With(kvs ...interface{}) Logger {
	return getLogger().With(kvs...)
}

func Trace(fmt string, args ...interface{}) {
	getLogger().Skip(1).Trace(fmt, args...)
}
func Debug(fmt string, args ...interface{}) {
	getLogger().Skip(1).Debug(fmt, args...)
}
func Info(fmt string, args ...interface{}) {
	getLogger().Skip(1).Info(fmt, args...)
}
func Warn(fmt string, args ...interface{}) {
	getLogger().Skip(1).Warn(fmt, args...)
}
func Error(fmt string, args ...interface{}) {
	getLogger().Skip(1).Error(fmt, args...)
}
