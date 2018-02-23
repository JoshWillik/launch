package config

// Settings contains the AWS credentials and bucket name
type Settings struct {
}

// ReadOrCreateSettings reads the existing project settings or prompts the user
// to create new ones
func ReadOrCreateSettings() Settings {
	return Settings{}
}
