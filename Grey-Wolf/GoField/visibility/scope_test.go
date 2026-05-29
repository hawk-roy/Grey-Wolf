package visibility

import "testing"

func TestPackageCanAccessPrivateConstant(t *testing.T) {
	if privateConstant == "" {
		t.Fatal("privateConstant should be accessible inside package visibility")
	}
}

func TestGlobalVarIsPackageLevelVariable(t *testing.T) {
	original := globalVar
	defer func() {
		globalVar = original
	}()

	globalVar = "changed in package visibility"
	if globalVar != "changed in package visibility" {
		t.Fatal("globalVar should be mutable from package scope")
	}
}
