// the "embed" tag decides whether to include this file or no_embed.go
// go build -tags embed
//go:build embed

package main

import (
	"embed"
)

//go:embed public
var embeddedFiles embed.FS

const hasEmbedded = true
