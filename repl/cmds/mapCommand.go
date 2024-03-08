package cmds

func MapCommand(config *Config, args ...string) error {
	if config.NextUrl == "" {
		config.NextUrl = "location"
	}
	return nil
}
