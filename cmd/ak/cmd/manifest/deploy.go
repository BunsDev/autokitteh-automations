package manifest

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"

	"go.autokitteh.dev/autokitteh/cmd/ak/common"
	"go.autokitteh.dev/autokitteh/internal/backend/webhookssvc"
	"go.autokitteh.dev/autokitteh/internal/kittehs"
	"go.autokitteh.dev/autokitteh/internal/manifest"
	"go.autokitteh.dev/autokitteh/internal/resolver"
	"go.autokitteh.dev/autokitteh/sdk/sdktypes"
)

var (
	env string

	filePaths, dirPaths []string
)

// Allow no dir/file - use manifest dir by default

var deployCmd = common.StandardCommand(&cobra.Command{
	Use:   "deploy <manifest file> [--project-name <name>] [--org org] [--dir <path> [...]] [--file <path> [...]] [--env <name or ID>] [--quiet]",
	Short: "Create, configure, build, deploy, and activate project",
	Long:  `Create, configure, build, deploy, and activate project - see also the "build", "deployment", and "project" parent commands`,
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := common.LimitedContext()
		defer cancel()

		r := resolver.Resolver{Client: common.Client()}
		oid, err := r.Org(ctx, org)
		if err != nil {
			return err
		}

		// Step 1: apply the manifest file (see also the "manifest" parent command).
		effects, project, err := applyManifest(ctx, cmd, args, oid)
		if err != nil {
			return err
		}

		PrintSuggestions(cmd, effects)

		// Step 2: build the project
		// (see also the "build" and "project" parent commands).
		if len(dirPaths) == 0 && len(filePaths) == 0 {
			dirPaths = append(dirPaths, filepath.Dir(args[0]))
		}

		pid, err := r.ProjectNameOrID(ctx, oid, project)
		if err != nil {
			return err
		}

		bid, err := common.BuildProject(pid, dirPaths, filePaths)
		if err != nil {
			return err
		}
		logFunc(cmd, "exec")(fmt.Sprintf("create_build: created %q", bid))

		// Step 3: deploy the build
		// (see also the "deployment" and "project" parent commands).
		deployment, err := sdktypes.DeploymentFromProto(&sdktypes.DeploymentPB{
			ProjectId: pid.String(),
			BuildId:   bid.String(),
		})
		if err != nil {
			return fmt.Errorf("invalid deployment: %w", err)
		}

		dep := common.Client().Deployments()
		did, err := dep.Create(ctx, deployment)
		if err != nil {
			return fmt.Errorf("create deployment: %w", err)
		}
		logFunc(cmd, "exec")(fmt.Sprintf("create_deployment: created %q", did))

		// Step 5: activate the deployment
		// (see also the "deployment" parent command).
		if err := dep.Activate(ctx, did); err != nil {
			return fmt.Errorf("activate deployment: %w", err)
		}
		logFunc(cmd, "exec")("activate_deployment: activated")

		return nil
	},
})

func init() {
	// Command-specific flags.
	deployCmd.Flags().StringArrayVarP(&dirPaths, "dir", "d", []string{}, "0 or more directory paths (default = manifest directory)")
	deployCmd.Flags().StringArrayVarP(&filePaths, "file", "f", []string{}, "0 or more file paths")
	kittehs.Must0(deployCmd.MarkFlagDirname("dir"))
	kittehs.Must0(deployCmd.MarkFlagFilename("file"))
	deployCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "only show errors, if any")
	deployCmd.Flags().StringVarP(&env, "env", "e", "", "environment name or ID")
	deployCmd.Flags().StringVarP(&projectName, "project-name", "n", "", "project name")
	deployCmd.Flags().StringVarP(&org, "org", "o", "", "org name or id")
}

func applyManifest(ctx context.Context, cmd *cobra.Command, args []string, oid sdktypes.OrgID) (manifest.Effects, string, error) {
	client := common.Client()

	// Read and parse the manifest file.
	data, err := common.Consume(args)
	if err != nil {
		return nil, "", err
	}

	actions, err := plan(cmd, data, projectName, oid)
	if err != nil {
		return nil, "", err
	}

	// Execute the plan.
	effects, err := manifest.Execute(ctx, actions, client, logFunc(cmd, "exec"))
	if err != nil {
		return nil, "", err
	}

	pids := effects.ProjectIDs()

	if len(pids) == 0 {
		// Execute didn't return a new project ID because the project already exists,
		// so get the project name from the manifest instead. It's safe to ignore the
		// error here because we already ran Read() successfully inside plan() above.
		m := kittehs.Must1(manifest.Read(data))
		return effects, m.Project.Name, nil
	}
	if len(pids) > 1 {
		return nil, "", fmt.Errorf("expected 1 project ID, got %d", len(pids))
	}

	return effects, pids[0].String(), nil // Exactly one new project created.
}

func PrintSuggestions(cmd *cobra.Command, effects manifest.Effects) {
	for _, effect := range effects {
		if effect.Type != manifest.Created {
			continue
		}

		if tid, ok := effect.SubjectID.(sdktypes.TriggerID); ok {
			ctx, cancel := common.LimitedContext()
			defer cancel()

			t, err := common.Client().Triggers().Get(ctx, tid)
			if err != nil {
				fmt.Fprintf(cmd.ErrOrStderr(), "error getting trigger %v: %v\n", tid, err)
				continue
			}

			if slug := t.WebhookSlug(); slug != "" {
				urlPath := webhookssvc.WebhooksPathPrefix + slug
				fmt.Fprintf(cmd.ErrOrStderr(), "[!!!!] trigger %q created, webhook path is %q\n", t.Name(), urlPath)
			}
		}

		if cid, ok := effect.SubjectID.(sdktypes.ConnectionID); ok {
			ctx, cancel := common.LimitedContext()
			defer cancel()

			c, err := common.Client().Connections().Get(ctx, cid)
			if err != nil {
				fmt.Fprintf(cmd.ErrOrStderr(), "error getting connection %v: %v\n", cid, err)
				continue
			}

			if l := c.Links().InitURL(); l != "" {
				action := "and can be initialized"
				if c.Capabilities().RequiresConnectionInit() {
					action = "but requires initialization"
				}
				fmt.Fprintf(cmd.ErrOrStderr(), "[!!!!] Connection created, %s. Please run this to complete: ak connection init %v\n", action, cid)
			}
		}
	}
}
