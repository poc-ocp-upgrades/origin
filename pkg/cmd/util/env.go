package util

import (
	"fmt"
	"os"
	"strconv"
)

func EnvInt(key string, defaultValue int32, minValue int32) int32 {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	value, err := strconv.ParseInt(Env(key, fmt.Sprintf("%d", defaultValue)), 10, 32)
	if err != nil || int32(value) < minValue {
		return defaultValue
	}
	return int32(value)
}
func Env(key string, defaultValue string) string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	val := os.Getenv(key)
	if len(val) == 0 {
		return defaultValue
	}
	return val
}
func GetEnv(key string) (string, bool) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	val := os.Getenv(key)
	if len(val) == 0 {
		return "", false
	}
	return val, true
}
