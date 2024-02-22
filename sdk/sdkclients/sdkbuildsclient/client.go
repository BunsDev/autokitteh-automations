package sdkbuildsclient

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"connectrpc.com/connect"

	"go.autokitteh.dev/autokitteh/internal/kittehs"
	buildsv1 "go.autokitteh.dev/autokitteh/proto/gen/go/autokitteh/builds/v1"
	"go.autokitteh.dev/autokitteh/proto/gen/go/autokitteh/builds/v1/buildsv1connect"
	"go.autokitteh.dev/autokitteh/sdk/internal/rpcerrors"
	"go.autokitteh.dev/autokitteh/sdk/sdkclients/internal"
	"go.autokitteh.dev/autokitteh/sdk/sdkclients/sdkclient"
	"go.autokitteh.dev/autokitteh/sdk/sdkservices"
	"go.autokitteh.dev/autokitteh/sdk/sdktypes"
)

type client struct {
	client buildsv1connect.BuildsServiceClient
}

// Download implements sdkservices.Builds.
func (c *client) Download(ctx context.Context, buildID sdktypes.BuildID) (io.ReadCloser, error) {
	resp, err := c.client.Download(ctx, connect.NewRequest(&buildsv1.DownloadRequest{BuildId: buildID.String()}))
	if err != nil {
		return nil, rpcerrors.TranslateError(err)
	}

	if err := internal.Validate(resp.Msg); err != nil {
		return nil, err
	}

	reader := io.NopCloser(bytes.NewReader(resp.Msg.Data))
	return reader, nil
}

// Get implements sdkservices.Builds.
func (c *client) Get(ctx context.Context, buildID sdktypes.BuildID) (sdktypes.Build, error) {
	resp, err := c.client.Get(ctx, connect.NewRequest(&buildsv1.GetRequest{BuildId: buildID.String()}))
	if err != nil {
		return nil, rpcerrors.TranslateError(err)
	}

	if err := internal.Validate(resp.Msg); err != nil {
		return nil, err
	}
	if resp.Msg.Build == nil {
		return nil, nil
	}

	build, err := sdktypes.StrictBuildFromProto(resp.Msg.Build)
	if err != nil {
		return nil, fmt.Errorf("invalid build: %w", err)
	}
	return build, nil
}

// List implements sdkservices.Builds.
func (c *client) List(ctx context.Context, filter sdkservices.ListBuildsFilter) ([]sdktypes.Build, error) {
	resp, err := c.client.List(ctx, connect.NewRequest(&buildsv1.ListRequest{
		Limit: filter.Limit,
	}))
	if err != nil {
		return nil, rpcerrors.TranslateError(err)
	}

	if err := internal.Validate(resp.Msg); err != nil {
		return nil, err
	}

	builds, err := kittehs.TransformError(resp.Msg.Builds, sdktypes.StrictBuildFromProto)
	if err != nil {
		return nil, fmt.Errorf("invalid build: %w", err)
	}
	return builds, nil
}

// Remove implements sdkservices.Builds.
func (c *client) Remove(ctx context.Context, buildID sdktypes.BuildID) error {
	resp, err := c.client.Remove(ctx, connect.NewRequest(&buildsv1.RemoveRequest{BuildId: buildID.String()}))
	if err != nil {
		return rpcerrors.TranslateError(err)
	}

	if err := internal.Validate(resp.Msg); err != nil {
		return err
	}

	return nil
}

// Save implements sdkservices.Builds.
func (c *client) Save(ctx context.Context, build sdktypes.Build, data []byte) (sdktypes.BuildID, error) {
	resp, err := c.client.Save(ctx, connect.NewRequest(&buildsv1.SaveRequest{Build: build.ToProto(), Data: data}))
	if err != nil {
		return nil, rpcerrors.TranslateError(err)
	}

	if err := internal.Validate(resp.Msg); err != nil {
		return nil, err
	}

	buildID, err := sdktypes.StrictParseBuildID(resp.Msg.BuildId)
	if err != nil {
		return nil, fmt.Errorf("invalid build: %w", err)
	}
	return buildID, nil
}

func New(p sdkclient.Params) sdkservices.Builds {
	return &client{client: internal.New(buildsv1connect.NewBuildsServiceClient, p)}
}
