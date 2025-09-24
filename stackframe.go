package pd

import (
	"runtime"
	"strings"
)

type stackFrame struct {
	pc       uintptr
	file     string
	function string
	line     int
}

func newStackFrame(skip int) (stackFrame, bool) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return stackFrame{}, false
	}

	function := runtime.FuncForPC(pc)
	if function == nil {
		return stackFrame{}, false
	}

	functionNameParts := strings.Split(function.Name(), "/")
	functionName := functionNameParts[len(functionNameParts)-1]

	return stackFrame{
		pc:       pc,
		file:     file,
		line:     line,
		function: functionName,
	}, true
}
