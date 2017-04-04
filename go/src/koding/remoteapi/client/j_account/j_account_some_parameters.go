package j_account

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

// NewJAccountSomeParams creates a new JAccountSomeParams object
// with the default values initialized.
func NewJAccountSomeParams() *JAccountSomeParams {
	var ()
	return &JAccountSomeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJAccountSomeParamsWithTimeout creates a new JAccountSomeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJAccountSomeParamsWithTimeout(timeout time.Duration) *JAccountSomeParams {
	var ()
	return &JAccountSomeParams{

		timeout: timeout,
	}
}

// NewJAccountSomeParamsWithContext creates a new JAccountSomeParams object
// with the default values initialized, and the ability to set a context for a request
func NewJAccountSomeParamsWithContext(ctx context.Context) *JAccountSomeParams {
	var ()
	return &JAccountSomeParams{

		Context: ctx,
	}
}

/*JAccountSomeParams contains all the parameters to send to the API endpoint
for the j account some operation typically these are written to a http.Request
*/
type JAccountSomeParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the j account some params
func (o *JAccountSomeParams) WithTimeout(timeout time.Duration) *JAccountSomeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j account some params
func (o *JAccountSomeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j account some params
func (o *JAccountSomeParams) WithContext(ctx context.Context) *JAccountSomeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j account some params
func (o *JAccountSomeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the j account some params
func (o *JAccountSomeParams) WithBody(body models.DefaultSelector) *JAccountSomeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the j account some params
func (o *JAccountSomeParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *JAccountSomeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
