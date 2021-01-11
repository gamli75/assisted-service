// Code generated by go-swagger; DO NOT EDIT.

package operators

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

// NewListOperatorPropertiesParams creates a new ListOperatorPropertiesParams object
// with the default values initialized.
func NewListOperatorPropertiesParams() *ListOperatorPropertiesParams {
	var ()
	return &ListOperatorPropertiesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListOperatorPropertiesParamsWithTimeout creates a new ListOperatorPropertiesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListOperatorPropertiesParamsWithTimeout(timeout time.Duration) *ListOperatorPropertiesParams {
	var ()
	return &ListOperatorPropertiesParams{

		timeout: timeout,
	}
}

// NewListOperatorPropertiesParamsWithContext creates a new ListOperatorPropertiesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListOperatorPropertiesParamsWithContext(ctx context.Context) *ListOperatorPropertiesParams {
	var ()
	return &ListOperatorPropertiesParams{

		Context: ctx,
	}
}

// NewListOperatorPropertiesParamsWithHTTPClient creates a new ListOperatorPropertiesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListOperatorPropertiesParamsWithHTTPClient(client *http.Client) *ListOperatorPropertiesParams {
	var ()
	return &ListOperatorPropertiesParams{
		HTTPClient: client,
	}
}

/*ListOperatorPropertiesParams contains all the parameters to send to the API endpoint
for the list operator properties operation typically these are written to a http.Request
*/
type ListOperatorPropertiesParams struct {

	/*OperatorType
	  The operator type.

	*/
	OperatorType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list operator properties params
func (o *ListOperatorPropertiesParams) WithTimeout(timeout time.Duration) *ListOperatorPropertiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list operator properties params
func (o *ListOperatorPropertiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list operator properties params
func (o *ListOperatorPropertiesParams) WithContext(ctx context.Context) *ListOperatorPropertiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list operator properties params
func (o *ListOperatorPropertiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list operator properties params
func (o *ListOperatorPropertiesParams) WithHTTPClient(client *http.Client) *ListOperatorPropertiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list operator properties params
func (o *ListOperatorPropertiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOperatorType adds the operatorType to the list operator properties params
func (o *ListOperatorPropertiesParams) WithOperatorType(operatorType string) *ListOperatorPropertiesParams {
	o.SetOperatorType(operatorType)
	return o
}

// SetOperatorType adds the operatorType to the list operator properties params
func (o *ListOperatorPropertiesParams) SetOperatorType(operatorType string) {
	o.OperatorType = operatorType
}

// WriteToRequest writes these params to a swagger request
func (o *ListOperatorPropertiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param operator_type
	if err := r.SetPathParam("operator_type", o.OperatorType); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
