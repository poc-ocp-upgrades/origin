package system

type KernelValidatorHelper interface{ GetKernelReleaseVersion() (string, error) }
