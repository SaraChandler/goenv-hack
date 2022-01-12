package pkg

import "github.com/coreos/go-semver/semver"

// Validate a given version string
func ValidateVersion(version string) error {
	_, err := semver.NewVersion(version)
	return err
}

// Compare version equality
// -1 = A < B
// 0 = A == B
// 1 = A > B
// 1.16 == 1.16.8 vA.Minor == vB.Minor??
// 1.16.8 != 1.16.6
func CompareVersions(versionA string, versionB string) (int, error) {
	vA := semver.New(versionA)
	vB := semver.New(versionB)

	return vA.Compare(*vB), nil
}
