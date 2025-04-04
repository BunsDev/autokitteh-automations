// Package resolver contains functions that resolve names and ID
// strings of autokitteh entities to their concrete SDK types.
package resolver

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"go.autokitteh.dev/autokitteh/sdk/sdkerrors"
	"go.autokitteh.dev/autokitteh/sdk/sdkservices"
	"go.autokitteh.dev/autokitteh/sdk/sdktypes"
)

const (
	separator = "/"
)

type Resolver struct {
	Client sdkservices.DBServices
}

// FIXME: move to sdkerrors
type NotFoundError struct {
	Type, Name string
}

var ErrNotFound = new(NotFoundError)

func (e NotFoundError) Error() string {
	name := e.Name
	if name != "" {
		name = fmt.Sprintf(" %q", name)
	}
	return e.Type + name + " not found"
}

func translateError[O sdktypes.Object](err error, obj O, typ, idOrName string) error {
	// hack - rework
	what := "ID"
	if len(idOrName) > 4 && idOrName[4] != '_' {
		what = "name"
	}

	if err != nil {
		if !errors.Is(err, sdkerrors.ErrNotFound) {
			return fmt.Errorf("get %s %s %q: %w", typ, what, idOrName, err)
		}
		return err // not found
	}
	// no error. But most of the services.Get() methods filtering out notFound errors.
	// check sdktype.IsValid() to cover this case
	if !obj.IsValid() {
		return sdkerrors.ErrNotFound
	}
	return nil
}

// BuildID returns a build, based on the given ID.
// It does NOT accept empty input.
func (r Resolver) BuildID(ctx context.Context, id string) (b sdktypes.Build, bid sdktypes.BuildID, err error) {
	if id == "" {
		err = errors.New("missing build ID")
		return
	}

	if bid, err = sdktypes.StrictParseBuildID(id); err != nil {
		err = fmt.Errorf("invalid build ID %q: %w", id, err)
		return
	}

	b, err = r.Client.Builds().Get(ctx, bid)
	err = translateError(err, b, "build", id)
	return
}

// DeploymentID returns a deployment, based on the given ID.
// It does NOT accept empty input.
func (r Resolver) DeploymentID(ctx context.Context, id string) (d sdktypes.Deployment, did sdktypes.DeploymentID, err error) {
	if id == "" {
		err = errors.New("missing deployment ID")
		return
	}

	if did, err = sdktypes.Strict(sdktypes.ParseDeploymentID(id)); err != nil {
		err = fmt.Errorf("invalid deployment ID %q: %w", id, err)
		return
	}

	d, err = r.Client.Deployments().Get(ctx, did)
	err = translateError(err, d, "deployment", id)
	return
}

// ConnectionNameOrID returns a connection, based on the given name or
// ID. If the input is empty, we return nil but not an error.
func (r Resolver) ConnectionNameOrID(ctx context.Context, nameOrID, project string, oid sdktypes.OrgID) (c sdktypes.Connection, cid sdktypes.ConnectionID, err error) {
	if nameOrID == "" {
		return
	}

	if sdktypes.IsConnectionID(nameOrID) {
		return r.connectionByID(ctx, nameOrID)
	}

	parts := strings.Split(nameOrID, separator)
	switch len(parts) {
	case 1:
		if project == "" {
			err = fmt.Errorf("invalid connection name %q: missing project prefix", nameOrID)
		} else {
			return r.connectionByFullName(ctx, oid, project, parts[0], nameOrID)
		}
		return
	case 2:
		return r.connectionByFullName(ctx, oid, parts[0], parts[1], nameOrID)
	default:
		err = fmt.Errorf("invalid connection name %q: too many parts", nameOrID)
		return
	}
}

// TriggerNameOrID returns a trigger, based on the given name or
// ID. If the input is empty, we return nil but not an error.
func (r Resolver) TriggerNameOrID(ctx context.Context, oid sdktypes.OrgID, nameOrID, project string) (c sdktypes.Trigger, cid sdktypes.TriggerID, err error) {
	if nameOrID == "" {
		return
	}

	if sdktypes.IsTriggerID(nameOrID) {
		return r.TriggerID(ctx, nameOrID)
	}

	parts := strings.Split(nameOrID, separator)
	switch len(parts) {
	case 1:
		if project == "" {
			err = fmt.Errorf("invalid trigger name %q: missing project prefix", nameOrID)
		} else {
			return r.triggerByFullName(ctx, oid, project, parts[0], nameOrID)
		}
		return
	case 2:
		return r.triggerByFullName(ctx, oid, parts[0], parts[1], nameOrID)
	default:
		err = fmt.Errorf("invalid trigger name %q: too many parts", nameOrID)
		return
	}
}

