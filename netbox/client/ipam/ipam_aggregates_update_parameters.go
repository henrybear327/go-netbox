// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package ipam

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
	"github.com/go-openapi/swag"

	"github.com/henrybear327/go-netbox/netbox/models"
)

// NewIpamAggregatesUpdateParams creates a new IpamAggregatesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamAggregatesUpdateParams() *IpamAggregatesUpdateParams {
	return &IpamAggregatesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamAggregatesUpdateParamsWithTimeout creates a new IpamAggregatesUpdateParams object
// with the ability to set a timeout on a request.
func NewIpamAggregatesUpdateParamsWithTimeout(timeout time.Duration) *IpamAggregatesUpdateParams {
	return &IpamAggregatesUpdateParams{
		timeout: timeout,
	}
}

// NewIpamAggregatesUpdateParamsWithContext creates a new IpamAggregatesUpdateParams object
// with the ability to set a context for a request.
func NewIpamAggregatesUpdateParamsWithContext(ctx context.Context) *IpamAggregatesUpdateParams {
	return &IpamAggregatesUpdateParams{
		Context: ctx,
	}
}

// NewIpamAggregatesUpdateParamsWithHTTPClient creates a new IpamAggregatesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamAggregatesUpdateParamsWithHTTPClient(client *http.Client) *IpamAggregatesUpdateParams {
	return &IpamAggregatesUpdateParams{
		HTTPClient: client,
	}
}

/*
IpamAggregatesUpdateParams contains all the parameters to send to the API endpoint

	for the ipam aggregates update operation.

	Typically these are written to a http.Request.
*/
type IpamAggregatesUpdateParams struct {

	// Data.
	Data *models.WritableAggregate

	/* ID.

	   A unique integer value identifying this aggregate.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam aggregates update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamAggregatesUpdateParams) WithDefaults() *IpamAggregatesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam aggregates update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamAggregatesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam aggregates update params
func (o *IpamAggregatesUpdateParams) WithTimeout(timeout time.Duration) *IpamAggregatesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam aggregates update params
func (o *IpamAggregatesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam aggregates update params
func (o *IpamAggregatesUpdateParams) WithContext(ctx context.Context) *IpamAggregatesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam aggregates update params
func (o *IpamAggregatesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam aggregates update params
func (o *IpamAggregatesUpdateParams) WithHTTPClient(client *http.Client) *IpamAggregatesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam aggregates update params
func (o *IpamAggregatesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam aggregates update params
func (o *IpamAggregatesUpdateParams) WithData(data *models.WritableAggregate) *IpamAggregatesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam aggregates update params
func (o *IpamAggregatesUpdateParams) SetData(data *models.WritableAggregate) {
	o.Data = data
}

// WithID adds the id to the ipam aggregates update params
func (o *IpamAggregatesUpdateParams) WithID(id int64) *IpamAggregatesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam aggregates update params
func (o *IpamAggregatesUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamAggregatesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
