package term

import (
	"io"
	"os"
	"github.com/docker/docker/pkg/term"
	wordwrap "github.com/mitchellh/go-wordwrap"
	kterm "k8s.io/kubernetes/pkg/kubectl/util/term"
)

type wordWrapWriter struct {
	limit	uint
	writer	io.Writer
}

func NewResponsiveWriter(w io.Writer) io.Writer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	file, ok := w.(*os.File)
	if !ok {
		return w
	}
	fd := file.Fd()
	if !term.IsTerminal(fd) {
		return w
	}
	terminalSize := kterm.GetSize(fd)
	if terminalSize == nil {
		return w
	}
	var limit uint
	switch {
	case terminalSize.Width >= 120:
		limit = 120
	case terminalSize.Width >= 100:
		limit = 100
	case terminalSize.Width >= 80:
		limit = 80
	}
	return NewWordWrapWriter(w, limit)
}
func NewWordWrapWriter(w io.Writer, limit uint) io.Writer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &wordWrapWriter{limit: limit, writer: w}
}
func (w wordWrapWriter) Write(p []byte) (nn int, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if w.limit == 0 {
		return w.writer.Write(p)
	}
	original := string(p)
	wrapped := wordwrap.WrapString(original, w.limit)
	l, e := w.writer.Write([]byte(wrapped))
	if e == nil || l > len(original) {
		l = len(original)
	}
	return l, e
}
func NewPunchCardWriter(w io.Writer) io.Writer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return NewWordWrapWriter(w, 80)
}

type maxWidthWriter struct {
	maxWidth	uint
	currentWidth	uint
	written		uint
	writer		io.Writer
}

func NewMaxWidthWriter(w io.Writer, maxWidth uint) io.Writer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &maxWidthWriter{maxWidth: maxWidth, writer: w}
}
func (m maxWidthWriter) Write(p []byte) (nn int, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, b := range p {
		if m.currentWidth == m.maxWidth {
			m.writer.Write([]byte{'\n'})
			m.currentWidth = 0
		}
		if b == '\n' {
			m.currentWidth = 0
		}
		_, err := m.writer.Write([]byte{b})
		if err != nil {
			return int(m.written), err
		}
		m.written++
		m.currentWidth++
	}
	return len(p), nil
}
