package pkg

import (
	"github.com/coreos/go-semver/semver"
)

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

//fuzzyCompare("1.16.5", "1.16") == true, fuzzyCompare("1.15.9", "1.16") == false
func fuzzyCompare(versionA string, versionB string) bool {
	shortVersion := ""
	if len(versionA) > len(versionB) {
		shortVersion = versionB
	} else {
		shortVersion = versionA
	}

	for i := 0; i < len(shortVersion); i++ {
		if versionA[i] != versionB[i] {
			return false
		}
		if versionA[i] == versionB[i] && i == 3 {
			return true
		}
	}
	return false
}
