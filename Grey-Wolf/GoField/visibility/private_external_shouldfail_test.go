//go:build shouldfail
// +build shouldfail

package visibility_test

import "gofield/visibility"

var _ = visibility.privateConstant
