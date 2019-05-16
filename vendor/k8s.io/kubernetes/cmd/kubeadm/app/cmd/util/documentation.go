package util

import (
	"k8s.io/kubernetes/pkg/util/normalizer"
)

var (
	AlphaDisclaimer = `
		Alpha Disclaimer: this command is currently alpha.
	`
	MacroCommandLongDescription = normalizer.LongDesc(`
		This command is not meant to be run on its own. See list of available subcommands.
	`)
)
