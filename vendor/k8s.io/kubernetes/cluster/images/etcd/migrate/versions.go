package main

import (
	"fmt"
	"github.com/blang/semver"
	"strings"
)

type EtcdVersion struct{ semver.Version }

func ParseEtcdVersion(s string) (*EtcdVersion, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	v, err := semver.Make(s)
	if err != nil {
		return nil, err
	}
	return &EtcdVersion{v}, nil
}
func MustParseEtcdVersion(s string) *EtcdVersion {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &EtcdVersion{semver.MustParse(s)}
}
func (v *EtcdVersion) String() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return v.Version.String()
}
func (v *EtcdVersion) Equals(o *EtcdVersion) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return v.Version.Equals(o.Version)
}
func (v *EtcdVersion) MajorMinorEquals(o *EtcdVersion) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return v.Major == o.Major && v.Minor == o.Minor
}

type EtcdStorageVersion int

const (
	storageUnknown EtcdStorageVersion = iota
	storageEtcd2
	storageEtcd3
)

func ParseEtcdStorageVersion(s string) (EtcdStorageVersion, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch s {
	case "etcd2":
		return storageEtcd2, nil
	case "etcd3":
		return storageEtcd3, nil
	default:
		return storageUnknown, fmt.Errorf("unrecognized storage version: %s", s)
	}
}
func MustParseEtcdStorageVersion(s string) EtcdStorageVersion {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	version, err := ParseEtcdStorageVersion(s)
	if err != nil {
		panic(err)
	}
	return version
}
func (v EtcdStorageVersion) String() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	switch v {
	case storageEtcd2:
		return "etcd2"
	case storageEtcd3:
		return "etcd3"
	default:
		panic(fmt.Sprintf("enum value %d missing from EtcdStorageVersion String() function", v))
	}
}

type EtcdVersionPair struct {
	version        *EtcdVersion
	storageVersion EtcdStorageVersion
}

func ParseEtcdVersionPair(s string) (*EtcdVersionPair, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	parts := strings.Split(s, "/")
	if len(parts) != 2 {
		return nil, fmt.Errorf("Malformed version file, expected <major>.<minor>.<patch>/<storage> but got %s", s)
	}
	version, err := ParseEtcdVersion(parts[0])
	if err != nil {
		return nil, err
	}
	storageVersion, err := ParseEtcdStorageVersion(parts[1])
	if err != nil {
		return nil, err
	}
	return &EtcdVersionPair{version, storageVersion}, nil
}
func MustParseEtcdVersionPair(s string) *EtcdVersionPair {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	pair, err := ParseEtcdVersionPair(s)
	if err != nil {
		panic(err)
	}
	return pair
}
func (vp *EtcdVersionPair) String() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fmt.Sprintf("%s/%s", vp.version, vp.storageVersion)
}
func (vp *EtcdVersionPair) Equals(o *EtcdVersionPair) bool {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return vp.version.Equals(o.version) && vp.storageVersion == o.storageVersion
}

type SupportedVersions []*EtcdVersion

func (sv SupportedVersions) NextVersion(current *EtcdVersion) *EtcdVersion {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var nextVersion *EtcdVersion
	for i, supportedVersion := range sv {
		if current.MajorMinorEquals(supportedVersion) && len(sv) > i+1 {
			nextVersion = sv[i+1]
		}
	}
	return nextVersion
}
func (sv SupportedVersions) NextVersionPair(current *EtcdVersionPair) *EtcdVersionPair {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	nextVersion := sv.NextVersion(current.version)
	if nextVersion == nil {
		return nil
	}
	storageVersion := storageEtcd3
	if nextVersion.Major == 2 {
		storageVersion = storageEtcd2
	}
	return &EtcdVersionPair{version: nextVersion, storageVersion: storageVersion}
}
func ParseSupportedVersions(s string) (SupportedVersions, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	var err error
	list := strings.Split(s, ",")
	versions := make(SupportedVersions, len(list))
	for i, v := range list {
		versions[i], err = ParseEtcdVersion(strings.TrimSpace(v))
		if err != nil {
			return nil, err
		}
	}
	return versions, nil
}
func MustParseSupportedVersions(s string) SupportedVersions {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	versions, err := ParseSupportedVersions(s)
	if err != nil {
		panic(err)
	}
	return versions
}
