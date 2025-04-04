package events

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"go.autokitteh.dev/autokitteh/cmd/ak/common"
	"go.autokitteh.dev/autokitteh/internal/kittehs"
	"go.autokitteh.dev/autokitteh/internal/resolver"
	"go.autokitteh.dev/autokitteh/sdk/sdktypes"
)

var dispatchCmd = common.StandardCommand(&cobra.Command{
	Use:     "dispatch [--from-file=...] [override flags]",
	Short:   "Notify server's dispatcher about new event",
	Aliases: []string{"dis", "dsp", "d"},
	Args:    cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) error {
		var event sdktypes.Event
		pb := &sdktypes.EventPB{}

		r := resolver.Resolver{Client: common.Client()}
		ctx, cancel := common.LimitedContext()
		defer cancel()

		if filename != "" {
			text, err := os.ReadFile(filename)
			if err != nil {
				return fmt.Errorf("read file: %w", err)
			}
			if err := json.Unmarshal(text, &event); err != nil {
				return fmt.Errorf("unmarshal JSON in %q: %w", filename, err)
			}
			pb = event.ToProto()
		}

		if connection != "" {
			_, cid, err := r.ConnectionNameOrID(ctx, args[0], "", sdktypes.InvalidOrgID)
			if err != nil {
				return err
			}

			pb.DestinationId = cid.String()
		}
		if eventType != "" {
			pb.EventType = eventType
		}
		if len(data) > 0 {
			m, err := kittehs.ListToMapError(data, parseDataKeyValue)
			if err != nil {
				return err
			}
			pb.Data = m
		}
		if len(memos) > 0 {
			memoMap, err := kittehs.ListToMapError(memos, parseMemoKeyValue)
			if err != nil {
				return err
			}
			pb.Memo = memoMap
		}

		e, err := sdktypes.EventFromProto(pb)
		if err != nil {
			return fmt.Errorf("invalid event: %w", err)
		}

		eid, err := common.Client().Dispatcher().Dispatch(ctx, e, nil)
		if err != nil {
			return fmt.Errorf("dispatch event: %w", err)
		}

		common.RenderKVIfV("event_id", eid)
		return nil
	},
})

func init() {
	// Command-specific flags.
	dispatchCmd.Flags().StringVarP(&filename, "from_file", "f", "", "load event data from file")
	kittehs.Must0(dispatchCmd.MarkFlagFilename("from_file"))

	dispatchCmd.Flags().StringVarP(&connection, "connection", "c", "", "connection name or ID")
	dispatchCmd.Flags().StringVarP(&eventType, "event-type", "e", "", "event type")
	dispatchCmd.Flags().StringSliceVarP(&data, "data", "d", nil, `zero or more "key=value" pairs`)
	dispatchCmd.Flags().StringSliceVarP(&memos, "memo", "m", nil, `zero or more "key=value" pairs`)
}
