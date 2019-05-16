package file

import (
	"bufio"
	goformat "fmt"
	"io/ioutil"
	"os"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func ReadLines(fileName string) ([]string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
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
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if len(file) == 0 {
		return []byte{}, nil
	}
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return []byte{}, err
	}
	return bytes, nil
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
