package customresourcevalidation

func RequireNameCluster(name string, prefix bool) []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if name != "cluster" {
		return []string{"must be cluster"}
	}
	return nil
}
