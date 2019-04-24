package integration

import (
	"reflect"
	"strings"
	"testing"
	"time"
	dockerregistry "github.com/openshift/origin/pkg/image/importer/dockerv1client"
)

const (
	pulpRegistryName	= "registry.access.redhat.com"
	dockerHubV2RegistryName	= "index.docker.io"
	dockerHubV1RegistryName	= "registry.hub.docker.com"
	quayRegistryName	= "quay.io"
	maxRetryCount		= 4
	retryAfter		= time.Millisecond * 500
)

var (
	unreachableErrorPatterns	= []string{"dial tcp", "read tcp", "net/http", "could not be reached"}
	imageNotFoundErrorPatterns	= []string{"was not found and may have been deleted", "has not been set on repository"}
)

func retryOnErrors(t *testing.T, errorPatterns []string, f func() error) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	timeout := retryAfter
	attempt := 0
	for err := f(); err != nil; err = f() {
		match := false
		for _, pattern := range errorPatterns {
			if strings.Contains(err.Error(), pattern) {
				match = true
				break
			}
		}
		if !match || attempt >= maxRetryCount {
			return err
		}
		t.Logf("caught error \"%v\", retrying in %s", err, timeout.String())
		time.Sleep(timeout)
		timeout = timeout * 2
		attempt += 1
	}
	return nil
}
func retryWhenUnreachable(t *testing.T, f func() error, errorPatterns ...string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return retryOnErrors(t, append(errorPatterns, unreachableErrorPatterns...), f)
}
func TestRegistryClientConnect(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c := dockerregistry.NewClient(10*time.Second, true)
	conn, err := c.Connect("docker.io", false)
	if err != nil {
		t.Fatal(err)
	}
	for _, s := range []string{"index.docker.io", "https://docker.io", "https://index.docker.io"} {
		otherConn, err := c.Connect(s, false)
		if err != nil {
			t.Errorf("%s: can't connect: %v", s, err)
			continue
		}
		if !reflect.DeepEqual(otherConn, conn) {
			t.Errorf("%s: did not reuse connection: %#v %#v", s, conn, otherConn)
		}
	}
	otherConn, err := c.Connect("index.docker.io:443", false)
	if err != nil || reflect.DeepEqual(otherConn, conn) {
		t.Errorf("should not have reused index.docker.io:443: %v", err)
	}
	if _, err := c.Connect("http://ba%3/", false); err == nil {
		t.Error("Unexpected non-error")
	}
}
func TestRegistryClientConnectPulpRegistry(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c := dockerregistry.NewClient(10*time.Second, true)
	conn, err := c.Connect(pulpRegistryName, false)
	if err != nil {
		t.Fatal(err)
	}
	var image *dockerregistry.Image
	err = retryWhenUnreachable(t, func() error {
		image, err = conn.ImageByTag("library", "rhel", "latest")
		return err
	}, imageNotFoundErrorPatterns...)
	if err != nil {
		if strings.Contains(err.Error(), "x509: certificate has expired or is not yet valid") {
			t.Skipf("SKIPPING: due to expired certificate of %s: %v", pulpRegistryName, err)
		}
		t.Skip("pulp is failing")
	}
	if len(image.Image.ID) == 0 {
		t.Fatalf("image had no ID: %#v", image)
	}
}
func TestRegistryClientDockerHubV2(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c := dockerregistry.NewClient(10*time.Second, true)
	conn, err := c.Connect(dockerHubV2RegistryName, false)
	if err != nil {
		t.Fatal(err)
	}
	var image *dockerregistry.Image
	err = retryWhenUnreachable(t, func() error {
		image, err = conn.ImageByTag("kubernetes", "guestbook", "latest")
		return err
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(image.Image.ID) == 0 {
		t.Fatalf("image had no ID: %#v", image)
	}
}
func TestRegistryClientDockerHubV1(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c := dockerregistry.NewClient(10*time.Second, true)
	conn, err := c.Connect(dockerHubV1RegistryName, false)
	if err != nil {
		t.Fatal(err)
	}
	var image *dockerregistry.Image
	err = retryWhenUnreachable(t, func() error {
		image, err = conn.ImageByTag("kubernetes", "guestbook", "latest")
		return err
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(image.Image.ID) == 0 {
		t.Fatalf("image had no ID: %#v", image)
	}
}
func TestRegistryClientRegistryNotFound(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	conn, err := dockerregistry.NewClient(10*time.Second, true).Connect("localhost:65000", false)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := conn.ImageByID("foo", "bar", "baz"); !dockerregistry.IsRegistryNotFound(err) {
		t.Error(err)
	}
}
func doTestRegistryClientImage(t *testing.T, registry, reponame, version string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	conn, err := dockerregistry.NewClient(10*time.Second, version == "v2").Connect(registry, false)
	if err != nil {
		t.Fatal(err)
	}
	err = retryWhenUnreachable(t, func() error {
		_, err := conn.ImageByTag("openshift", "origin-not-found", "latest")
		return err
	})
	if err == nil || (!dockerregistry.IsRepositoryNotFound(err) && !dockerregistry.IsTagNotFound(err)) {
		t.Errorf("%s: unexpected error: %v", version, err)
	}
	var image *dockerregistry.Image
	err = retryWhenUnreachable(t, func() error {
		image, err = conn.ImageByTag("openshift", reponame, "latest")
		return err
	})
	if err != nil {
		t.Fatal(err)
	}
	if image.Image.Comment != "Imported from -" {
		t.Errorf("%s: unexpected image comment", version)
	}
	if image.Image.Architecture != "amd64" {
		t.Errorf("%s: unexpected image architecture", version)
	}
	if version == "v2" && !image.PullByID {
		t.Errorf("%s: should be able to pull by ID %s", version, image.Image.ID)
	}
	var other *dockerregistry.Image
	err = retryWhenUnreachable(t, func() error {
		other, err = conn.ImageByID("openshift", reponame, image.Image.ID)
		return err
	})
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(other.Image.ContainerConfig.Entrypoint, image.Image.ContainerConfig.Entrypoint) {
		t.Errorf("%s: unexpected image: %#v", version, other)
	}
}
func TestRegistryClientAPIv2ManifestV2Schema2(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.Log("openshift/schema-v2-test-repo was pushed by Docker 1.11.1")
	doTestRegistryClientImage(t, dockerHubV2RegistryName, "schema-v2-test-repo", "v2")
}
func TestRegistryClientAPIv2ManifestV2Schema1(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.Log("openshift/schema-v1-test-repo was pushed by Docker 1.8.2")
	doTestRegistryClientImage(t, dockerHubV2RegistryName, "schema-v1-test-repo", "v2")
}
func TestRegistryClientAPIv1(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.Log("openshift/schema-v1-test-repo was pushed by Docker 1.8.2")
	doTestRegistryClientImage(t, dockerHubV1RegistryName, "schema-v1-test-repo", "v1")
}
func TestRegistryClientQuayIOImage(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	conn, err := dockerregistry.NewClient(10*time.Second, true).Connect("quay.io", false)
	if err != nil {
		t.Fatal(err)
	}
	err = retryWhenUnreachable(t, func() error {
		_, err := conn.ImageByTag("coreos", "etcd", "latest")
		return err
	}, imageNotFoundErrorPatterns...)
	if err != nil {
		t.Skipf("SKIPPING: unexpected error from quay.io: %v", err)
	}
}
