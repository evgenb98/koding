package j_url_alias

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

// NewJURLAliasCreateParams creates a new JURLAliasCreateParams object
// with the default values initialized.
func NewJURLAliasCreateParams() *JURLAliasCreateParams {
	var ()
	return &JURLAliasCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJURLAliasCreateParamsWithTimeout creates a new JURLAliasCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJURLAliasCreateParamsWithTimeout(timeout time.Duration) *JURLAliasCreateParams {
	var ()
	return &JURLAliasCreateParams{

		timeout: timeout,
	}
}

// NewJURLAliasCreateParamsWithContext creates a new JURLAliasCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewJURLAliasCreateParamsWithContext(ctx context.Context) *JURLAliasCreateParams {
	var ()
	return &JURLAliasCreateParams{

		Context: ctx,
	}
}

/*JURLAliasCreateParams contains all the parameters to send to the API endpoint
for the j Url alias create operation typically these are written to a http.Request
*/
type JURLAliasCreateParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the j Url alias create params
func (o *JURLAliasCreateParams) WithTimeout(timeout time.Duration) *JURLAliasCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j Url alias create params
func (o *JURLAliasCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j Url alias create params
func (o *JURLAliasCreateParams) WithContext(ctx context.Context) *JURLAliasCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j Url alias create params
func (o *JURLAliasCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the j Url alias create params
func (o *JURLAliasCreateParams) WithBody(body models.DefaultSelector) *JURLAliasCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the j Url alias create params
func (o *JURLAliasCreateParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *JURLAliasCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
