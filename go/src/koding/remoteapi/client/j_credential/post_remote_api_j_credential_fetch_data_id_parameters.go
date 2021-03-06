package j_credential

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// NewPostRemoteAPIJCredentialFetchDataIDParams creates a new PostRemoteAPIJCredentialFetchDataIDParams object
// with the default values initialized.
func NewPostRemoteAPIJCredentialFetchDataIDParams() *PostRemoteAPIJCredentialFetchDataIDParams {
	var ()
	return &PostRemoteAPIJCredentialFetchDataIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJCredentialFetchDataIDParamsWithTimeout creates a new PostRemoteAPIJCredentialFetchDataIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJCredentialFetchDataIDParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJCredentialFetchDataIDParams {
	var ()
	return &PostRemoteAPIJCredentialFetchDataIDParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJCredentialFetchDataIDParamsWithContext creates a new PostRemoteAPIJCredentialFetchDataIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJCredentialFetchDataIDParamsWithContext(ctx context.Context) *PostRemoteAPIJCredentialFetchDataIDParams {
	var ()
	return &PostRemoteAPIJCredentialFetchDataIDParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJCredentialFetchDataIDParams contains all the parameters to send to the API endpoint
for the post remote API j credential fetch data ID operation typically these are written to a http.Request
*/
type PostRemoteAPIJCredentialFetchDataIDParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector
	/*ID
	  Mongo ID of target instance

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API j credential fetch data ID params
func (o *PostRemoteAPIJCredentialFetchDataIDParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJCredentialFetchDataIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j credential fetch data ID params
func (o *PostRemoteAPIJCredentialFetchDataIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j credential fetch data ID params
func (o *PostRemoteAPIJCredentialFetchDataIDParams) WithContext(ctx context.Context) *PostRemoteAPIJCredentialFetchDataIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j credential fetch data ID params
func (o *PostRemoteAPIJCredentialFetchDataIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API j credential fetch data ID params
func (o *PostRemoteAPIJCredentialFetchDataIDParams) WithBody(body models.DefaultSelector) *PostRemoteAPIJCredentialFetchDataIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API j credential fetch data ID params
func (o *PostRemoteAPIJCredentialFetchDataIDParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WithID adds the id to the post remote API j credential fetch data ID params
func (o *PostRemoteAPIJCredentialFetchDataIDParams) WithID(id string) *PostRemoteAPIJCredentialFetchDataIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post remote API j credential fetch data ID params
func (o *PostRemoteAPIJCredentialFetchDataIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJCredentialFetchDataIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