func (r Resolver) connectionByID(ctx context.Context, id string) (c sdktypes.Connection, cid sdktypes.ConnectionID, err error) {
	if cid, err = sdktypes.StrictParseConnectionID(id); err != nil {
		err = fmt.Errorf("invalid connection ID %q: %w", id, err)
		return
	}

	if c, err = r.Client.Connections().Get(ctx, cid); err != nil {
		err = fmt.Errorf("get connection ID %q: %w", id, err)
		return
	}

	return
}

// TODO: add type and maybe id to sdkerrors.ErrNotFound and replace NotFoundError below
func (r Resolver) connectionByFullName(ctx context.Context, oid sdktypes.OrgID, projNameOrID, connName, fullName string) (sdktypes.Connection, sdktypes.ConnectionID, error) {
	pid, err := r.ProjectNameOrID(ctx, oid, projNameOrID)
	if err != nil {
		return sdktypes.InvalidConnection, sdktypes.InvalidConnectionID, err
	}
	if !pid.IsValid() {
		return sdktypes.InvalidConnection, sdktypes.InvalidConnectionID, NotFoundError{Type: "project", Name: projNameOrID}
	}

	f := sdkservices.ListConnectionsFilter{ProjectID: pid}
	cs, err := r.Client.Connections().List(ctx, f)
	if err != nil {
		return sdktypes.InvalidConnection, sdktypes.InvalidConnectionID, fmt.Errorf("list connections: %w", err)
	}

	for _, c := range cs {
		if c.Name().String() == connName {
			return c, c.ID(), nil
		}
	}

	return sdktypes.InvalidConnection, sdktypes.InvalidConnectionID, NotFoundError{Type: "connection", Name: fullName}
}

// TODO: add type and maybe id to sdkerrors.ErrNotFound and replace NotFoundError below
func (r Resolver) triggerByFullName(ctx context.Context, oid sdktypes.OrgID, projNameOrID, triggerName, fullName string) (sdktypes.Trigger, sdktypes.TriggerID, error) {
	pid, err := r.ProjectNameOrID(ctx, oid, projNameOrID)
	if err != nil {
		return sdktypes.InvalidTrigger, sdktypes.InvalidTriggerID, err
	}
	if !pid.IsValid() {
		return sdktypes.InvalidTrigger, sdktypes.InvalidTriggerID, NotFoundError{Type: "project", Name: projNameOrID}
	}

	f := sdkservices.ListTriggersFilter{ProjectID: pid}
	cs, err := r.Client.Triggers().List(ctx, f)
	if err != nil {
		return sdktypes.InvalidTrigger, sdktypes.InvalidTriggerID, fmt.Errorf("list triggers: %w", err)
	}

	for _, c := range cs {
		if c.Name().String() == triggerName {
			return c, c.ID(), nil
		}
	}

	return sdktypes.InvalidTrigger, sdktypes.InvalidTriggerID, NotFoundError{Type: "trigger", Name: fullName}
}

// EventID returns an event, based on the given ID.
// It does NOT accept empty input.
func (r Resolver) EventID(ctx context.Context, id string) (e sdktypes.Event, eid sdktypes.EventID, err error) {
	if id == "" {
		err = errors.New("missing event ID")
		return
	}

	if eid, err = sdktypes.Strict(sdktypes.ParseEventID(id)); err != nil {
		err = fmt.Errorf("invalid event ID %q: %w", id, err)
		return
	}

	e, err = r.Client.Events().Get(ctx, eid)
	err = translateError(err, e, "event", id)
	return
}

// IntegrationNameOrID returns an integration, based on the given
// name or ID. If the input is empty, we return nil but not an error.
func (r Resolver) IntegrationNameOrID(ctx context.Context, nameOrID string) (sdktypes.Integration, sdktypes.IntegrationID, error) {
	if nameOrID == "" {
		return sdktypes.InvalidIntegration, sdktypes.InvalidIntegrationID, nil
	}

	if sdktypes.IsIntegrationID(nameOrID) {
		return r.integrationByID(ctx, nameOrID)
	}

	return r.integrationByName(ctx, nameOrID)
}

func (r Resolver) integrationByID(ctx context.Context, id string) (sdktypes.Integration, sdktypes.IntegrationID, error) {
	iid, err := sdktypes.Strict(sdktypes.ParseIntegrationID(id))
	if err != nil {
		return sdktypes.InvalidIntegration, sdktypes.InvalidIntegrationID, fmt.Errorf("invalid integration ID %q: %w", id, err)
	}

	is, err := r.Client.Integrations().List(ctx, "")
	if err != nil {
		return sdktypes.InvalidIntegration, sdktypes.InvalidIntegrationID, fmt.Errorf("list integrations: %w", err)
	}

	for _, i := range is {
		if i.ID() == iid {
			return i, iid, nil
		}
	}

	return sdktypes.InvalidIntegration, iid, nil
}

