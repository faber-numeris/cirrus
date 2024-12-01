package openapi

import (
	"embed"
	_ "embed"
)

//go:embed *
var Files embed.FS
