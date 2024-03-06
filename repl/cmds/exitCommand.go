package cmds

import "os"

func ExitCommand() error {
	defer os.Exit(0)
	return nil
}
