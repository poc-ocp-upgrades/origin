package dockerfile

import (
	"bytes"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"strings"
	"text/scanner"
	"unicode"
)

type ShellLex struct{ escapeToken rune }

func NewShellLex(escapeToken rune) *ShellLex {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &ShellLex{escapeToken: escapeToken}
}
func (s *ShellLex) ProcessWord(word string, env []string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	word, _, err := s.process(word, env)
	return word, err
}
func (s *ShellLex) ProcessWords(word string, env []string) ([]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, words, err := s.process(word, env)
	return words, err
}
func (s *ShellLex) process(word string, env []string) (string, []string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	sw := &shellWord{envs: env, escapeToken: s.escapeToken}
	sw.scanner.Init(strings.NewReader(word))
	return sw.process(word)
}

type shellWord struct {
	scanner     scanner.Scanner
	envs        []string
	escapeToken rune
}

func (sw *shellWord) process(source string) (string, []string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	word, words, err := sw.processStopOn(scanner.EOF)
	if err != nil {
		err = errors.Wrapf(err, "failed to process %q", source)
	}
	return word, words, err
}

type wordsStruct struct {
	word   string
	words  []string
	inWord bool
}

func (w *wordsStruct) addChar(ch rune) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if unicode.IsSpace(ch) && w.inWord {
		if len(w.word) != 0 {
			w.words = append(w.words, w.word)
			w.word = ""
			w.inWord = false
		}
	} else if !unicode.IsSpace(ch) {
		w.addRawChar(ch)
	}
}
func (w *wordsStruct) addRawChar(ch rune) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	w.word += string(ch)
	w.inWord = true
}
func (w *wordsStruct) addString(str string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var scan scanner.Scanner
	scan.Init(strings.NewReader(str))
	for scan.Peek() != scanner.EOF {
		w.addChar(scan.Next())
	}
}
func (w *wordsStruct) addRawString(str string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	w.word += str
	w.inWord = true
}
func (w *wordsStruct) getWords() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(w.word) > 0 {
		w.words = append(w.words, w.word)
		w.word = ""
		w.inWord = false
	}
	return w.words
}
func (sw *shellWord) processStopOn(stopChar rune) (string, []string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var result bytes.Buffer
	var words wordsStruct
	var charFuncMapping = map[rune]func() (string, error){'\'': sw.processSingleQuote, '"': sw.processDoubleQuote, '$': sw.processDollar}
	for sw.scanner.Peek() != scanner.EOF {
		ch := sw.scanner.Peek()
		if stopChar != scanner.EOF && ch == stopChar {
			sw.scanner.Next()
			break
		}
		if fn, ok := charFuncMapping[ch]; ok {
			tmp, err := fn()
			if err != nil {
				return "", []string{}, err
			}
			result.WriteString(tmp)
			if ch == rune('$') {
				words.addString(tmp)
			} else {
				words.addRawString(tmp)
			}
		} else {
			ch = sw.scanner.Next()
			if ch == sw.escapeToken {
				ch = sw.scanner.Next()
				if ch == scanner.EOF {
					break
				}
				words.addRawChar(ch)
			} else {
				words.addChar(ch)
			}
			result.WriteRune(ch)
		}
	}
	return result.String(), words.getWords(), nil
}
func (sw *shellWord) processSingleQuote() (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var result bytes.Buffer
	sw.scanner.Next()
	for {
		ch := sw.scanner.Next()
		switch ch {
		case scanner.EOF:
			return "", errors.New("unexpected end of statement while looking for matching single-quote")
		case '\'':
			return result.String(), nil
		}
		result.WriteRune(ch)
	}
}
func (sw *shellWord) processDoubleQuote() (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var result bytes.Buffer
	sw.scanner.Next()
	for {
		switch sw.scanner.Peek() {
		case scanner.EOF:
			return "", errors.New("unexpected end of statement while looking for matching double-quote")
		case '"':
			sw.scanner.Next()
			return result.String(), nil
		case '$':
			value, err := sw.processDollar()
			if err != nil {
				return "", err
			}
			result.WriteString(value)
		default:
			ch := sw.scanner.Next()
			if ch == sw.escapeToken {
				switch sw.scanner.Peek() {
				case scanner.EOF:
					continue
				case '"', '$', sw.escapeToken:
					ch = sw.scanner.Next()
				}
			}
			result.WriteRune(ch)
		}
	}
}
func (sw *shellWord) processDollar() (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	sw.scanner.Next()
	if sw.scanner.Peek() != '{' {
		name := sw.processName()
		if name == "" {
			return "$", nil
		}
		return sw.getEnv(name), nil
	}
	sw.scanner.Next()
	name := sw.processName()
	ch := sw.scanner.Peek()
	if ch == '}' {
		sw.scanner.Next()
		return sw.getEnv(name), nil
	}
	if ch == ':' {
		sw.scanner.Next()
		modifier := sw.scanner.Next()
		word, _, err := sw.processStopOn('}')
		if err != nil {
			return "", err
		}
		newValue := sw.getEnv(name)
		switch modifier {
		case '+':
			if newValue != "" {
				newValue = word
			}
			return newValue, nil
		case '-':
			if newValue == "" {
				newValue = word
			}
			return newValue, nil
		default:
			return "", errors.Errorf("unsupported modifier (%c) in substitution", modifier)
		}
	}
	return "", errors.Errorf("missing ':' in substitution")
}
func (sw *shellWord) processName() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var name bytes.Buffer
	for sw.scanner.Peek() != scanner.EOF {
		ch := sw.scanner.Peek()
		if name.Len() == 0 && unicode.IsDigit(ch) {
			ch = sw.scanner.Next()
			return string(ch)
		}
		if !unicode.IsLetter(ch) && !unicode.IsDigit(ch) && ch != '_' {
			break
		}
		ch = sw.scanner.Next()
		name.WriteRune(ch)
	}
	return name.String()
}
func (sw *shellWord) getEnv(name string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, env := range sw.envs {
		i := strings.Index(env, "=")
		if i < 0 {
			if equalEnvKeys(name, env) {
				return ""
			}
			continue
		}
		compareName := env[:i]
		if !equalEnvKeys(name, compareName) {
			continue
		}
		return env[i+1:]
	}
	return ""
}
func normalizeWorkdir(_ string, current string, requested string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if requested == "" {
		return "", errors.New("cannot normalize nothing")
	}
	current = filepath.FromSlash(current)
	requested = filepath.FromSlash(requested)
	if !filepath.IsAbs(requested) {
		return filepath.Join(string(os.PathSeparator), current, requested), nil
	}
	return requested, nil
}
func equalEnvKeys(from, to string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return from == to
}
