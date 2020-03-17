package vaultpkitest

import (
	"testing"

	"github.com/giantswarm/vaultpki"
)

func Test_VaultPKITest_New_Interface(t *testing.T) {
	// Create an anonymus function that takes the interface as argument and
	// provide the test implementation in order to verify it implements the actual
	// interface.
	func(v vaultpki.Interface) {}(New())
}
