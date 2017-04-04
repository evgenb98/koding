package j_group

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

// NewJGroupFetchPermissionsParams creates a new JGroupFetchPermissionsParams object
// with the default values initialized.
func NewJGroupFetchPermissionsParams() *JGroupFetchPermissionsParams {
	var ()
	return &JGroupFetchPermissionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJGroupFetchPermissionsParamsWithTimeout creates a new JGroupFetchPermissionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJGroupFetchPermissionsParamsWithTimeout(timeout time.Duration) *JGroupFetchPermissionsParams {
	var ()
	return &JGroupFetchPermissionsParams{

		timeout: timeout,
	}
}

// NewJGroupFetchPermissionsParamsWithContext creates a new JGroupFetchPermissionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewJGroupFetchPermissionsParamsWithContext(ctx context.Context) *JGroupFetchPermissionsParams {
	var ()
	return &JGroupFetchPermissionsParams{

		Context: ctx,
	}
}

/*JGroupFetchPermissionsParams contains all the parameters to send to the API endpoint
for the j group fetch permissions operation typically these are written to a http.Request
*/
type JGroupFetchPermissionsParams struct {

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

// WithTimeout adds the timeout to the j group fetch permissions params
func (o *JGroupFetchPermissionsParams) WithTimeout(timeout time.Duration) *JGroupFetchPermissionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j group fetch permissions params
func (o *JGroupFetchPermissionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j group fetch permissions params
func (o *JGroupFetchPermissionsParams) WithContext(ctx context.Context) *JGroupFetchPermissionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j group fetch permissions params
func (o *JGroupFetchPermissionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the j group fetch permissions params
func (o *JGroupFetchPermissionsParams) WithBody(body models.DefaultSelector) *JGroupFetchPermissionsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the j group fetch permissions params
func (o *JGroupFetchPermissionsParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WithID adds the id to the j group fetch permissions params
func (o *JGroupFetchPermissionsParams) WithID(id string) *JGroupFetchPermissionsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the j group fetch permissions params
func (o *JGroupFetchPermissionsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *JGroupFetchPermissionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
