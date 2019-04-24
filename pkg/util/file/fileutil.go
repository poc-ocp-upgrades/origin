package file

import (
	"bufio"
	"io/ioutil"
	"os"
)

func ReadLines(fileName string) ([]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
func LoadData(file string) ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(file) == 0 {
		return []byte{}, nil
	}
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return []byte{}, err
	}
	return bytes, nil
}
