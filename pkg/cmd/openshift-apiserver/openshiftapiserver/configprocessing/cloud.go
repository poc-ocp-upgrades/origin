package configprocessing

import "fmt"

func GetCloudProviderConfigFile(args map[string][]string) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	filenames, ok := args["cloud-config"]
	if !ok {
		return "", nil
	}
	if len(filenames) != 1 {
		return "", fmt.Errorf(`one or zero "--cloud-config" required, not %v`, filenames)
	}
	return filenames[0], nil
}