func (r Resolver) integrationByName(ctx context.Context, name string) (sdktypes.Integration, sdktypes.IntegrationID, error) {
	is, err := r.Client.Integrations().List(ctx, name)
	if err != nil {
		return sdktypes.InvalidIntegration, sdktypes.InvalidIntegrationID, fmt.Errorf("list integrations: %w", err)
	}

	for _, i := range is {
		if i.UniqueName().String() == name {
			return i, i.ID(), nil
		}
		if i.DisplayName() == name {
			return i, i.ID(), nil
		}
	}

	return sdktypes.InvalidIntegration, sdktypes.InvalidIntegrationID, nil
}

// TriggerID returns a trigger, based on the given ID.
// If the input is empty, we return nil but not an error.
func (r Resolver) TriggerID(ctx context.Context, id string) (t sdktypes.Trigger, tid sdktypes.TriggerID, err error) {
	if id == "" {
		err = errors.New("missing trigger ID")
		return
	}

	if tid, err = sdktypes.StrictParseTriggerID(id); err != nil {
		err = fmt.Errorf("invalid trigger ID %q: %w", id, err)
		return
	}

	t, err = r.Client.Triggers().Get(ctx, tid)
	err = translateError(err, t, "trigger", id)
	return
}

// ProjectNameOrID returns a project, based on the given name or ID.
// If the input is empty, we return nil but not an error.
func (r Resolver) ProjectNameOrID(ctx context.Context, oid sdktypes.OrgID, nameOrID string) (sdktypes.ProjectID, error) {
	if nameOrID == "" {
		return sdktypes.InvalidProjectID, nil
	}

	if sdktypes.IsProjectID(nameOrID) {
		return r.projectByID(nameOrID)
	}

	return r.projectByName(ctx, oid, nameOrID)
}

func (r Resolver) projectByID(id string) (pid sdktypes.ProjectID, err error) {
	if pid, err = sdktypes.StrictParseProjectID(id); err != nil {
		err = fmt.Errorf("invalid project ID %q: %w", id, err)
		return
	}

	return
}

func (r Resolver) projectByName(ctx context.Context, oid sdktypes.OrgID, name string) (pid sdktypes.ProjectID, err error) {
	org, proj, ok := strings.Cut(name, ".")
	if !ok {
		org, proj = "", name
	}

	var n sdktypes.Symbol
	if n, err = sdktypes.Strict(sdktypes.ParseSymbol(proj)); err != nil {
		err = fmt.Errorf("invalid project name %q: %w", name, err)
		return
	}

	if org != "" {
		oid, err = r.Org(ctx, org)
		if err != nil {
			return
		}
	}

	p, err := r.Client.Projects().GetByName(ctx, oid, n)
	err = translateError(err, p, "project", name)
	pid = p.ID()
	return
}

// SessionID returns a session, based on the given ID.
// If the input is empty, we return nil but not an error.
func (r Resolver) SessionID(ctx context.Context, id string) (s sdktypes.Session, sid sdktypes.SessionID, err error) {
	if id == "" {
		err = errors.New("missing session ID")
		return
	}

	if sid, err = sdktypes.StrictParseSessionID(id); err != nil {
		err = fmt.Errorf("invalid session ID %q: %w", id, err)
		return
	}

	s, err = r.Client.Sessions().Get(ctx, sid)
	err = translateError(err, s, "session", id)
	return
}

func (r Resolver) UserID(ctx context.Context, emailOrID string) (uid sdktypes.UserID, err error) {
	if emailOrID == "" {
		return
	}

	if sdktypes.IsUserID(emailOrID) {
		return sdktypes.StrictParseUserID(emailOrID)
	}

	uid, err = r.Client.Users().GetID(ctx, emailOrID)
	return
}

// User returns a user, based on the given email address or user ID.
func (r Resolver) User(ctx context.Context, emailOrID string) (u sdktypes.User, uid sdktypes.UserID, err error) {
	if emailOrID == "" {
		return
	}

	var email string

	if sdktypes.IsUserID(emailOrID) {
		if uid, err = sdktypes.StrictParseUserID(emailOrID); err != nil {
			err = fmt.Errorf("invalid user ID %q: %w", emailOrID, err)
			return
		}
	} else if strings.Contains(emailOrID, "@") {
		email = emailOrID
	}

	u, err = r.Client.Users().Get(ctx, uid, email)
	if u.IsValid() {
		uid = u.ID()
	}
	err = translateError(err, u, "user", emailOrID)
	return
}

// Org returns an org, based on the or ID supplied by org.
func (r Resolver) Org(ctx context.Context, org string) (oid sdktypes.OrgID, err error) {
	if org == "" {
		return
	}

	if sdktypes.IsID(org) {
		return sdktypes.ParseOrgID(org)
	}

	n, err := sdktypes.ParseSymbol(org)
	if err != nil {
		return sdktypes.InvalidOrgID, fmt.Errorf("invalid org name %q: %w", org, err)
	}

	o, err := r.Client.Orgs().GetByName(ctx, n)
	if o.IsValid() {
		oid = o.ID()
	}
	err = translateError(err, o, "org", org)
	return
}
