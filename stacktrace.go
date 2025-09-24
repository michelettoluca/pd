package pd

import (
	"fmt"
	"strings"
)

var MaxStackDepth = 10

type stackTrace struct {
	frames []stackFrame
}

func newStackTrace() stackTrace {
	var frames []stackFrame

	for i := 0; len(frames) < MaxStackDepth; i++ {
		sf, ok := newStackFrame(i)
		if !ok {
			break
		}

		isPdError := strings.HasPrefix(sf.function, "pd.")
		isGoError := strings.Contains(sf.function, "runtime.")
		if isPdError || isGoError {
			continue
		}

		frames = append(frames, sf)
	}

	return stackTrace{
		frames: frames,
	}
}

func (s stackTrace) String() string {
	var text string
	for _, frame := range s.frames {
		text += fmt.Sprintf("%s:%d %s\n", frame.file, frame.line, frame.function)
	}

	return text
}
