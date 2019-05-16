package main

import (
	"fmt"
	goformat "fmt"
	"github.com/mvdan/xurls"
	flag "github.com/spf13/pflag"
	"io/ioutil"
	"net/http"
	"os"
	goos "os"
	"path/filepath"
	"regexp"
	godefaultruntime "runtime"
	"strconv"
	"strings"
	"time"
	gotime "time"
)

var (
	rootDir          = flag.String("root-dir", "", "Root directory containing documents to be processed.")
	fileSuffix       = flag.StringSlice("file-suffix", []string{"types.go", ".md"}, "suffix of files to be checked")
	regWhiteList     = []*regexp.Regexp{regexp.MustCompile(`https://kubernetes-site\.appspot\.com`), regexp.MustCompile(`https?://[^A-Za-z].*`), regexp.MustCompile(`https?://localhost.*`)}
	fullURLWhiteList = map[string]struct{}{"http://github.com/some/repo.git": {}, "http://stackoverflow.com/questions/ask?tags=kubernetes": {}, "https://github.com/$YOUR_GITHUB_USERNAME/kubernetes.git": {}, "https://github.com/$YOUR_GITHUB_USERNAME/kubernetes": {}, "http://storage.googleapis.com/kubernetes-release/release/v${K8S_VERSION}/bin/darwin/amd64/kubectl": {}, "http://supervisord.org/": {}, "http://kubernetes.io/vX.Y/docs": {}, "http://kubernetes.io/vX.Y/docs/": {}, "http://kubernetes.io/vX.Y/": {}}
	visitedURLs      = map[string]struct{}{}
	htmlpreviewReg   = regexp.MustCompile(`https://htmlpreview\.github\.io/\?`)
	httpOrhttpsReg   = regexp.MustCompile(`https?.*`)
)

func newWalkFunc(invalidLink *bool, client *http.Client) filepath.WalkFunc {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return func(filePath string, info os.FileInfo, err error) error {
		hasSuffix := false
		for _, suffix := range *fileSuffix {
			hasSuffix = hasSuffix || strings.HasSuffix(info.Name(), suffix)
		}
		if !hasSuffix {
			return nil
		}
		fileBytes, err := ioutil.ReadFile(filePath)
		if err != nil {
			return err
		}
		foundInvalid := false
		allURLs := xurls.Strict.FindAll(fileBytes, -1)
		fmt.Fprintf(os.Stdout, "\nChecking file %s\n", filePath)
	URL:
		for _, URL := range allURLs {
			if !httpOrhttpsReg.Match(URL) {
				continue
			}
			for _, whiteURL := range regWhiteList {
				if whiteURL.Match(URL) {
					continue URL
				}
			}
			if _, found := fullURLWhiteList[string(URL)]; found {
				continue
			}
			processedURL := htmlpreviewReg.ReplaceAll(URL, []byte{})
			if _, found := visitedURLs[string(processedURL)]; found {
				continue
			}
			visitedURLs[string(processedURL)] = struct{}{}
			retry := 0
			const maxRetry int = 3
			backoff := 100
			for retry < maxRetry {
				fmt.Fprintf(os.Stdout, "Visiting %s\n", string(processedURL))
				resp, err := client.Head(string(processedURL))
				if err != nil {
					break
				}
				if resp.StatusCode == http.StatusTooManyRequests {
					retryAfter := resp.Header.Get("Retry-After")
					if seconds, err := strconv.Atoi(retryAfter); err != nil {
						backoff = seconds + 10
					}
					fmt.Fprintf(os.Stderr, "Got %d visiting %s, retry after %d seconds.\n", resp.StatusCode, string(URL), backoff)
					time.Sleep(time.Duration(backoff) * time.Second)
					backoff *= 2
					retry++
				} else if resp.StatusCode == http.StatusNotFound {
					resp, err = client.Get(string(processedURL))
					if err != nil {
						break
					}
					if resp.StatusCode != http.StatusNotFound {
						continue URL
					}
					foundInvalid = true
					fmt.Fprintf(os.Stderr, "Failed: in file %s, Got %d visiting %s\n", filePath, resp.StatusCode, string(URL))
					break
				} else {
					break
				}
			}
			if retry == maxRetry {
				foundInvalid = true
				fmt.Fprintf(os.Stderr, "Failed: in file %s, still got 429 visiting %s after %d retries\n", filePath, string(URL), maxRetry)
			}
		}
		if foundInvalid {
			*invalidLink = true
		}
		return nil
	}
}
func main() {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	flag.Parse()
	if *rootDir == "" {
		flag.Usage()
		os.Exit(2)
	}
	client := http.Client{Timeout: time.Duration(5 * time.Second)}
	invalidLink := false
	if err := filepath.Walk(*rootDir, newWalkFunc(&invalidLink, &client)); err != nil {
		fmt.Fprintf(os.Stderr, "Fail: %v.\n", err)
		os.Exit(2)
	}
	if invalidLink {
		os.Exit(1)
	}
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
