package logger

import (
	"fmt"
	"os"
)

type consoleWriter struct {
}

func (f *consoleWriter) Write(data interface{}) error {
	str := fmt.Sprintf("%v\n", data)

	_, err := os.Stdout.Write([]byte(str))

	return err
}

func newConsoleWriter() *consoleWriter {
	return &consoleWriter{}
}
