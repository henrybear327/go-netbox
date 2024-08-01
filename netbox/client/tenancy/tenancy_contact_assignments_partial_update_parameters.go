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

package tenancy

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

// NewTenancyContactAssignmentsPartialUpdateParams creates a new TenancyContactAssignmentsPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTenancyContactAssignmentsPartialUpdateParams() *TenancyContactAssignmentsPartialUpdateParams {
	return &TenancyContactAssignmentsPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTenancyContactAssignmentsPartialUpdateParamsWithTimeout creates a new TenancyContactAssignmentsPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewTenancyContactAssignmentsPartialUpdateParamsWithTimeout(timeout time.Duration) *TenancyContactAssignmentsPartialUpdateParams {
	return &TenancyContactAssignmentsPartialUpdateParams{
		timeout: timeout,
	}
}

// NewTenancyContactAssignmentsPartialUpdateParamsWithContext creates a new TenancyContactAssignmentsPartialUpdateParams object
// with the ability to set a context for a request.
func NewTenancyContactAssignmentsPartialUpdateParamsWithContext(ctx context.Context) *TenancyContactAssignmentsPartialUpdateParams {
	return &TenancyContactAssignmentsPartialUpdateParams{
		Context: ctx,
	}
}

// NewTenancyContactAssignmentsPartialUpdateParamsWithHTTPClient creates a new TenancyContactAssignmentsPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewTenancyContactAssignmentsPartialUpdateParamsWithHTTPClient(client *http.Client) *TenancyContactAssignmentsPartialUpdateParams {
	return &TenancyContactAssignmentsPartialUpdateParams{
		HTTPClient: client,
	}
}

/*
TenancyContactAssignmentsPartialUpdateParams contains all the parameters to send to the API endpoint

	for the tenancy contact assignments partial update operation.

	Typically these are written to a http.Request.
*/
type TenancyContactAssignmentsPartialUpdateParams struct {

	// Data.
	Data *models.WritableContactAssignment

	/* ID.

	   A unique integer value identifying this contact assignment.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tenancy contact assignments partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyContactAssignmentsPartialUpdateParams) WithDefaults() *TenancyContactAssignmentsPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tenancy contact assignments partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyContactAssignmentsPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tenancy contact assignments partial update params
func (o *TenancyContactAssignmentsPartialUpdateParams) WithTimeout(timeout time.Duration) *TenancyContactAssignmentsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tenancy contact assignments partial update params
func (o *TenancyContactAssignmentsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tenancy contact assignments partial update params
func (o *TenancyContactAssignmentsPartialUpdateParams) WithContext(ctx context.Context) *TenancyContactAssignmentsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tenancy contact assignments partial update params
func (o *TenancyContactAssignmentsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tenancy contact assignments partial update params
func (o *TenancyContactAssignmentsPartialUpdateParams) WithHTTPClient(client *http.Client) *TenancyContactAssignmentsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tenancy contact assignments partial update params
func (o *TenancyContactAssignmentsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the tenancy contact assignments partial update params
func (o *TenancyContactAssignmentsPartialUpdateParams) WithData(data *models.WritableContactAssignment) *TenancyContactAssignmentsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the tenancy contact assignments partial update params
func (o *TenancyContactAssignmentsPartialUpdateParams) SetData(data *models.WritableContactAssignment) {
	o.Data = data
}

// WithID adds the id to the tenancy contact assignments partial update params
func (o *TenancyContactAssignmentsPartialUpdateParams) WithID(id int64) *TenancyContactAssignmentsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the tenancy contact assignments partial update params
func (o *TenancyContactAssignmentsPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *TenancyContactAssignmentsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
