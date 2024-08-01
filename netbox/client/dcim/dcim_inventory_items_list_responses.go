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

// DcimInventoryItemsListReader is a Reader for the DcimInventoryItemsList structure.
type DcimInventoryItemsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInventoryItemsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInventoryItemsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimInventoryItemsListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimInventoryItemsListOK creates a DcimInventoryItemsListOK with default headers values
func NewDcimInventoryItemsListOK() *DcimInventoryItemsListOK {
	return &DcimInventoryItemsListOK{}
}

/*
DcimInventoryItemsListOK describes a response with status code 200, with default header values.

DcimInventoryItemsListOK dcim inventory items list o k
*/
type DcimInventoryItemsListOK struct {
	Payload *DcimInventoryItemsListOKBody
}

// IsSuccess returns true when this dcim inventory items list o k response has a 2xx status code
func (o *DcimInventoryItemsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim inventory items list o k response has a 3xx status code
func (o *DcimInventoryItemsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim inventory items list o k response has a 4xx status code
func (o *DcimInventoryItemsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim inventory items list o k response has a 5xx status code
func (o *DcimInventoryItemsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim inventory items list o k response a status code equal to that given
func (o *DcimInventoryItemsListOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimInventoryItemsListOK) Error() string {
	return fmt.Sprintf("[GET /dcim/inventory-items/][%d] dcimInventoryItemsListOK  %+v", 200, o.Payload)
}

func (o *DcimInventoryItemsListOK) String() string {
	return fmt.Sprintf("[GET /dcim/inventory-items/][%d] dcimInventoryItemsListOK  %+v", 200, o.Payload)
}

func (o *DcimInventoryItemsListOK) GetPayload() *DcimInventoryItemsListOKBody {
	return o.Payload
}

func (o *DcimInventoryItemsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DcimInventoryItemsListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimInventoryItemsListDefault creates a DcimInventoryItemsListDefault with default headers values
func NewDcimInventoryItemsListDefault(code int) *DcimInventoryItemsListDefault {
	return &DcimInventoryItemsListDefault{
		_statusCode: code,
	}
}

/*
DcimInventoryItemsListDefault describes a response with status code -1, with default header values.

DcimInventoryItemsListDefault dcim inventory items list default
*/
type DcimInventoryItemsListDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim inventory items list default response
func (o *DcimInventoryItemsListDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim inventory items list default response has a 2xx status code
func (o *DcimInventoryItemsListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim inventory items list default response has a 3xx status code
func (o *DcimInventoryItemsListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim inventory items list default response has a 4xx status code
func (o *DcimInventoryItemsListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim inventory items list default response has a 5xx status code
func (o *DcimInventoryItemsListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim inventory items list default response a status code equal to that given
func (o *DcimInventoryItemsListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimInventoryItemsListDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/inventory-items/][%d] dcim_inventory-items_list default  %+v", o._statusCode, o.Payload)
}

func (o *DcimInventoryItemsListDefault) String() string {
	return fmt.Sprintf("[GET /dcim/inventory-items/][%d] dcim_inventory-items_list default  %+v", o._statusCode, o.Payload)
}

func (o *DcimInventoryItemsListDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimInventoryItemsListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
DcimInventoryItemsListOKBody dcim inventory items list o k body
swagger:model DcimInventoryItemsListOKBody
*/
type DcimInventoryItemsListOKBody struct {

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
	Results []*models.InventoryItem `json:"results"`
}

// Validate validates this dcim inventory items list o k body
func (o *DcimInventoryItemsListOKBody) Validate(formats strfmt.Registry) error {
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

func (o *DcimInventoryItemsListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("dcimInventoryItemsListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *DcimInventoryItemsListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimInventoryItemsListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimInventoryItemsListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimInventoryItemsListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimInventoryItemsListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("dcimInventoryItemsListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimInventoryItemsListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dcimInventoryItemsListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this dcim inventory items list o k body based on the context it is used
func (o *DcimInventoryItemsListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DcimInventoryItemsListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {
			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimInventoryItemsListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dcimInventoryItemsListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DcimInventoryItemsListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DcimInventoryItemsListOKBody) UnmarshalBinary(b []byte) error {
	var res DcimInventoryItemsListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
