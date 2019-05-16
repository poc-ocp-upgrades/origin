package strings

import (
	goformat "fmt"
	goos "os"
	godefaultruntime "runtime"
	gotime "time"
)

func IsWildcardMatch(s string, p string) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	dp := make([][]bool, len(p)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s)+1)
		for j := range dp[i] {
			dp[i][j] = false
		}
	}
	dp[0][0] = true
	for j := 1; j <= len(p); j++ {
		pattern := p[j-1]
		dp[j][0] = dp[j-1][0] && pattern == '*'
		for i := 1; i <= len(s); i++ {
			letter := s[i-1]
			if pattern != '*' {
				dp[j][i] = dp[j-1][i-1] && (pattern == '?' || pattern == letter)
			} else {
				dp[j][i] = dp[j][i-1] || dp[j-1][i]
			}
		}
	}
	return dp[len(p)][len(s)]
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
