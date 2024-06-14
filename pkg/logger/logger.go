package logger

import (
	"fmt"
	"log"
	"path"
	"runtime"
	"runtime/debug"
)

const callerDepth = 2

const (
	lvlInfo = iota
	lvlImp
	lvlWarn
	lvlDebug
	lvlErr
)

type level struct {
	format string
	prefix string
}

var levels = map[int]level{
	lvlErr:   {format: "\033[1;31m%s\033[0m", prefix: "ERROR"},
	lvlInfo:  {format: "\033[1;32m%s\033[0m", prefix: "INFO"},
	lvlWarn:  {format: "\033[1;33m%s\033[0m", prefix: "WARN"},
	lvlDebug: {format: "\033[1;34m%s\033[0m", prefix: "DEBUG"},
	lvlImp:   {format: "\033[1;35m%s\033[0m", prefix: "IMPORTANT"},
}

func callerInfo() (string, int) {
	_, file, line, ok := runtime.Caller(callerDepth)
	if !ok {
		file = "???"
		line = 0
	}
	return file, line
}

func SDump(args ...any) string {
	dumpStr := ""
	for _, arg := range args {
		dumpStr += fmt.Sprintf(" %+v", arg)
	}
	return dumpStr[1:]
}

func Info(args ...any) {
	file, line := callerInfo()
	prefix := fmt.Sprintf(levels[lvlInfo].format, levels[lvlInfo].prefix)
	prefix = fmt.Sprintf("%s:%d | %s", path.Base(file), line, prefix)
	args = append([]any{prefix}, args...)
	log.Println(args...)
}

func Debug(args ...any) {
	file, line := callerInfo()
	prefix := fmt.Sprintf(levels[lvlDebug].format, levels[lvlDebug].prefix)
	prefix = fmt.Sprintf("%s:%d | %s", path.Base(file), line, prefix)
	log.Println(prefix, SDump(args...))
}

func Imp(args ...any) {
	file, line := callerInfo()
	prefix := fmt.Sprintf(levels[lvlImp].format, levels[lvlImp].prefix)
	prefix = fmt.Sprintf("%s:%d | %s ", path.Base(file), line, prefix)
	args = append([]any{prefix}, args...)
	log.Println(args...)
}

func Warn(args ...any) {
	file, line := callerInfo()
	prefix := fmt.Sprintf(levels[lvlWarn].format, levels[lvlWarn].prefix)
	prefix = fmt.Sprintf("%s:%d | %s ", path.Base(file), line, prefix)
	args = append([]any{prefix}, args...)
	log.Println(args...)
}

func Err(args ...any) {
	file, line := callerInfo()
	prefix := fmt.Sprintf(levels[lvlErr].format, levels[lvlErr].prefix)
	prefix = fmt.Sprintf("%s:%d | %s ", path.Base(file), line, prefix)
	args = append([]any{prefix}, args...)
	args = append(args, "\nstack trace:\n\n", string(debug.Stack()))
	log.Println(args...)
}

func Fatal(args ...any) {
	log.Panicln(args...)
}
