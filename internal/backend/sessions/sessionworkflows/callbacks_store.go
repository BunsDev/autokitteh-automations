package sessionworkflows

import (
	"context"

	"go.temporal.io/sdk/workflow"

	"go.autokitteh.dev/autokitteh/sdk/sdktypes"
)

func (w *sessionWorkflow) listStoreValues(wctx workflow.Context) func(context.Context, sdktypes.RunID) ([]string, error) {
	return func(context.Context, sdktypes.RunID) ([]string, error) {
		var vs []string

		if err := workflow.ExecuteActivity(wctx, listStoreValuesActivityName, w.data.Session.ProjectID()).Get(wctx, &vs); err != nil {
			return nil, err
		}

		return vs, nil
	}
}

func (w *sessionWorkflow) getStoreValue(wctx workflow.Context) func(context.Context, sdktypes.RunID, string) (sdktypes.Value, error) {
	return func(_ context.Context, _ sdktypes.RunID, key string) (sdktypes.Value, error) {
		var vs map[string]sdktypes.Value

		if err := workflow.ExecuteActivity(wctx, getStoreValueActivityName, w.data.Session.ProjectID(), []string{key}).Get(wctx, &vs); err != nil {
			return sdktypes.InvalidValue, err
		}

		if v, ok := vs[key]; ok {
			return v, nil
		}

		return sdktypes.Nothing, nil
	}
}

func (w *sessionWorkflow) mutateStoreValue(wctx workflow.Context) func(context.Context, sdktypes.RunID, string, string, ...sdktypes.Value) (sdktypes.Value, error) {
	return func(_ context.Context, rid sdktypes.RunID, key, op string, operands ...sdktypes.Value) (sdktypes.Value, error) {
		var v sdktypes.Value

		if err := workflow.ExecuteActivity(wctx, mutateStoreValueActivityName, w.data.Session.ProjectID(), key, op, operands).Get(wctx, &v); err != nil {
			return sdktypes.InvalidValue, err
		}

		return v, nil
	}
}
