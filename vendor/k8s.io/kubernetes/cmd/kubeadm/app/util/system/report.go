package system

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os"
)

type ValidationResultType int32

const (
	good ValidationResultType = iota
	bad
	warn
)

type color int32

const (
	red    color = 31
	green        = 32
	yellow       = 33
	white        = 37
)

func colorize(s string, c color) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fmt.Sprintf("\033[0;%dm%s\033[0m", c, s)
}

type StreamReporter struct{ WriteStream io.Writer }

func (dr *StreamReporter) Report(key, value string, resultType ValidationResultType) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var c color
	switch resultType {
	case good:
		c = green
	case bad:
		c = red
	case warn:
		c = yellow
	default:
		c = white
	}
	if dr.WriteStream == nil {
		return errors.New("WriteStream has to be defined for this reporter")
	}
	fmt.Fprintf(dr.WriteStream, "%s: %s\n", colorize(key, white), colorize(value, c))
	return nil
}

var DefaultReporter = &StreamReporter{WriteStream: os.Stdout}
