package image_ecosystem

import "fmt"

type ImageBaseType string
type tc struct {
	Version			string
	Cmd			string
	Expected		string
	Repository		string
	DockerImageReference	string
}

var s2iImages = map[string][]tc{"ruby": {{Version: "22", Cmd: "ruby --version", Expected: "ruby 2.2", Repository: "centos"}, {Version: "23", Cmd: "ruby --version", Expected: "ruby 2.3", Repository: "centos"}, {Version: "24", Cmd: "ruby --version", Expected: "ruby 2.4", Repository: "centos"}}, "python": {{Version: "27", Cmd: "python --version", Expected: "Python 2.7", Repository: "centos"}, {Version: "34", Cmd: "python --version", Expected: "Python 3.4", Repository: "centos"}, {Version: "35", Cmd: "python --version", Expected: "Python 3.5", Repository: "centos"}, {Version: "36", Cmd: "python --version", Expected: "Python 3.6", Repository: "centos"}}, "nodejs": {{Version: "4", Cmd: "node --version", Expected: "v4", Repository: "centos"}, {Version: "6", Cmd: "node --version", Expected: "v6", Repository: "centos"}}, "perl": {{Version: "520", Cmd: "perl --version", Expected: "v5.20", Repository: "centos"}, {Version: "524", Cmd: "perl --version", Expected: "v5.24", Repository: "centos"}}, "php": {{Version: "56", Cmd: "php --version", Expected: "5.6", Repository: "centos"}, {Version: "70", Cmd: "php --version", Expected: "7.0", Repository: "centos"}}}

func GetTestCaseForImages() map[string][]tc {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result := make(map[string][]tc)
	for name, variants := range s2iImages {
		for i := range variants {
			resolveDockerImageReference(name, &variants[i])
			result[name] = append(result[name], variants[i])
		}
	}
	return result
}
func resolveDockerImageReference(name string, t *tc) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(t.Repository) == 0 {
		t.Repository = "openshift"
	}
	t.DockerImageReference = fmt.Sprintf("%s/%s-%s-centos7", t.Repository, name, t.Version)
}
