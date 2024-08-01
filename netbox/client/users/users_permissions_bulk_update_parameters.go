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

package users

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

	"github.com/henrybear327/go-netbox/netbox/models"
)

// NewUsersPermissionsBulkUpdateParams creates a new UsersPermissionsBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUsersPermissionsBulkUpdateParams() *UsersPermissionsBulkUpdateParams {
	return &UsersPermissionsBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUsersPermissionsBulkUpdateParamsWithTimeout creates a new UsersPermissionsBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewUsersPermissionsBulkUpdateParamsWithTimeout(timeout time.Duration) *UsersPermissionsBulkUpdateParams {
	return &UsersPermissionsBulkUpdateParams{
		timeout: timeout,
	}
}

// NewUsersPermissionsBulkUpdateParamsWithContext creates a new UsersPermissionsBulkUpdateParams object
// with the ability to set a context for a request.
func NewUsersPermissionsBulkUpdateParamsWithContext(ctx context.Context) *UsersPermissionsBulkUpdateParams {
	return &UsersPermissionsBulkUpdateParams{
		Context: ctx,
	}
}

// NewUsersPermissionsBulkUpdateParamsWithHTTPClient creates a new UsersPermissionsBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewUsersPermissionsBulkUpdateParamsWithHTTPClient(client *http.Client) *UsersPermissionsBulkUpdateParams {
	return &UsersPermissionsBulkUpdateParams{
		HTTPClient: client,
	}
}

/*
UsersPermissionsBulkUpdateParams contains all the parameters to send to the API endpoint

	for the users permissions bulk update operation.

	Typically these are written to a http.Request.
*/
type UsersPermissionsBulkUpdateParams struct {

	// Data.
	Data *models.WritableObjectPermission

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the users permissions bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersPermissionsBulkUpdateParams) WithDefaults() *UsersPermissionsBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the users permissions bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersPermissionsBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the users permissions bulk update params
func (o *UsersPermissionsBulkUpdateParams) WithTimeout(timeout time.Duration) *UsersPermissionsBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users permissions bulk update params
func (o *UsersPermissionsBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users permissions bulk update params
func (o *UsersPermissionsBulkUpdateParams) WithContext(ctx context.Context) *UsersPermissionsBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users permissions bulk update params
func (o *UsersPermissionsBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users permissions bulk update params
func (o *UsersPermissionsBulkUpdateParams) WithHTTPClient(client *http.Client) *UsersPermissionsBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users permissions bulk update params
func (o *UsersPermissionsBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the users permissions bulk update params
func (o *UsersPermissionsBulkUpdateParams) WithData(data *models.WritableObjectPermission) *UsersPermissionsBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the users permissions bulk update params
func (o *UsersPermissionsBulkUpdateParams) SetData(data *models.WritableObjectPermission) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *UsersPermissionsBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
