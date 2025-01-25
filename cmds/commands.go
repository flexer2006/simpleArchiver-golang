// Package cmds handles command initialization and registration for the CLI application.
// It provides functionality to add subcommands to the root command and ensures proper
// error handling during command setup through panic recovery mechanisms.
package cmds

import (
	"github.com/flexer2006/simpleArchiver-golang/internal/application"
	"github.com/flexer2006/simpleArchiver-golang/pkg/vlcPack"
	"github.com/flexer2006/simpleArchiver-golang/pkg/vlcUnpack"
)

// InitCommands registers subcommands with the root command and wraps the initialization
// process with panic recovery. This function:
//   - Adds vlcPack.VlcPackCmd as a subcommand for packing operations
//   - Adds vlcUnpack.VlcUnpackCmd as a subcommand for unpacking operations
//   - Uses application.HandlePanic to ensure safe command registration
//
// Should be called during application startup before executing the root command.
func InitCommands() {
	application.HandlePanic(func() {
		application.RootCmd.AddCommand(vlcPack.VlcPackCmd)
		application.RootCmd.AddCommand(vlcUnpack.VlcUnpackCmd)
	})
}
