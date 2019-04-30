package release

import (
	"encoding/json"
	"fmt"
	"time"
)

func createReleaseSignatureMessage(signer string, now time.Time, releaseDigest, pullSpec string) ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(signer) == 0 || now.IsZero() || len(releaseDigest) == 0 || len(pullSpec) == 0 {
		return nil, fmt.Errorf("you must specify a signer, current timestamp, release digest, and pull spec to sign")
	}
	sig := &signature{Critical: criticalSignature{Type: "atomic container signature", Image: criticalImage{DockerManifestDigest: releaseDigest}, Identity: criticalIdentity{DockerReference: pullSpec}}, Optional: optionalSignature{Creator: signer, Timestamp: now.Unix()}}
	return json.MarshalIndent(sig, "", "  ")
}

type signature struct {
	Critical	criticalSignature	`json:"critical"`
	Optional	optionalSignature	`json:"optional"`
}
type criticalSignature struct {
	Type		string			`json:"type"`
	Image		criticalImage		`json:"image"`
	Identity	criticalIdentity	`json:"identity"`
}
type criticalImage struct {
	DockerManifestDigest string `json:"docker-manifest-digest"`
}
type criticalIdentity struct {
	DockerReference string `json:"docker-reference"`
}
type optionalSignature struct {
	Creator		string	`json:"creator"`
	Timestamp	int64	`json:"timestamp"`
}
