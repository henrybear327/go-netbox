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

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/henrybear327/go-netbox/netbox/models"
)

// DcimSitesListReader is a Reader for the DcimSitesList structure.
type DcimSitesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimSitesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimSitesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimSitesListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimSitesListOK creates a DcimSitesListOK with default headers values
func NewDcimSitesListOK() *DcimSitesListOK {
	return &DcimSitesListOK{}
}

/*
DcimSitesListOK describes a response with status code 200, with default header values.

DcimSitesListOK dcim sites list o k
*/
type DcimSitesListOK struct {
	Payload *DcimSitesListOKBody
}

// IsSuccess returns true when this dcim sites list o k response has a 2xx status code
func (o *DcimSitesListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim sites list o k response has a 3xx status code
func (o *DcimSitesListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim sites list o k response has a 4xx status code
func (o *DcimSitesListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim sites list o k response has a 5xx status code
func (o *DcimSitesListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim sites list o k response a status code equal to that given
func (o *DcimSitesListOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimSitesListOK) Error() string {
	return fmt.Sprintf("[GET /dcim/sites/][%d] dcimSitesListOK  %+v", 200, o.Payload)
}

func (o *DcimSitesListOK) String() string {
	return fmt.Sprintf("[GET /dcim/sites/][%d] dcimSitesListOK  %+v", 200, o.Payload)
}

func (o *DcimSitesListOK) GetPayload() *DcimSitesListOKBody {
	return o.Payload
}

func (o *DcimSitesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DcimSitesListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimSitesListDefault creates a DcimSitesListDefault with default headers values
func NewDcimSitesListDefault(code int) *DcimSitesListDefault {
	return &DcimSitesListDefault{
		_statusCode: code,
	}
}

/*
DcimSitesListDefault describes a response with status code -1, with default header values.

DcimSitesListDefault dcim sites list default
*/
type DcimSitesListDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim sites list default response
func (o *DcimSitesListDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim sites list default response has a 2xx status code
func (o *DcimSitesListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim sites list default response has a 3xx status code
func (o *DcimSitesListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim sites list default response has a 4xx status code
func (o *DcimSitesListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim sites list default response has a 5xx status code
func (o *DcimSitesListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim sites list default response a status code equal to that given
func (o *DcimSitesListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimSitesListDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/sites/][%d] dcim_sites_list default  %+v", o._statusCode, o.Payload)
}

func (o *DcimSitesListDefault) String() string {
	return fmt.Sprintf("[GET /dcim/sites/][%d] dcim_sites_list default  %+v", o._statusCode, o.Payload)
}

func (o *DcimSitesListDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimSitesListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
DcimSitesListOKBody dcim sites list o k body
swagger:model DcimSitesListOKBody
*/
type DcimSitesListOKBody struct {

	// count
	// Required: true
	Count *int64 `json:"count"`

	// next
	// Format: uri
	Next *strfmt.URI `json:"next,omitempty"`

	// previous
	// Format: uri
	Previous *strfmt.URI `json:"previous,omitempty"`

	// results
	// Required: true
	Results []*models.Site `json:"results"`
}

// Validate validates this dcim sites list o k body
func (o *DcimSitesListOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePrevious(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DcimSitesListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("dcimSitesListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *DcimSitesListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimSitesListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimSitesListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimSitesListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimSitesListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("dcimSitesListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimSitesListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dcimSitesListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this dcim sites list o k body based on the context it is used
func (o *DcimSitesListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DcimSitesListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {
			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimSitesListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dcimSitesListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DcimSitesListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DcimSitesListOKBody) UnmarshalBinary(b []byte) error {
	var res DcimSitesListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
