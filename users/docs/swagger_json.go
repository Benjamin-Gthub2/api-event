/*
 * File: swagger_json.go
 * Author: Benjamin
 * Copyright: 2025, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file defines the component for the swagger_json.go.
 *
 * Last Modified: 2025-04-11
 */

package docs

import (
	_ "embed"
)

//go:embed swagger3.json
var DocTemplateJson string
