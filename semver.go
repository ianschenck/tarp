package tarp

import (
	"regexp"
)

const semverExpression = `^v?((\d+)\.(\d+)\.(\d+))(?:-([\dA-Za-z\-]+(?:\.[\dA-Za-z\-]+)*))?(?:\+([\dA-Za-z\-]+(?:\.[\dA-Za-z\-]+)*))?$`

var SemverRE = regexp.MustCompile(semverExpression)
