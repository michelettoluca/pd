package pd

var _ error = Error{}

const (
	messageUndefined = ""
	statusUndefined  = -1
	codeUndefined    = ""
)

var (
	FallbackMessage = messageUndefined
	FallbackStatus  = statusUndefined
	FallbackCode    = codeUndefined
)

type Error struct {
	err        error
	message    string
	status     int
	code       string
	stackTrace stackTrace
}

func (pd Error) Unwrap() error {
	return pd.err
}

func (pd Error) Error() string {
	return pd.Message()
}

func (pd Error) Message() string {
	if match, ok := findClosestMatch(pd, func(err Error) bool { return err.message != messageUndefined }); ok {
		return match.message
	}

	return FallbackMessage
}

func (pd Error) Status() int {
	if match, ok := findClosestMatch(pd, func(err Error) bool { return err.status != statusUndefined }); ok {
		return match.status
	}

	return FallbackStatus
}

func (pd Error) Code() string {
	if match, ok := findClosestMatch(pd, func(err Error) bool { return err.code != codeUndefined }); ok {
		return match.code
	}

	return FallbackCode
}

func (pd Error) StackTrace() string {
	return findDeepest(pd).stackTrace.String()
}
