package docs

import (
	_ "embed"
)

//go:embed swagger3.json
var DocTemplateJson string
