package projects

import (
	"github.com/spf13/cobra"

	"go.autokitteh.dev/autokitteh/cmd/ak/common"
	"go.autokitteh.dev/autokitteh/internal/resolver"
	"go.autokitteh.dev/autokitteh/sdk/sdktypes"
)

var getCmd = common.StandardCommand(&cobra.Command{
	Use:     "get <project name or ID> [--fail]",
	Short:   "Get project details",
	Aliases: []string{"g"},
	Args:    cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		r := resolver.Resolver{Client: common.Client()}
		ctx, cancel := common.LimitedContext()
		defer cancel()

		pid, err := r.ProjectNameOrID(ctx, args[0])
		err = common.AddNotFoundErrIfCond(err, pid.IsValid())
		if err = common.ToExitCodeWithSkipNotFoundFlag(cmd, err, "project"); err == nil {
			var p sdktypes.Project
			p, err = projects().GetByID(ctx, pid)
			err = common.AddNotFoundErrIfCond(err, p.IsValid())
			if err = common.ToExitCodeWithSkipNotFoundFlag(cmd, err, "project"); err == nil {
				common.RenderKVIfV("project", p)
			}
		}
		return err
	},
})

func init() {
	// Command-specific flags.
	common.AddFailIfNotFoundFlag(getCmd)
}
