package debug

import (
	"errors"
	"fmt"
	"io"
	"log"
	"reflect"
	"runtime"
	"strings"
)

// Debug prints a debug information to the log with file and line.
func Debug(format string, a ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	info := fmt.Sprintf(format, a...)

	log.Printf("[Debug] %s:%d %v", file, line, info)
}

//MatchErrorWithPanic recover panic
//if there is panic , set error to panic
func MatchErrorWithPanic(err error, errStr string) {
	if r := recover(); r != nil {
		switch t := r.(type) {
		case error:
			err = t
		default:
			err = errors.New(errStr)
		}
	}
}

//GetFileLine return file name and line
func GetFileLine() (string, int) {
	_, file, line, _ := runtime.Caller(1)
	return file, line

}

//GetFullPackageAndFunctionName return full path funciton name
func GetFullPackageAndFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

//GetShortPackageAndFunctionName return package and function name of passed function
func GetShortPackageAndFunctionName(function interface{}) string {
	str := runtime.FuncForPC(reflect.ValueOf(function).Pointer()).Name()
	idx := strings.LastIndex(str, "/")
	return str[idx+1:]

}

//GetFullPackageAndFunctionNameAndLine return package and function name and line
func GetFullPackageAndFunctionNameAndLine(function interface{}) string {
	str := runtime.FuncForPC(reflect.ValueOf(function).Pointer()).Name()

	_, _, line, _ := runtime.Caller(1)
	str += "." + string(line)
	return str
}

//GetShortPackageAndFunctionNameAndLine return package and function name and line
func GetShortPackageAndFunctionNameAndLine(function interface{}) string {
	str := runtime.FuncForPC(reflect.ValueOf(function).Pointer()).Name()
	idx := strings.LastIndex(str, "/")

	_, _, line, _ := runtime.Caller(1)

	str = str[idx+1:]
	str += "." + string(line)
	return str
}

//LogInfo log into out with [Info] prefix
func LogInfo(out io.Writer, v ...interface{}) {
	lr := log.New(out, "[Info]", log.LstdFlags|log.Lshortfile)
	lr.Output(5, fmt.Sprintln(v))
	//calldepth 從一開始，越大越上層，要自己測
}

//LogWarning log into out with [Warning] prefix
func LogWarning(out io.Writer, v ...interface{}) {
	lr := log.New(out, "[Warning]", log.LstdFlags|log.Lshortfile)
	lr.Output(5, fmt.Sprintln(v))
}

//LogError log into out with [Error] prefix
func LogError(out io.Writer, v ...interface{}) {
	lr := log.New(out, "[Error]", log.LstdFlags|log.Lshortfile)
	lr.Output(5, fmt.Sprintln(v))
}

//LogPanic log into stdout with [Panic] prefix.
//Followed by a call to panic().
func LogPanic(out io.Writer, v ...interface{}) {
	lr := log.New(out, "[Painc]", log.LstdFlags|log.Lshortfile)
	lr.Panic(v)
}

//LogFatal log into stdout with [Fatal] prefix.
//Followed by a call to exit(1)
func LogFatal(out io.Writer, v ...interface{}) {
	lr := log.New(out, "[Fatal]", log.LstdFlags|log.Lshortfile)
	lr.Fatal(v)

}
