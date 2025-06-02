package helper

import (
	"github.com/labstack/gommon/log"
	"strings"
)

// IdentifierMatcher checks if keyName matches any identifier in keyIdentifier string
// Returns true if found and the matched string
func IdentifierMatcher(keyIdentifier, keyName string) (bool, string) {
	// Handle empty cases
	if keyIdentifier == "" || keyName == "" {
		return false, ""
	}

	// Split identifier by pipe and process each part
	identifiers := strings.Split(keyIdentifier, "|")

	// Normalize the wmsName (trim spaces and convert to lowercase)
	searchName := strings.ToLower(strings.TrimSpace(keyName))

	log.Infof("helper IdentifierMatcher with equal COMPARISON keyName=%s", keyName)
	for _, id := range identifiers {
		// Normalize the identifier
		cleanId := strings.TrimSpace(id)

		// Simple case-insensitive comparison
		if strings.ToLower(cleanId) == searchName {
			return true, cleanId
		}
	}

	log.Infof("helper IdentifierMatcher with equal CONTAINS keyName=%s", keyName)
	for _, id := range identifiers {
		// Normalize the identifier
		cleanId := strings.TrimSpace(id)
		cleanIdLowerCase := strings.ToLower(cleanId)

		// Simple case-insensitive comparison
		if strings.Contains(cleanIdLowerCase, searchName) {
			log.Infof("helper IdentifierMatcher with equal CONTAINS MATCHED keyName=%s", keyName)
			return true, cleanId
		}
	}

	return false, ""
}
