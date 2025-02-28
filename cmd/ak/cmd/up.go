package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"go.autokitteh.dev/autokitteh/cmd/ak/common"
)

var upCmd = common.StandardCommand(&cobra.Command{
	Use:   "up [--mode {default|dev|test}]",
	Short: "Start local server",
	Args:  cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Root().Context()

		app, err := common.NewSvc()
		if err != nil {
			return fmt.Errorf("new service: %w", err)
		}

		if err := app.Start(ctx); err != nil {
			return fmt.Errorf("fx app start: %w", err)
		}

		<-app.Wait()

		fmt.Println() // End the output with "\n".
		return nil
	},
})

func init() {
	// Command-specific flags.
	common.AddModeFlag(upCmd)
	common.AddSilentFlag(upCmd)
}
