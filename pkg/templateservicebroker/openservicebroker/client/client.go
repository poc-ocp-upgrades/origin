package client

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	goformat "fmt"
	"github.com/openshift/origin/pkg/templateservicebroker/openservicebroker/api"
	"golang.org/x/net/context"
	"io"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/authentication/user"
	"net/http"
	goos "os"
	godefaultruntime "runtime"
	"time"
	gotime "time"
)

type Client interface {
	Client() *http.Client
	Catalog(ctx context.Context) (*api.CatalogResponse, error)
	Provision(ctx context.Context, u user.Info, instanceID string, preq *api.ProvisionRequest) (*api.ProvisionResponse, error)
	Deprovision(ctx context.Context, u user.Info, instanceID string) error
	Bind(ctx context.Context, u user.Info, instanceID, bindingID string, breq *api.BindRequest) (*api.BindResponse, error)
	Unbind(ctx context.Context, u user.Info, instanceID, bindingID string) error
}
type client struct {
	cli  *http.Client
	root string
}

func NewClient(cli *http.Client, root string) Client {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &client{cli: cli, root: root}
}

type ServerError struct {
	StatusCode  int
	Description string
}

func (e *ServerError) Error() string {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return fmt.Sprintf("%s: %s", http.StatusText(e.StatusCode), e.Description)
}
func newServerError(statusCode int, description string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return &ServerError{StatusCode: statusCode, Description: description}
}
func (c *client) Client() *http.Client {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	return c.cli
}
func OriginatingIdentityHeader(u user.Info) (string, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	templatereq := api.ConvertUserToTemplateInstanceRequester(u)
	b, err := json.Marshal(&templatereq)
	if err != nil {
		return "", err
	}
	encodeVal := base64.StdEncoding.EncodeToString(b)
	return api.OriginatingIdentitySchemeKubernetes + " " + encodeVal, nil
}
func (c *client) Catalog(ctx context.Context) (*api.CatalogResponse, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	req, err := http.NewRequest(http.MethodGet, c.root+"/v2/catalog", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(api.XBrokerAPIVersion, api.APIVersion)
	resp, err := c.cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.Header.Get("Content-Type") != "application/json" {
		return nil, newServerError(resp.StatusCode, "invalid content type")
	}
	d := json.NewDecoder(resp.Body)
	if resp.StatusCode == http.StatusOK {
		var r *api.CatalogResponse
		err = d.Decode(&r)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
	var r *api.ErrorResponse
	err = d.Decode(&r)
	if err != nil {
		return nil, err
	}
	return nil, newServerError(resp.StatusCode, r.Description)
}
func (c *client) Provision(ctx context.Context, u user.Info, instanceID string, preq *api.ProvisionRequest) (*api.ProvisionResponse, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if errs := api.ValidateUUID(field.NewPath("instanceID"), instanceID); len(errs) > 0 {
		return nil, errs.ToAggregate()
	}
	pr, pw := io.Pipe()
	go func() {
		e := json.NewEncoder(pw)
		pw.CloseWithError(e.Encode(preq))
	}()
	req, err := http.NewRequest(http.MethodPut, c.root+"/v2/service_instances/"+instanceID+"?accepts_incomplete=true", pr)
	if err != nil {
		return nil, err
	}
	req.Header.Add(api.XBrokerAPIVersion, api.APIVersion)
	req.Header.Add("Content-Type", "application/json")
	identity, err := OriginatingIdentityHeader(u)
	if err != nil {
		return nil, err
	}
	req.Header.Add(api.XBrokerAPIOriginatingIdentity, identity)
	resp, err := c.cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.Header.Get("Content-Type") != "application/json" {
		return nil, newServerError(resp.StatusCode, "invalid content type")
	}
	d := json.NewDecoder(resp.Body)
	if resp.StatusCode == http.StatusCreated || resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusAccepted {
		var r *api.ProvisionResponse
		err = d.Decode(&r)
		if err != nil {
			return nil, err
		}
		if resp.StatusCode == http.StatusAccepted {
			var state api.LastOperationState
			state, err = c.WaitForOperation(ctx, u, instanceID, r.Operation)
			if err != nil {
				return nil, err
			}
			if state != api.LastOperationStateSucceeded {
				return nil, fmt.Errorf("operation returned state %s", string(state))
			}
		}
		return r, nil
	}
	var r *api.ErrorResponse
	err = d.Decode(&r)
	if err != nil {
		return nil, err
	}
	return nil, newServerError(resp.StatusCode, r.Description)
}
func (c *client) Deprovision(ctx context.Context, u user.Info, instanceID string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if errs := api.ValidateUUID(field.NewPath("instanceID"), instanceID); len(errs) > 0 {
		return errs.ToAggregate()
	}
	req, err := http.NewRequest(http.MethodDelete, c.root+"/v2/service_instances/"+instanceID+"?accepts_incomplete=true", nil)
	if err != nil {
		return err
	}
	req.Header.Add(api.XBrokerAPIVersion, api.APIVersion)
	identity, err := OriginatingIdentityHeader(u)
	if err != nil {
		return err
	}
	req.Header.Add(api.XBrokerAPIOriginatingIdentity, identity)
	resp, err := c.cli.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.Header.Get("Content-Type") != "application/json" {
		return newServerError(resp.StatusCode, "invalid content type")
	}
	d := json.NewDecoder(resp.Body)
	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusAccepted || resp.StatusCode == http.StatusGone {
		var r *api.DeprovisionResponse
		err = d.Decode(&r)
		if err != nil {
			return err
		}
		if resp.StatusCode == http.StatusAccepted {
			var state api.LastOperationState
			state, err = c.WaitForOperation(ctx, u, instanceID, r.Operation)
			if err != nil {
				return err
			}
			if state != api.LastOperationStateSucceeded {
				return fmt.Errorf("operation returned state %s", string(state))
			}
		}
		return nil
	}
	var r *api.ErrorResponse
	err = d.Decode(&r)
	if err != nil {
		return err
	}
	return newServerError(resp.StatusCode, r.Description)
}
func (c *client) LastOperation(ctx context.Context, u user.Info, instanceID string, operation api.Operation) (*api.LastOperationResponse, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if errs := api.ValidateUUID(field.NewPath("instanceID"), instanceID); len(errs) > 0 {
		return nil, errs.ToAggregate()
	}
	req, err := http.NewRequest(http.MethodGet, c.root+"/v2/service_instances/"+instanceID+"/last_operation?operation="+string(operation), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(api.XBrokerAPIVersion, api.APIVersion)
	identity, err := OriginatingIdentityHeader(u)
	if err != nil {
		return nil, err
	}
	req.Header.Add(api.XBrokerAPIOriginatingIdentity, identity)
	resp, err := c.cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.Header.Get("Content-Type") != "application/json" {
		return nil, newServerError(resp.StatusCode, "invalid content type")
	}
	d := json.NewDecoder(resp.Body)
	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusGone {
		var r *api.LastOperationResponse
		err = d.Decode(&r)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
	var r *api.ErrorResponse
	err = d.Decode(&r)
	if err != nil {
		return nil, err
	}
	return nil, newServerError(resp.StatusCode, r.Description)
}
func (c *client) WaitForOperation(ctx context.Context, u user.Info, instanceID string, operation api.Operation) (api.LastOperationState, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	done := ctx.Done()
	for {
		op, err := c.LastOperation(ctx, u, instanceID, operation)
		if err != nil {
			return api.LastOperationStateFailed, err
		}
		if op.State != api.LastOperationStateInProgress {
			return op.State, nil
		}
		select {
		case <-done:
			return api.LastOperationStateFailed, ctx.Err()
		default:
		}
		time.Sleep(1 * time.Second)
	}
}
func (c *client) Bind(ctx context.Context, u user.Info, instanceID, bindingID string, breq *api.BindRequest) (*api.BindResponse, error) {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if errs := api.ValidateUUID(field.NewPath("instanceID"), instanceID); len(errs) > 0 {
		return nil, errs.ToAggregate()
	}
	if errs := api.ValidateUUID(field.NewPath("bindingID"), bindingID); len(errs) > 0 {
		return nil, errs.ToAggregate()
	}
	pr, pw := io.Pipe()
	go func() {
		e := json.NewEncoder(pw)
		pw.CloseWithError(e.Encode(breq))
	}()
	req, err := http.NewRequest(http.MethodPut, c.root+"/v2/service_instances/"+instanceID+"/service_bindings/"+bindingID, pr)
	if err != nil {
		return nil, err
	}
	req.Header.Add(api.XBrokerAPIVersion, api.APIVersion)
	req.Header.Add("Content-Type", "application/json")
	identity, err := OriginatingIdentityHeader(u)
	if err != nil {
		return nil, err
	}
	req.Header.Add(api.XBrokerAPIOriginatingIdentity, identity)
	resp, err := c.cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.Header.Get("Content-Type") != "application/json" {
		return nil, newServerError(resp.StatusCode, "invalid content type")
	}
	d := json.NewDecoder(resp.Body)
	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated {
		var r *api.BindResponse
		err = d.Decode(&r)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
	var r *api.ErrorResponse
	err = d.Decode(&r)
	if err != nil {
		return nil, err
	}
	return nil, newServerError(resp.StatusCode, r.Description)
}
func (c *client) Unbind(ctx context.Context, u user.Info, instanceID, bindingID string) error {
	_logClusterCodePath("Entered function: ")
	defer _logClusterCodePath("Exited function: ")
	if errs := api.ValidateUUID(field.NewPath("instanceID"), instanceID); len(errs) > 0 {
		return errs.ToAggregate()
	}
	if errs := api.ValidateUUID(field.NewPath("bindingID"), bindingID); len(errs) > 0 {
		return errs.ToAggregate()
	}
	req, err := http.NewRequest(http.MethodDelete, c.root+"/v2/service_instances/"+instanceID+"/service_bindings/"+bindingID, nil)
	if err != nil {
		return err
	}
	req.Header.Add(api.XBrokerAPIVersion, api.APIVersion)
	identity, err := OriginatingIdentityHeader(u)
	if err != nil {
		return err
	}
	req.Header.Add(api.XBrokerAPIOriginatingIdentity, identity)
	resp, err := c.cli.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.Header.Get("Content-Type") != "application/json" {
		return newServerError(resp.StatusCode, "invalid content type")
	}
	d := json.NewDecoder(resp.Body)
	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusGone {
		var r *api.UnbindResponse
		err = d.Decode(&r)
		if err != nil {
			return err
		}
		return nil
	}
	var r *api.ErrorResponse
	err = d.Decode(&r)
	if err != nil {
		return err
	}
	return newServerError(resp.StatusCode, r.Description)
}
func _logClusterCodePath(op string) {
	pc, _, _, _ := godefaultruntime.Caller(1)
	goformat.Fprintf(goos.Stderr, "[%v][ANALYTICS] %s%s\n", gotime.Now().UTC(), op, godefaultruntime.FuncForPC(pc).Name())
}
