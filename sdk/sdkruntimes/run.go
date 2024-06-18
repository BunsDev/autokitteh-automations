package sdkruntimes

import (
	"context"
	"errors"
	"fmt"

	"go.autokitteh.dev/autokitteh/sdk/sdkbuildfile"
	"go.autokitteh.dev/autokitteh/sdk/sdkerrors"
	"go.autokitteh.dev/autokitteh/sdk/sdkservices"
	"go.autokitteh.dev/autokitteh/sdk/sdktypes"
)

type RunParams struct {
	Runtimes             sdkservices.Runtimes
	BuildFile            *sdkbuildfile.BuildFile
	Globals              map[string]sdktypes.Value
	RunID                sdktypes.RunID
	FallthroughCallbacks sdkservices.RunCallbacks
	EntryPointPath       string
}

// Run executes a build file and manages it across multiple runtimes.
// fallthourghCallbacks.{Load,Call} are called only for dynamic modules
// (modules that are supplied from integrations).
func Run(ctx context.Context, params RunParams) (sdkservices.Run, error) {
	group := &group{
		mainID: params.RunID,
		runs:   make(map[sdktypes.ExecutorID]sdkservices.Run),
	}

	cbs := sdkservices.RunCallbacks{
		Print: params.FallthroughCallbacks.SafePrint,
		Call: func(ctx context.Context, runID sdktypes.RunID, v sdktypes.Value, args []sdktypes.Value, kwargs map[string]sdktypes.Value) (sdktypes.Value, error) {
			fv := v.GetFunction()

			if fv.IsConst() {
				return fv.ConstValue()
			}

			xid := fv.ExecutorID()
			if !xid.IsValid() {
				return sdktypes.InvalidValue, fmt.Errorf("invalid executor id: %w", sdkerrors.ErrNotFound)
			}

			if !fv.HasFlag(sdktypes.PureFunctionFlag) || group.runs[xid] == nil {
				return params.FallthroughCallbacks.SafeCall(ctx, runID, v, args, kwargs)
			}

			return group.Call(ctx, v, args, kwargs)
		},
		NewRunID:   params.FallthroughCallbacks.NewRunID,
		DebugTrace: params.FallthroughCallbacks.SafeDebugTrace,
	}

	cache := make(map[string]map[string]sdktypes.Value)

	cbs.Load = func(ctx context.Context, rid sdktypes.RunID, path string) (map[string]sdktypes.Value, error) {
		exports, ok := cache[path]
		if ok {
			if exports == nil {
				return nil, fmt.Errorf("detected circular dependency detected involving %q", path)
			}

			return exports, nil
		}

		loadRunID := params.FallthroughCallbacks.SafeNewRunID()

		runParams := params
		runParams.Globals = params.Globals
		runParams.RunID = loadRunID
		runParams.FallthroughCallbacks = cbs

		if vs, err := params.FallthroughCallbacks.SafeLoad(ctx, rid, path); err == nil {
			cache[path] = vs
			return vs, nil
		} else if !errors.Is(err, sdkerrors.ErrNotFound) {
			return nil, err
		}

		cache[path] = nil

		r, err := run(ctx, runParams, path)
		if err != nil {
			if errors.Is(err, sdkerrors.ErrNotFound) {
				return params.FallthroughCallbacks.SafeLoad(ctx, rid, path)
			}

			return nil, err
		}

		group.runs[sdktypes.NewExecutorID(loadRunID)] = r

		cache[path] = r.Values()

		return r.Values(), err
	}

	runParams := params
	runParams.FallthroughCallbacks = cbs

	r, err := run(ctx, runParams, params.EntryPointPath)
	if r != nil {
		r.Close()
	}

	group.runs[sdktypes.NewExecutorID(group.mainID)] = r

	return group, err
}

func run(ctx context.Context, params RunParams, path string) (sdkservices.Run, error) {
	ls, err := params.Runtimes.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("list runtimes: %w", err)
	}

	rtd, ok := MatchRuntimeByPath(ls, path)
	if !ok {
		return nil, sdkerrors.ErrNotFound
	}

	found := -1

	for i, brt := range params.BuildFile.Runtimes {
		if brt.Info.Name == rtd.Name() {
			found = i
			break
		}
	}

	if found < 0 {
		return nil, fmt.Errorf("no matching runtime for path %q", path)
	}

	brt := params.BuildFile.Runtimes[found]

	rt, err := params.Runtimes.New(ctx, brt.Info.Name)
	if err != nil {
		return nil, fmt.Errorf("new runtime: %w", err)
	}

	return rt.Run(
		ctx,
		params.RunID,
		path,
		brt.Artifact.CompiledData(),
		params.Globals,
		&params.FallthroughCallbacks,
	)
}
