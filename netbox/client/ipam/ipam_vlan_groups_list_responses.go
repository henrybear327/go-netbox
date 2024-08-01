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

// IpamVlanGroupsListReader is a Reader for the IpamVlanGroupsList structure.
type IpamVlanGroupsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVlanGroupsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamVlanGroupsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamVlanGroupsListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamVlanGroupsListOK creates a IpamVlanGroupsListOK with default headers values
func NewIpamVlanGroupsListOK() *IpamVlanGroupsListOK {
	return &IpamVlanGroupsListOK{}
}

/*
IpamVlanGroupsListOK describes a response with status code 200, with default header values.

IpamVlanGroupsListOK ipam vlan groups list o k
*/
type IpamVlanGroupsListOK struct {
	Payload *IpamVlanGroupsListOKBody
}

// IsSuccess returns true when this ipam vlan groups list o k response has a 2xx status code
func (o *IpamVlanGroupsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam vlan groups list o k response has a 3xx status code
func (o *IpamVlanGroupsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam vlan groups list o k response has a 4xx status code
func (o *IpamVlanGroupsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam vlan groups list o k response has a 5xx status code
func (o *IpamVlanGroupsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam vlan groups list o k response a status code equal to that given
func (o *IpamVlanGroupsListOK) IsCode(code int) bool {
	return code == 200
}

func (o *IpamVlanGroupsListOK) Error() string {
	return fmt.Sprintf("[GET /ipam/vlan-groups/][%d] ipamVlanGroupsListOK  %+v", 200, o.Payload)
}

func (o *IpamVlanGroupsListOK) String() string {
	return fmt.Sprintf("[GET /ipam/vlan-groups/][%d] ipamVlanGroupsListOK  %+v", 200, o.Payload)
}

func (o *IpamVlanGroupsListOK) GetPayload() *IpamVlanGroupsListOKBody {
	return o.Payload
}

func (o *IpamVlanGroupsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(IpamVlanGroupsListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamVlanGroupsListDefault creates a IpamVlanGroupsListDefault with default headers values
func NewIpamVlanGroupsListDefault(code int) *IpamVlanGroupsListDefault {
	return &IpamVlanGroupsListDefault{
		_statusCode: code,
	}
}

/*
IpamVlanGroupsListDefault describes a response with status code -1, with default header values.

IpamVlanGroupsListDefault ipam vlan groups list default
*/
type IpamVlanGroupsListDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam vlan groups list default response
func (o *IpamVlanGroupsListDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam vlan groups list default response has a 2xx status code
func (o *IpamVlanGroupsListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam vlan groups list default response has a 3xx status code
func (o *IpamVlanGroupsListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam vlan groups list default response has a 4xx status code
func (o *IpamVlanGroupsListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam vlan groups list default response has a 5xx status code
func (o *IpamVlanGroupsListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam vlan groups list default response a status code equal to that given
func (o *IpamVlanGroupsListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamVlanGroupsListDefault) Error() string {
	return fmt.Sprintf("[GET /ipam/vlan-groups/][%d] ipam_vlan-groups_list default  %+v", o._statusCode, o.Payload)
}

func (o *IpamVlanGroupsListDefault) String() string {
	return fmt.Sprintf("[GET /ipam/vlan-groups/][%d] ipam_vlan-groups_list default  %+v", o._statusCode, o.Payload)
}

func (o *IpamVlanGroupsListDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamVlanGroupsListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
IpamVlanGroupsListOKBody ipam vlan groups list o k body
swagger:model IpamVlanGroupsListOKBody
*/
type IpamVlanGroupsListOKBody struct {

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
	Results []*models.VLANGroup `json:"results"`
}

// Validate validates this ipam vlan groups list o k body
func (o *IpamVlanGroupsListOKBody) Validate(formats strfmt.Registry) error {
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

func (o *IpamVlanGroupsListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("ipamVlanGroupsListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *IpamVlanGroupsListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("ipamVlanGroupsListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *IpamVlanGroupsListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("ipamVlanGroupsListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *IpamVlanGroupsListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("ipamVlanGroupsListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ipamVlanGroupsListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ipamVlanGroupsListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this ipam vlan groups list o k body based on the context it is used
func (o *IpamVlanGroupsListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *IpamVlanGroupsListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {
			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ipamVlanGroupsListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ipamVlanGroupsListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *IpamVlanGroupsListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *IpamVlanGroupsListOKBody) UnmarshalBinary(b []byte) error {
	var res IpamVlanGroupsListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
