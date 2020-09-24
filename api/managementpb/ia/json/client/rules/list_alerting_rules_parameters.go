// Code generated by go-swagger; DO NOT EDIT.

package rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewListAlertingRulesParams creates a new ListAlertingRulesParams object
// with the default values initialized.
func NewListAlertingRulesParams() *ListAlertingRulesParams {
	var ()
	return &ListAlertingRulesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListAlertingRulesParamsWithTimeout creates a new ListAlertingRulesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListAlertingRulesParamsWithTimeout(timeout time.Duration) *ListAlertingRulesParams {
	var ()
	return &ListAlertingRulesParams{

		timeout: timeout,
	}
}

// NewListAlertingRulesParamsWithContext creates a new ListAlertingRulesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListAlertingRulesParamsWithContext(ctx context.Context) *ListAlertingRulesParams {
	var ()
	return &ListAlertingRulesParams{

		Context: ctx,
	}
}

// NewListAlertingRulesParamsWithHTTPClient creates a new ListAlertingRulesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListAlertingRulesParamsWithHTTPClient(client *http.Client) *ListAlertingRulesParams {
	var ()
	return &ListAlertingRulesParams{
		HTTPClient: client,
	}
}

/*ListAlertingRulesParams contains all the parameters to send to the API endpoint
for the list alerting rules operation typically these are written to a http.Request
*/
type ListAlertingRulesParams struct {

	/*Body*/
	Body interface{}

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list alerting rules params
func (o *ListAlertingRulesParams) WithTimeout(timeout time.Duration) *ListAlertingRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list alerting rules params
func (o *ListAlertingRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list alerting rules params
func (o *ListAlertingRulesParams) WithContext(ctx context.Context) *ListAlertingRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list alerting rules params
func (o *ListAlertingRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list alerting rules params
func (o *ListAlertingRulesParams) WithHTTPClient(client *http.Client) *ListAlertingRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list alerting rules params
func (o *ListAlertingRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the list alerting rules params
func (o *ListAlertingRulesParams) WithBody(body interface{}) *ListAlertingRulesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the list alerting rules params
func (o *ListAlertingRulesParams) SetBody(body interface{}) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ListAlertingRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
