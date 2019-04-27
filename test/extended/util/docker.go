package util

import (
	"fmt"
	"strings"
	g "github.com/onsi/ginkgo"
	gson "encoding/json"
	dockerClient "github.com/fsouza/go-dockerclient"
	tutil "github.com/openshift/origin/test/util"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kubernetes/pkg/credentialprovider"
)

func ListImages() ([]string, error) {
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
	client, err := tutil.NewDockerClient()
	if err != nil {
		return nil, err
	}
	imageList, err := client.ListImages(dockerClient.ListImagesOptions{})
	if err != nil {
		return nil, err
	}
	returnIds := make([]string, 0)
	for _, image := range imageList {
		for _, tag := range image.RepoTags {
			returnIds = append(returnIds, tag)
		}
	}
	return returnIds, nil
}
func BuildAuthConfiguration(credKey string, oc *CLI) (*dockerClient.AuthConfiguration, error) {
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
	authCfg := &dockerClient.AuthConfiguration{}
	secretList, err := oc.AdminKubeClient().CoreV1().Secrets(oc.Namespace()).List(metav1.ListOptions{})
	g.By(fmt.Sprintf("get secret list err %v ", err))
	if err == nil {
		for _, secret := range secretList.Items {
			g.By(fmt.Sprintf("secret name %s ", secret.ObjectMeta.Name))
			if strings.Contains(secret.ObjectMeta.Name, "builder-dockercfg") {
				dockercfgToken := secret.Data[".dockercfg"]
				dockercfgTokenJson := string(dockercfgToken)
				g.By(fmt.Sprintf("docker cfg token json %s ", dockercfgTokenJson))
				creds := credentialprovider.DockerConfig{}
				err = gson.Unmarshal(dockercfgToken, &creds)
				g.By(fmt.Sprintf("json unmarshal err %v ", err))
				if err == nil {
					keyring := credentialprovider.BasicDockerKeyring{}
					keyring.Add(creds)
					authConfs, found := keyring.Lookup(credKey)
					g.By(fmt.Sprintf("found auth %v with auth cfg len %d ", found, len(authConfs)))
					if !found || len(authConfs) == 0 {
						return authCfg, err
					}
					if len(authConfs[0].ServerAddress) == 0 {
						authConfs[0].ServerAddress = credKey
					}
					g.By(fmt.Sprintf("dockercfg with svrAddr %s user %s pass %s email %s ", authConfs[0].ServerAddress, authConfs[0].Username, authConfs[0].Password, authConfs[0].Email))
					c := dockerClient.AuthConfiguration{Username: authConfs[0].Username, ServerAddress: authConfs[0].ServerAddress, Password: authConfs[0].Password}
					return &c, err
				}
			}
		}
	}
	return authCfg, err
}

type MissingTagError struct{ Tags []string }

func (mte MissingTagError) Error() string {
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
	return fmt.Sprintf("the tag %s passed in was invalid, and not found in the list of images returned from docker", mte.Tags)
}
func GetImageIDForTags(comps []string) ([]string, error) {
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
	client, dcerr := tutil.NewDockerClient()
	if dcerr != nil {
		return nil, dcerr
	}
	imageList, serr := client.ListImages(dockerClient.ListImagesOptions{})
	if serr != nil {
		return nil, serr
	}
	returnTags := make([]string, 0)
	missingTags := make([]string, 0)
	for _, comp := range comps {
		var found bool
		for _, image := range imageList {
			for _, repTag := range image.RepoTags {
				if repTag == comp {
					found = true
					returnTags = append(returnTags, image.ID)
					break
				}
			}
			if found {
				break
			}
		}
		if !found {
			returnTags = append(returnTags, "")
			missingTags = append(missingTags, comp)
		}
	}
	if len(missingTags) == 0 {
		return returnTags, nil
	} else {
		mte := MissingTagError{Tags: missingTags}
		return returnTags, mte
	}
}
