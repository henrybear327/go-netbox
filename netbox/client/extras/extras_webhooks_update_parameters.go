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

package extras

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

// NewExtrasWebhooksUpdateParams creates a new ExtrasWebhooksUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasWebhooksUpdateParams() *ExtrasWebhooksUpdateParams {
	return &ExtrasWebhooksUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasWebhooksUpdateParamsWithTimeout creates a new ExtrasWebhooksUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasWebhooksUpdateParamsWithTimeout(timeout time.Duration) *ExtrasWebhooksUpdateParams {
	return &ExtrasWebhooksUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasWebhooksUpdateParamsWithContext creates a new ExtrasWebhooksUpdateParams object
// with the ability to set a context for a request.
func NewExtrasWebhooksUpdateParamsWithContext(ctx context.Context) *ExtrasWebhooksUpdateParams {
	return &ExtrasWebhooksUpdateParams{
		Context: ctx,
	}
}

// NewExtrasWebhooksUpdateParamsWithHTTPClient creates a new ExtrasWebhooksUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasWebhooksUpdateParamsWithHTTPClient(client *http.Client) *ExtrasWebhooksUpdateParams {
	return &ExtrasWebhooksUpdateParams{
		HTTPClient: client,
	}
}

/*
ExtrasWebhooksUpdateParams contains all the parameters to send to the API endpoint

	for the extras webhooks update operation.

	Typically these are written to a http.Request.
*/
type ExtrasWebhooksUpdateParams struct {

	// Data.
	Data *models.Webhook

	/* ID.

	   A unique integer value identifying this webhook.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras webhooks update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasWebhooksUpdateParams) WithDefaults() *ExtrasWebhooksUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras webhooks update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasWebhooksUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras webhooks update params
func (o *ExtrasWebhooksUpdateParams) WithTimeout(timeout time.Duration) *ExtrasWebhooksUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras webhooks update params
func (o *ExtrasWebhooksUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras webhooks update params
func (o *ExtrasWebhooksUpdateParams) WithContext(ctx context.Context) *ExtrasWebhooksUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras webhooks update params
func (o *ExtrasWebhooksUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras webhooks update params
func (o *ExtrasWebhooksUpdateParams) WithHTTPClient(client *http.Client) *ExtrasWebhooksUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras webhooks update params
func (o *ExtrasWebhooksUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras webhooks update params
func (o *ExtrasWebhooksUpdateParams) WithData(data *models.Webhook) *ExtrasWebhooksUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras webhooks update params
func (o *ExtrasWebhooksUpdateParams) SetData(data *models.Webhook) {
	o.Data = data
}

// WithID adds the id to the extras webhooks update params
func (o *ExtrasWebhooksUpdateParams) WithID(id int64) *ExtrasWebhooksUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras webhooks update params
func (o *ExtrasWebhooksUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasWebhooksUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
