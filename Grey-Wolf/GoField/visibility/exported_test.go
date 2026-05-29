package visibility_test

import (
	"testing"

	"gofield/visibility"
)

func TestPublicConstantCanBeAccessedFromAnotherPackage(t *testing.T) {
	if visibility.PublicConstant == "" {
		t.Fatal("PublicConstant should be accessible from another package")
	}
}
