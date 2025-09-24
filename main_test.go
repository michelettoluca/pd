package pd

import (
	"errors"
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	err := wrapError()
	err = Status(10).Code("nuovo codice").Wrapf(err, "nuovo errore")

	var pdErr Error
	if errors.As(err, &pdErr) {
		fmt.Printf("Message: %s\n", pdErr.Message())
		fmt.Printf("Status: %d\n", pdErr.Status())
		fmt.Printf("Code: %s\n", pdErr.Code())
		fmt.Printf("Stacktrace: %s\n", pdErr.StackTrace())
	}
}

func newError() error {
	return Status(420).Code("codice originale").New("errore originale")
}

func wrapError() error {
	err := newError()
	return Status(421).Code("codice originale 2").Wrapf(err, "errore originale 2")
}
