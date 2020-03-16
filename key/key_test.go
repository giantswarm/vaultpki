package key

import (
	"strconv"
	"testing"
)

func Test_IsMountPath(t *testing.T) {
	testCases := []struct {
		name    string
		path    string
		matches bool
	}{
		{
			name:    "case 0 matches a usual Tenant Cluster PKI backend",
			path:    "pki-al9qy/",
			matches: true,
		},
		{
			name:    "case 1 matches a usual Tenant Cluster PKI backend",
			path:    "pki-mn4k1/",
			matches: true,
		},
		{
			name:    "case 2 ensures to not match non-cluster-id patterns",
			path:    "pki-92384/",
			matches: false,
		},
		{
			name:    "case 3 ensures to not match non-cluster-id patterns",
			path:    "pki-hdmed/",
			matches: false,
		},
		{
			name:    "case 4 ensures to not match the Control Plane PKI backend as this should stay untouched",
			path:    "pki-g8s/",
			matches: false,
		},
		{
			name:    "case 5 ensures to not match other paths",
			path:    "g8s/",
			matches: false,
		},
		{
			name:    "case 6 ensures to not match other Vault specific paths",
			path:    "secrets/",
			matches: false,
		},
		{
			name:    "case 7 ensures to not match other Vault specific paths",
			path:    "sys/",
			matches: false,
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Log(tc.name)

			matches := IsMountPath(tc.path)

			if tc.matches && !matches {
				t.Fatalf("expected path %#q to match", tc.path)
			}
			if !tc.matches && matches {
				t.Fatalf("expected path %#q not to match", tc.path)
			}
		})
	}
}
