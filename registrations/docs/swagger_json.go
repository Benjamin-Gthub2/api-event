/*
 * File: swagger_json.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the swagger json template.
 *
 * Last Modified: 2026-04-21
 */

package docs

import (
	_ "embed"
)

//go:embed swagger3.json
var DocTemplateJson string
