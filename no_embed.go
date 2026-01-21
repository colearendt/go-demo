// the "embed" tag decides whether to include this file or embed.go
// go build
// go build -tags embed=false
//go:build !embed

package main

import (
	"embed"
)

var embeddedFiles embed.FS

const hasEmbedded = false
