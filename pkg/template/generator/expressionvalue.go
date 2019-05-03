package generator

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

type ExpressionValueGenerator struct{ seed *rand.Rand }

const (
	Alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Numerals = "0123456789"
	Symbols  = "~!@#$%^&*()-_+={}[]\\|<,>.?/\"';:`"
	ASCII    = Alphabet + Numerals + Symbols
)

var (
	rangeExp      = regexp.MustCompile(`([\\]?[a-zA-Z0-9]\-?[a-zA-Z0-9]?)`)
	generatorsExp = regexp.MustCompile(`\[([a-zA-Z0-9\-\\]+)\](\{([0-9]+)\})`)
	expressionExp = regexp.MustCompile(`\[(\\w|\\d|\\a|\\A)|([a-zA-Z0-9]\-[a-zA-Z0-9])+\]`)
)

func NewExpressionValueGenerator(seed *rand.Rand) ExpressionValueGenerator {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ExpressionValueGenerator{seed: seed}
}
func (g ExpressionValueGenerator) GenerateValue(expression string) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for {
		r := generatorsExp.FindStringIndex(expression)
		if r == nil {
			break
		}
		ranges, length, err := rangesAndLength(expression[r[0]:r[1]])
		if err != nil {
			return "", err
		}
		err = replaceWithGenerated(&expression, expression[r[0]:r[1]], findExpressionPos(ranges), length, g.seed)
		if err != nil {
			return "", err
		}
	}
	return expression, nil
}
func alphabetSlice(from, to byte) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	leftPos := strings.Index(ASCII, string(from))
	rightPos := strings.LastIndex(ASCII, string(to))
	if leftPos > rightPos {
		return "", fmt.Errorf("invalid range specified: %s-%s", string(from), string(to))
	}
	return ASCII[leftPos:rightPos], nil
}
func replaceWithGenerated(s *string, expression string, ranges [][]byte, length int, seed *rand.Rand) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var alphabet string
	for _, r := range ranges {
		switch string(r[0]) + string(r[1]) {
		case `\w`:
			alphabet += Alphabet + Numerals + "_"
		case `\d`:
			alphabet += Numerals
		case `\a`:
			alphabet += Alphabet + Numerals
		case `\A`:
			alphabet += Symbols
		default:
			slice, err := alphabetSlice(r[0], r[1])
			if err != nil {
				return err
			}
			alphabet += slice
		}
	}
	result := make([]byte, length)
	alphabet = removeDuplicateChars(alphabet)
	for i := 0; i < length; i++ {
		result[i] = alphabet[seed.Intn(len(alphabet))]
	}
	*s = strings.Replace(*s, expression, string(result), 1)
	return nil
}
func removeDuplicateChars(input string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	data := []byte(input)
	length := len(data) - 1
	for i := 0; i < length; i++ {
		for j := i + 1; j <= length; j++ {
			if data[i] == data[j] {
				data[j] = data[length]
				data = data[0:length]
				length--
				j--
			}
		}
	}
	return string(data)
}
func findExpressionPos(s string) [][]byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	matches := rangeExp.FindAllStringIndex(s, -1)
	result := make([][]byte, len(matches))
	for i, r := range matches {
		result[i] = []byte{s[r[0]], s[r[1]-1]}
	}
	return result
}
func rangesAndLength(s string) (string, int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	expr := s[0:strings.LastIndex(s, "{")]
	if !expressionExp.MatchString(expr) {
		return "", 0, fmt.Errorf("malformed expresion syntax: %s", expr)
	}
	length, _ := strconv.Atoi(s[strings.LastIndex(s, "{")+1 : len(s)-1])
	if length > 0 && length <= 255 {
		return expr, length, nil
	}
	return "", 0, fmt.Errorf("range must be within [1-255] characters (%d)", length)
}
