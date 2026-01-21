//go:build embed

package main

import (
	"embed"
)

//go:embed public
var embeddedFiles embed.FS

const hasEmbedded = true
