package cmds

import "os"

func ExitCommand(config *Config, args ...string) error {
	defer os.Exit(0)
	return nil
}
