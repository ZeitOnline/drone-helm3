package run

import (
	"os"
)

// Help is a step in a helm Plan that calls `helm help`.
type Help struct {
	cmd   cmd
	debug bool
}

// Run launches the command.
func (h *Help) Run() error {
	return h.cmd.Run()
}

// NewHelp returns a new Help.
func NewHelp(debug bool) *Help {
	h := Help{
		debug: debug,
	}

	h.cmd = command(helmBin, "help")
	h.cmd.Stdout(os.Stdout)
	h.cmd.Stderr(os.Stderr)

	return &h
}
