//go:build !embed

package main

import (
	"embed"
)

var embeddedFiles embed.FS

const hasEmbedded = false
