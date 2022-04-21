//go:build tools

package tools

import (
	_ "golang.org/x/tools/cmd/stringer"
	_ "mvdan.cc/gofumpt" // code formatting
)
