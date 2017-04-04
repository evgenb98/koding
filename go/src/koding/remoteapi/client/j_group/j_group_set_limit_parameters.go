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

// NewJGroupSetLimitParams creates a new JGroupSetLimitParams object
// with the default values initialized.
func NewJGroupSetLimitParams() *JGroupSetLimitParams {
	var ()
	return &JGroupSetLimitParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJGroupSetLimitParamsWithTimeout creates a new JGroupSetLimitParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJGroupSetLimitParamsWithTimeout(timeout time.Duration) *JGroupSetLimitParams {
	var ()
	return &JGroupSetLimitParams{

		timeout: timeout,
	}
}

// NewJGroupSetLimitParamsWithContext creates a new JGroupSetLimitParams object
// with the default values initialized, and the ability to set a context for a request
func NewJGroupSetLimitParamsWithContext(ctx context.Context) *JGroupSetLimitParams {
	var ()
	return &JGroupSetLimitParams{

		Context: ctx,
	}
}

/*JGroupSetLimitParams contains all the parameters to send to the API endpoint
for the j group set limit operation typically these are written to a http.Request
*/
type JGroupSetLimitParams struct {

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

// WithTimeout adds the timeout to the j group set limit params
func (o *JGroupSetLimitParams) WithTimeout(timeout time.Duration) *JGroupSetLimitParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j group set limit params
func (o *JGroupSetLimitParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j group set limit params
func (o *JGroupSetLimitParams) WithContext(ctx context.Context) *JGroupSetLimitParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j group set limit params
func (o *JGroupSetLimitParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the j group set limit params
func (o *JGroupSetLimitParams) WithBody(body models.DefaultSelector) *JGroupSetLimitParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the j group set limit params
func (o *JGroupSetLimitParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WithID adds the id to the j group set limit params
func (o *JGroupSetLimitParams) WithID(id string) *JGroupSetLimitParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the j group set limit params
func (o *JGroupSetLimitParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *JGroupSetLimitParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
