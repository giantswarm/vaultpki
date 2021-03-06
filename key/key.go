package key

import (
	"fmt"
	"regexp"
	"strings"
)

func CommonName(ID string, commonNameFormat string) string {
	return fmt.Sprintf(commonNameFormat, ID)
}

// ClusterIDFromMountPath aligns with the implementation details of IsMountPath
// and expects a valid mount path to work properly. A valid mount path must look
// something like the following, where "z4574" is the cluster ID.
//
//     pki-z4574/
//
func ClusterIDFromMountPath(path string) string {
	return path[4:9]
}

func ListMountsPath(ID string) string {
	return fmt.Sprintf("pki-%s/", ID)
}

// IsMountPath verifies if path is the expected mount path we use for our PKI
// backends. One special requirement is the format of our Tenant Cluster IDs.
// Further we must not mess around with the Control Plane specific G8s PKI
// backend. Thus IsMountPath does not consider "pki-g8s" to be a mount path.
func IsMountPath(path string) bool {
	if !strings.HasPrefix(path, "pki-") {
		return false
	}
	if !strings.HasSuffix(path, "/") {
		return false
	}
	if len(path) != 10 {
		return false
	}

	id := path[4:9]
	hasLetters := regexp.MustCompile(`[a-z]+`).MatchString(id)
	hasNumbers := regexp.MustCompile(`[0-9]+`).MatchString(id)

	return hasLetters && hasNumbers
}

func MountPKIPath(ID string) string {
	return fmt.Sprintf("pki-%s", ID)
}

func ReadCAPath(ID string) string {
	return fmt.Sprintf("pki-%s/cert/ca", ID)
}

func WriteCAPath(ID string, exportPrivateKey bool) string {
	generationType := "internal"
	if exportPrivateKey {
		generationType = "exported"
	}

	return fmt.Sprintf("pki-%s/root/generate/%s", ID, generationType)
}
