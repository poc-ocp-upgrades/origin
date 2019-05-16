package system

type KernelConfig struct {
	Name        string   `json:"name,omitempty"`
	Aliases     []string `json:"aliases,omitempty"`
	Description string   `json:"description,omitempty"`
}
type KernelSpec struct {
	Versions  []string       `json:"versions,omitempty"`
	Required  []KernelConfig `json:"required,omitempty"`
	Optional  []KernelConfig `json:"optional,omitempty"`
	Forbidden []KernelConfig `json:"forbidden,omitempty"`
}
type DockerSpec struct {
	Version     []string `json:"version,omitempty"`
	GraphDriver []string `json:"graphDriver,omitempty"`
}
type RuntimeSpec struct {
	*DockerSpec `json:",inline"`
}
type PackageSpec struct {
	Name         string `json:"name,omitempty"`
	VersionRange string `json:"versionRange,omitempty"`
	Description  string `json:"description,omitempty"`
}
type PackageSpecOverride struct {
	OSDistro     string        `json:"osDistro,omitempty"`
	Subtractions []PackageSpec `json:"subtractions,omitempty"`
	Additions    []PackageSpec `json:"additions,omitempty"`
}
type SysSpec struct {
	OS                   string                `json:"os,omitempty"`
	KernelSpec           KernelSpec            `json:"kernelSpec,omitempty"`
	Cgroups              []string              `json:"cgroups,omitempty"`
	RuntimeSpec          RuntimeSpec           `json:"runtimeSpec,omitempty"`
	PackageSpecs         []PackageSpec         `json:"packageSpecs,omitempty"`
	PackageSpecOverrides []PackageSpecOverride `json:"packageSpecOverrides,omitempty"`
}
