package logger

import (
	"bytes"
	"os"
)

// OutputSplitter custo logrus log
type OutputSplitter struct{}

// Write write
func (splitter *OutputSplitter) Write(p []byte) (n int, err error) {
	if bytes.Contains(p, []byte("level=error")) {
		return os.Stderr.Write(p)
	}
	return os.Stdout.Write(p)
}
