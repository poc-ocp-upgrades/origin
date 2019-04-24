package reference

import "regexp"

var (
	alphaNumericRegexp	= match(`[a-z0-9]+`)
	separatorRegexp		= match(`(?:[._]|__|[-]*)`)
	nameComponentRegexp	= expression(alphaNumericRegexp, optional(repeated(separatorRegexp, alphaNumericRegexp)))
	hostnameComponentRegexp	= match(`(?:[a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9])`)
	hostnameRegexp		= expression(hostnameComponentRegexp, optional(repeated(literal(`.`), hostnameComponentRegexp)), optional(literal(`:`), match(`[0-9]+`)))
	TagRegexp		= match(`[\w][\w.-]{0,127}`)
	anchoredTagRegexp	= anchored(TagRegexp)
	DigestRegexp		= match(`[A-Za-z][A-Za-z0-9]*(?:[-_+.][A-Za-z][A-Za-z0-9]*)*[:][[:xdigit:]]{32,}`)
	anchoredDigestRegexp	= anchored(DigestRegexp)
	NameRegexp		= expression(optional(hostnameRegexp, literal(`/`)), nameComponentRegexp, optional(repeated(literal(`/`), nameComponentRegexp)))
	anchoredNameRegexp	= anchored(optional(capture(hostnameRegexp), literal(`/`)), capture(nameComponentRegexp, optional(repeated(literal(`/`), nameComponentRegexp))))
	ReferenceRegexp		= anchored(capture(NameRegexp), optional(literal(":"), capture(TagRegexp)), optional(literal("@"), capture(DigestRegexp)))
)
var match = regexp.MustCompile

func literal(s string) *regexp.Regexp {
	_logClusterCodePath()
	defer _logClusterCodePath()
	re := match(regexp.QuoteMeta(s))
	if _, complete := re.LiteralPrefix(); !complete {
		panic("must be a literal")
	}
	return re
}
func expression(res ...*regexp.Regexp) *regexp.Regexp {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var s string
	for _, re := range res {
		s += re.String()
	}
	return match(s)
}
func optional(res ...*regexp.Regexp) *regexp.Regexp {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return match(group(expression(res...)).String() + `?`)
}
func repeated(res ...*regexp.Regexp) *regexp.Regexp {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return match(group(expression(res...)).String() + `+`)
}
func group(res ...*regexp.Regexp) *regexp.Regexp {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return match(`(?:` + expression(res...).String() + `)`)
}
func capture(res ...*regexp.Regexp) *regexp.Regexp {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return match(`(` + expression(res...).String() + `)`)
}
func anchored(res ...*regexp.Regexp) *regexp.Regexp {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return match(`^` + expression(res...).String() + `$`)
}
