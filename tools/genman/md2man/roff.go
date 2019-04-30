package md2man

import (
	"bytes"
	"fmt"
	"html"
	"strings"
	"github.com/russross/blackfriday"
)

type roffRenderer struct{}

func RoffRenderer(flags int) blackfriday.Renderer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &roffRenderer{}
}
func (r *roffRenderer) GetFlags() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return 0
}
func (r *roffRenderer) TitleBlock(out *bytes.Buffer, text []byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.WriteString(".TH ")
	splitText := bytes.Split(text, []byte("\n"))
	for i, line := range splitText {
		line = bytes.TrimPrefix(line, []byte("% "))
		if i == 0 {
			line = bytes.Replace(line, []byte("("), []byte("\" \""), 1)
			line = bytes.Replace(line, []byte(")"), []byte("\" \""), 1)
		}
		line = append([]byte("\""), line...)
		line = append(line, []byte("\" ")...)
		out.Write(line)
	}
	out.WriteString(" \"\"\n")
}
func (r *roffRenderer) BlockCode(out *bytes.Buffer, text []byte, lang string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.WriteString("\n.PP\n.RS\n\n.nf\n")
	escapeSpecialChars(out, text)
	out.WriteString("\n.fi\n.RE\n")
}
func (r *roffRenderer) BlockQuote(out *bytes.Buffer, text []byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.WriteString("\n.PP\n.RS\n")
	out.Write(text)
	out.WriteString("\n.RE\n")
}
func (r *roffRenderer) BlockHtml(out *bytes.Buffer, text []byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Write(text)
}
func (r *roffRenderer) Header(out *bytes.Buffer, text func() bool, level int, id string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	marker := out.Len()
	switch {
	case marker == 0:
		out.WriteString(".TH ")
	case level == 1:
		out.WriteString("\n\n.SH ")
	case level == 2:
		out.WriteString("\n.SH ")
	default:
		out.WriteString("\n.SS ")
	}
	if !text() {
		out.Truncate(marker)
		return
	}
}
func (r *roffRenderer) HRule(out *bytes.Buffer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.WriteString("\n.ti 0\n\\l'\\n(.lu'\n")
}
func (r *roffRenderer) List(out *bytes.Buffer, text func() bool, flags int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	marker := out.Len()
	out.WriteString(".IP ")
	if flags&blackfriday.LIST_TYPE_ORDERED != 0 {
		out.WriteString("\\(bu 2")
	}
	out.WriteString("\n")
	if !text() {
		out.Truncate(marker)
		return
	}
}
func (r *roffRenderer) ListItem(out *bytes.Buffer, text []byte, flags int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.WriteString("\n.IP\n\\(bu ")
	out.Write(text)
}
func (r *roffRenderer) Paragraph(out *bytes.Buffer, text func() bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	marker := out.Len()
	out.WriteString("\n.PP\n")
	if !text() {
		out.Truncate(marker)
		return
	}
	if marker != 0 {
		out.WriteString("\n")
	}
}
func (r *roffRenderer) Table(out *bytes.Buffer, header []byte, body []byte, columnData []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.WriteString(".TS\nallbox;\n")
	out.Write(header)
	out.Write(body)
	out.WriteString("\n.TE\n")
}
func (r *roffRenderer) TableRow(out *bytes.Buffer, text []byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if out.Len() > 0 {
		out.WriteString("\n")
	}
	out.Write(text)
	out.WriteString("\n")
}
func (r *roffRenderer) TableHeaderCell(out *bytes.Buffer, text []byte, align int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if out.Len() > 0 {
		out.WriteString(" ")
	}
	out.Write(text)
	out.WriteString(" ")
}
func (r *roffRenderer) TableCell(out *bytes.Buffer, text []byte, align int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if out.Len() > 0 {
		out.WriteString("\t")
	}
	out.Write(text)
	out.WriteString("\t")
}
func (r *roffRenderer) Footnotes(out *bytes.Buffer, text func() bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (r *roffRenderer) FootnoteItem(out *bytes.Buffer, name, text []byte, flags int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (r *roffRenderer) AutoLink(out *bytes.Buffer, link []byte, kind int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.WriteString("\n\\[la]")
	out.Write(link)
	out.WriteString("\\[ra]")
}
func (r *roffRenderer) CodeSpan(out *bytes.Buffer, text []byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.WriteString("\\fB\\fC")
	escapeSpecialChars(out, text)
	out.WriteString("\\fR")
}
func (r *roffRenderer) DoubleEmphasis(out *bytes.Buffer, text []byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.WriteString("\\fB")
	out.Write(text)
	out.WriteString("\\fP")
}
func (r *roffRenderer) Emphasis(out *bytes.Buffer, text []byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.WriteString("\\fI")
	out.Write(text)
	out.WriteString("\\fP")
}
func (r *roffRenderer) Image(out *bytes.Buffer, link []byte, title []byte, alt []byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (r *roffRenderer) LineBreak(out *bytes.Buffer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.WriteString("\n.br\n")
}
func (r *roffRenderer) Link(out *bytes.Buffer, link []byte, title []byte, content []byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r.AutoLink(out, link, 0)
}
func (r *roffRenderer) RawHtmlTag(out *bytes.Buffer, tag []byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Write(tag)
}
func (r *roffRenderer) TripleEmphasis(out *bytes.Buffer, text []byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.WriteString("\\s+2")
	out.Write(text)
	out.WriteString("\\s-2")
}
func (r *roffRenderer) StrikeThrough(out *bytes.Buffer, text []byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (r *roffRenderer) FootnoteRef(out *bytes.Buffer, ref []byte, id int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (r *roffRenderer) Entity(out *bytes.Buffer, entity []byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.WriteString(html.UnescapeString(string(entity)))
}
func processFooterText(text []byte) []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	text = bytes.TrimPrefix(text, []byte("% "))
	newText := []byte{}
	textArr := strings.Split(string(text), ") ")
	for i, w := range textArr {
		if i == 0 {
			w = strings.Replace(w, "(", "\" \"", 1)
			w = fmt.Sprintf("\"%s\"", w)
		} else {
			w = fmt.Sprintf(" \"%s\"", w)
		}
		newText = append(newText, []byte(w)...)
	}
	newText = append(newText, []byte(" \"\"")...)
	return newText
}
func (r *roffRenderer) NormalText(out *bytes.Buffer, text []byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	escapeSpecialChars(out, text)
}
func (r *roffRenderer) DocumentHeader(out *bytes.Buffer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (r *roffRenderer) DocumentFooter(out *bytes.Buffer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func needsBackslash(c byte) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, r := range []byte("-_&\\~") {
		if c == r {
			return true
		}
	}
	return false
}
func escapeSpecialChars(out *bytes.Buffer, text []byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := 0; i < len(text); i++ {
		org := i
		for i < len(text) && !needsBackslash(text[i]) {
			i++
		}
		if i > org {
			out.Write(text[org:i])
		}
		if i >= len(text) {
			break
		}
		out.WriteByte('\\')
		out.WriteByte(text[i])
	}
}
