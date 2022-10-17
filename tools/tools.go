//go:build tools
// +build tools

package tools

// from https://github.com/golang/go/issues/25922#issuecomment-412992431

import (
	_ "github.com/cloudspannerecosystem/wrench"
	_ "github.com/gcpug/zagane"
	_ "go.mercari.io/yo"
)
