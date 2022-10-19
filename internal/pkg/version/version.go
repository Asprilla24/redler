package version

var (
	serviceName = "unknown"
	buildID     = "unknown"
	buildDate   = "unknown"
	gitHash     = "unknown"
)

// Short returns a short version
func Short() map[string]interface{} {
	return map[string]interface{}{
		"name":      serviceName,
		"buildDate": buildDate,
		"version":   buildID,
		"gitHash":   gitHash,
	}
}

// GetVersion returns version
func GetVersion() string {
	return buildID
}
