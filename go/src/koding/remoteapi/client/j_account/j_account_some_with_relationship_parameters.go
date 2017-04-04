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

// NewJAccountSomeWithRelationshipParams creates a new JAccountSomeWithRelationshipParams object
// with the default values initialized.
func NewJAccountSomeWithRelationshipParams() *JAccountSomeWithRelationshipParams {
	var ()
	return &JAccountSomeWithRelationshipParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJAccountSomeWithRelationshipParamsWithTimeout creates a new JAccountSomeWithRelationshipParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJAccountSomeWithRelationshipParamsWithTimeout(timeout time.Duration) *JAccountSomeWithRelationshipParams {
	var ()
	return &JAccountSomeWithRelationshipParams{

		timeout: timeout,
	}
}

// NewJAccountSomeWithRelationshipParamsWithContext creates a new JAccountSomeWithRelationshipParams object
// with the default values initialized, and the ability to set a context for a request
func NewJAccountSomeWithRelationshipParamsWithContext(ctx context.Context) *JAccountSomeWithRelationshipParams {
	var ()
	return &JAccountSomeWithRelationshipParams{

		Context: ctx,
	}
}

/*JAccountSomeWithRelationshipParams contains all the parameters to send to the API endpoint
for the j account some with relationship operation typically these are written to a http.Request
*/
type JAccountSomeWithRelationshipParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the j account some with relationship params
func (o *JAccountSomeWithRelationshipParams) WithTimeout(timeout time.Duration) *JAccountSomeWithRelationshipParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j account some with relationship params
func (o *JAccountSomeWithRelationshipParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j account some with relationship params
func (o *JAccountSomeWithRelationshipParams) WithContext(ctx context.Context) *JAccountSomeWithRelationshipParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j account some with relationship params
func (o *JAccountSomeWithRelationshipParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the j account some with relationship params
func (o *JAccountSomeWithRelationshipParams) WithBody(body models.DefaultSelector) *JAccountSomeWithRelationshipParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the j account some with relationship params
func (o *JAccountSomeWithRelationshipParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *JAccountSomeWithRelationshipParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
