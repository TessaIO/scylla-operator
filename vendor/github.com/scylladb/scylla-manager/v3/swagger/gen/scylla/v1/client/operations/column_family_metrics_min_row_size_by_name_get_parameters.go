// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewColumnFamilyMetricsMinRowSizeByNameGetParams creates a new ColumnFamilyMetricsMinRowSizeByNameGetParams object
// with the default values initialized.
func NewColumnFamilyMetricsMinRowSizeByNameGetParams() *ColumnFamilyMetricsMinRowSizeByNameGetParams {
	var ()
	return &ColumnFamilyMetricsMinRowSizeByNameGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewColumnFamilyMetricsMinRowSizeByNameGetParamsWithTimeout creates a new ColumnFamilyMetricsMinRowSizeByNameGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewColumnFamilyMetricsMinRowSizeByNameGetParamsWithTimeout(timeout time.Duration) *ColumnFamilyMetricsMinRowSizeByNameGetParams {
	var ()
	return &ColumnFamilyMetricsMinRowSizeByNameGetParams{

		timeout: timeout,
	}
}

// NewColumnFamilyMetricsMinRowSizeByNameGetParamsWithContext creates a new ColumnFamilyMetricsMinRowSizeByNameGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewColumnFamilyMetricsMinRowSizeByNameGetParamsWithContext(ctx context.Context) *ColumnFamilyMetricsMinRowSizeByNameGetParams {
	var ()
	return &ColumnFamilyMetricsMinRowSizeByNameGetParams{

		Context: ctx,
	}
}

// NewColumnFamilyMetricsMinRowSizeByNameGetParamsWithHTTPClient creates a new ColumnFamilyMetricsMinRowSizeByNameGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewColumnFamilyMetricsMinRowSizeByNameGetParamsWithHTTPClient(client *http.Client) *ColumnFamilyMetricsMinRowSizeByNameGetParams {
	var ()
	return &ColumnFamilyMetricsMinRowSizeByNameGetParams{
		HTTPClient: client,
	}
}

/*
ColumnFamilyMetricsMinRowSizeByNameGetParams contains all the parameters to send to the API endpoint
for the column family metrics min row size by name get operation typically these are written to a http.Request
*/
type ColumnFamilyMetricsMinRowSizeByNameGetParams struct {

	/*Name
	  The column family name in keyspace:name format

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the column family metrics min row size by name get params
func (o *ColumnFamilyMetricsMinRowSizeByNameGetParams) WithTimeout(timeout time.Duration) *ColumnFamilyMetricsMinRowSizeByNameGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the column family metrics min row size by name get params
func (o *ColumnFamilyMetricsMinRowSizeByNameGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the column family metrics min row size by name get params
func (o *ColumnFamilyMetricsMinRowSizeByNameGetParams) WithContext(ctx context.Context) *ColumnFamilyMetricsMinRowSizeByNameGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the column family metrics min row size by name get params
func (o *ColumnFamilyMetricsMinRowSizeByNameGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the column family metrics min row size by name get params
func (o *ColumnFamilyMetricsMinRowSizeByNameGetParams) WithHTTPClient(client *http.Client) *ColumnFamilyMetricsMinRowSizeByNameGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the column family metrics min row size by name get params
func (o *ColumnFamilyMetricsMinRowSizeByNameGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the column family metrics min row size by name get params
func (o *ColumnFamilyMetricsMinRowSizeByNameGetParams) WithName(name string) *ColumnFamilyMetricsMinRowSizeByNameGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the column family metrics min row size by name get params
func (o *ColumnFamilyMetricsMinRowSizeByNameGetParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ColumnFamilyMetricsMinRowSizeByNameGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}